package util

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/nats-io/stan.go"
)

func TestPub(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sc, err := NewNSConnection("", "", "pubClient", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	// Simple Synchronous Publisher
	sc.Publish("foo", []byte("Hello World"+fmt.Sprintln(time.Now())))
}

func TestSub(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sc, err := NewNSConnection("", "", "subClient", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	// Simple Async Subscriber
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()

	time.Sleep(60 * time.Second)
}
