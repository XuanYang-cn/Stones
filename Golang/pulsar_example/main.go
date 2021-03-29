package main

import (
	"github.com/apache/pulsar-client-go/pulsar"

	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"unsafe"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:9876", nil))
	}()

	newClient := func() {
		var client pulsar.Client

		opts := pulsar.ClientOptions{
			URL: "pulsar://127.0.0.1:6650",
		}

		client, err := pulsar.NewClient(opts)
		if err != nil {
			panic(err)
		}

		channelName := "insert"
		pp, err := client.CreateProducer(pulsar.ProducerOptions{Topic: channelName})
		if err != nil {
			panic(err)
		}
		fmt.Println("producer: ", pp)

		receiveChannel := make(chan pulsar.ConsumerMessage, 1024)
		pc, err := client.Subscribe(pulsar.ConsumerOptions{
			Topic:                       channelName,
			SubscriptionName:            channelName,
			Type:                        pulsar.KeyShared,
			SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
			MessageChannel:              receiveChannel,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("consumer: ", pc)
		pp.Close()
		pc.Close()

		client.Close()
		f := reflect.ValueOf(client).Elem().FieldByName("cnxPool")
		f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
		f.MethodByName("Close").Call(nil)
	}

	sc := make(chan os.Signal, 1)

	go func() {
		for i := 0; i < 10000; i++ {
			newClient()
		}
	}()

	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sc
	fmt.Println("Get signal to exit", sig.String())

}
