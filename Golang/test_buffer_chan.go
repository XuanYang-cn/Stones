package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	testChannelClose()
}

// For an operand ch whose core type is a channel, the value of the receive operation <-ch is the value received from the channel ch.
// The channel direction must permit receive operations, and the type of the receive operation is the element type of the channel.
// The expression blocks until a value is available.
// Receiving from a nil channel blocks forever.
// A receive operation on a closed channel can always proceed immediately, yielding the element type's zero value after any previously sent values have been received.
// https://go.dev/ref/spec#Receive_operator
func testChannelClose() {
	stopCh := make(chan struct{})

	stopFn := func() {
		select {
		case <-stopCh:
			fmt.Println("return and stop")
		default:
			close(stopCh)
			fmt.Println("close and stop")
		}
	}

	for i := 0; i < 100; i++ {
		stopFn()
	}
}

func testBufferedChannel() {
	var wg sync.WaitGroup
	var bufCh = make(chan struct{}, 10)
	wg.Add(2)

	writer := func(ctx context.Context, bufCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case <-time.Tick(time.Millisecond * 250):
				bufCh <- struct{}{}
			}
		}
	}

	reader := func(ctx context.Context, bufCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case <-time.Tick(time.Second):
				a := len(bufCh)
				if a > 5 {
					fmt.Println("more than five")
					for i := 0; i < 5; i++ {
						<-bufCh
					}
				} else {
					fmt.Println("less than five")
					<-bufCh
				}
			}
		}
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	go reader(ctx, bufCh)
	go writer(ctx, bufCh)

	wg.Wait()

}
