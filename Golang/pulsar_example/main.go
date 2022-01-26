package main

import (
	"flag"

	"github.com/apache/pulsar-client-go/pulsar"

	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pulsar_example/pkg"
)

const (
	produceInterval = 200 * time.Millisecond
	numOfProducer   = 1
)

func newConsumer(client pulsar.Client, topicName string, subName string) pulsar.Consumer {
	receiveChannel := make(chan pulsar.ConsumerMessage, 1024)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:                       topicName,
		SubscriptionName:            subName,
		Type:                        pulsar.KeyShared,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
		MessageChannel:              receiveChannel,
	})

	if err != nil {
		log.Fatalf("new consumer %s for topic %s failed: %v", subName, topicName, err)
	}

	return consumer
}

func newProducer(client pulsar.Client, topicName string) pulsar.Producer {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})

	if err != nil {
		log.Fatalf("new producer for topic %s failed: %v", topicName, err)
	}

	return producer
}

// produce produces once per 20ms
func produce(ctx context.Context, client pulsar.Client, topicName string) {
	p := newProducer(client, topicName)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Exist produce for context done, producer: %v", p.Name())
			release(nil, p, nil)
			return
		case <-time.After(produceInterval):
			msg := time.Now().String() + "---" + p.Name()

			_, err := p.Send(context.Background(), &pulsar.ProducerMessage{
				Payload: []byte(msg),
			})
			if err != nil {
				log.Printf("%sERROR%s Produce: %v, err: %v", color.Red, color.Reset, msg, err)
				continue
			}

			log.Printf("%s[Produced]%s: %s", color.Green, color.Reset, msg)
		}
	}
}

func consume(ctx context.Context, c pulsar.Consumer) {
	msg, err := c.Receive(ctx)
	if err != nil {
		log.Fatalf("error when receive the first msg: %s", err.Error())
	}
	c.Seek(msg.ID())

	log.Printf("Consumer seek to msgID: %s", msg.ID())

	for {
		select {
		case <-ctx.Done():
			log.Printf("Exist consumer for context done, consumer: %v", c.Name())
			release(nil, nil, c)
			return
		default:
			msg, err := c.Receive(ctx)
			if err != nil {
				log.Printf("%sERROR%s Consume: %v", color.Red, color.Reset, err)
				continue
			}

			log.Printf("%s[Received]%s: -- content: '%s'\n", color.Blue, color.Reset, string(msg.Payload()))

			c.Ack(msg)
		}
	}
}

func release(client pulsar.Client, p pulsar.Producer, c pulsar.Consumer) {
	if c != nil {
		c.Close()
		log.Printf("Closed pulsar consumer: %v", c)
	}

	if p != nil {
		p.Close()
		log.Printf("Closed pulsar producer: %v", p)
	}

	if client != nil {
		client.Close()
		log.Printf("Closed pulsar client: %v", client)
	}
}

var pulsarHost = flag.String("pulsarhost", "pulsar://127.0.0.1:6650", "Pulsar URL")

func testSendAndReceive(url string) {
	opts := pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	}

	client, err := pulsar.NewClient(opts)
	defer release(client, nil, nil)
	if err != nil {
		log.Fatalf("new pulsar client failed: %v", err)
	}

	TopicName := "statistic-channel-twice"

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	for i := 0; i < numOfProducer; i++ {
		go produce(ctx, client, TopicName)
	}

	c := newConsumer(client, TopicName, "only-consumer")
	go consume(ctx, c)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sc
	fmt.Println("Get signal to exit", sig.String())
}

func main() {

	flag.Parse()

	pulsarURL := fmt.Sprintf("pulsar://%s:6650", *pulsarHost)
	log.Printf("Pulsar URL: %v", pulsarURL)

	testSendAndReceive(pulsarURL)

	// ....
}
