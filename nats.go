package util

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//
// checks provided nats topic only has alphanumeric & dot separators within the name
//
var topicRegex = regexp.MustCompile("^[A-Za-z0-9]([A-Za-z0-9.]*[A-Za-z0-9])?$")

//
// do regex check on topic names provided for nats
//
func ValidateNatsTopic(tName string) (bool, error) {

	valid := topicRegex.Match([]byte(tName))
	if valid {
		return valid, nil
	}
	return false, errors.New("Nats topic names must be alphanumeric only, can also contain (but not start or end with) period ( . ) as token delimiter.")

}

//
// creates new connection to nats streaming server
//
func NewConnection(host, cluster, client string, port int) (stan.Conn, error) {

	// Send PINGs every 10 seconds, and fail after 5 PINGs without any response.
	sc, err := stan.Connect(cluster, client,
		stan.NatsURL(fmt.Sprintf("nats://%s:%d", host, port)),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("\n\tReader shuting down: Connection to streaming server lost, reason: %v\n\n", reason)
			// attempt clean shutdown by raising sig int
			p, _ := os.FindProcess(os.Getpid())
			p.Signal(os.Interrupt)
		}))
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func NewNSConnection(host, cluster, client string, port int) (stan.Conn, error) {

	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(cluster) == 0 {
		cluster = "test-cluster" // default nats-streaming-server clusterID
	}
	if port == 0 {
		port = 4222 // nats-stream-server default port
	}

	return NewConnection(host, cluster, client, port)
}
