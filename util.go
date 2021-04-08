package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/nats-io/nuid"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	hashids "github.com/speps/go-hashids"
)

var (
	once      sync.Once
	netClient *http.Client
)

//
// create a singleton http client to ensure
// maximum reuse of connection
//
func newNetClient() *http.Client {
	once.Do(func() {
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 2 * time.Second,
		}
		netClient = &http.Client{
			Timeout:   time.Second * 2,
			Transport: netTransport,
		}
	})

	return netClient
}

//
// generate a short useful unique name - hashid in this case
// defaultName can be upper project name like "aligner", "reader", "leveller"
//
func GenerateName(defaultName string) string {

	name := defaultName // "aligner", "reader", "leveller"

	// generate a random number
	number0, err := rand.Int(rand.Reader, big.NewInt(10000000))
	if err != nil {
		log.Fatalf("error generating number %v", err)
	}

	hd := hashids.NewData()
	// hd.Salt = "otf-align random name generator 2020" // otf-reader, otf-level
	hd.Salt = fmt.Sprintf("%s random name generator %s", defaultName, time.Now().Format("01-02-2006"))
	hd.MinLength = 5
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Println("error auto-generating name: ", err)
		return name
	}
	e, err := h.EncodeInt64([]int64{number0.Int64()})
	if err != nil {
		log.Println("error encoding auto-generated name: ", err)
		return name
	}
	name = e

	return name

}

//
// generate a unique id - nuid in this case
//
func GenerateID() string {
	return nuid.Next()
}

//
// small utility function embedded in major ops
// to print a performance indicator.
//
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed.Truncate(time.Millisecond).String())
}

//
// find an available tcp port
//
func AvailablePort() (int, error) {

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, errors.Wrap(err, "cannot acquire a tcp port")
	}

	return listener.Addr().(*net.TCPAddr).Port, nil

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

//
// Makes network calls to other services (text-class, nias), and returns
// the response payload as bytes, or an error
//
// method - http method to invoke (post/put/get etc.)
// header - map of headers to include in request
// body - reader for any content to supply as request body
//
func Fetch(method string, url string, header map[string]string, body io.Reader) ([]byte, error) {

	// //
	// // TODO: turn off in production
	// //
	// fmt.Printf("\nmethod:%v\nurl:%v\n,header:%+v\n\n", method, url, header)

	// Create request.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// //
	// // TODO: turn off in production
	// //
	// reqDump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	//  fmt.Println("req-dump error: ", err)
	// }
	// fmt.Printf("\noutbound request\n\n%s\n\n", reqDump)

	// Add any required headers.
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// Perform the network call.
	res, err := newNetClient().Do(req)
	if err != nil {
		return nil, err
	}

	// //
	// // TODO: turn off in production
	// //
	// responseDump, err := httputil.DumpResponse(res, true)
	// if err != nil {
	//  fmt.Println("resp-dump error: ", err)
	// }
	// fmt.Printf("\nresponse:\n\n%s\n\n", responseDump)

	// If response from network call is not 200, return error.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Network call failed with response: %d", res.StatusCode))
	}

	// return response payload as bytes
	respByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read Fetch response")
	}
	res.Body.Close()

	return respByte, nil
}
