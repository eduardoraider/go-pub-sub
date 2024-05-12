package main

import (
	"fmt"
	"github.com/eduardoraider/go-pub-sub/pkg/pubsub"
	"time"
)

func main() {
	pubSub := pubsub.NewPubSub()

	// Subscriber 1
	sub1 := pubSub.Subscribe("topic1")
	go func() {
		for {
			select {
			case msg := <-sub1:
				fmt.Println("Subscriber 1:", msg)
			}
		}
	}()

	// Subscriber 2
	sub2 := pubSub.Subscribe("topic2")
	go func() {
		for {
			select {
			case msg := <-sub2:
				fmt.Println("Subscriber 2:", msg)
			}
		}
	}()

	// Publisher
	pubSub.Publish("topic1", "Hello, topic 1!")
	pubSub.Publish("topic2", "Greetings, topic 2!")
	pubSub.Publish("topic2", "New on topic 2!")

	// Use a channel to signal termination
	done := make(chan struct{})

	// Wait until all messages are processed by the subscribers
	go func() {
		// Wait until all subscribers finish
		pubSub.Wait()

		// Wait for a short period of time before exiting
		time.Sleep(1 * time.Second)

		// Signal termination
		close(done)
	}()

	// Wait for the channel to close to exit
	<-done

}
