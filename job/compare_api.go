package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/segmentio/ksuid"
)

var (
	brokers  string
	topic    string
	records  int
	certFile string
	caFile   string
	keyFile  string
)

func init() {
	flag.StringVar(&brokers, "brokers", "localhost:9092", "broker addresses, comma-separated")
	flag.StringVar(&topic, "topic", "topic", "topic to produce to")
	flag.IntVar(&records, "records", 250000, "number of records to read from kafka")
	flag.StringVar(&certFile, "cert", "_cert.pem", "tls cert")
	flag.StringVar(&caFile, "ca", "_ca.pem", "tls ca")
	flag.StringVar(&keyFile, "key", "_key.pem", "tls key")
	flag.Parse()
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	confluent()
	saramago()
}

func confluent() {
	groupID := ksuid.New().String()
	cm := &kafka.ConfigMap{
		"session.timeout.ms":              6000,
		"metadata.broker.list":            brokers,
		"enable.auto.commit":              false,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"group.id":                        groupID,
		"default.topic.config": kafka.ConfigMap{
			"auto.offset.reset": "earliest",
		},

		"security.protocol":        "ssl",
		"ssl.ca.location":          caFile,
		"ssl.certificate.location": certFile,
		"ssl.key.location":         keyFile,
	}

	consumer, err := kafka.NewConsumer(cm)
	check(err)
	defer consumer.Close()

	check(consumer.Subscribe(topic, nil))

	var start time.Time
	count := 0

loop:
	for {
		select {
		case m, ok := <-consumer.Events():
			if !ok {
				panic("unexpected eof")
			}

			switch event := m.(type) {
			case kafka.AssignedPartitions:
				consumer.Assign(event.Partitions)

			case kafka.PartitionEOF:
				// nop

			case kafka.RevokedPartitions:
				consumer.Unassign()

			case *kafka.Message:
				count++
				if count == 1 {
					start = time.Now()
				}
				if count == records {
					break loop
				}

			default:
				panic(m)
			}
		}
	}
	elapsed := time.Now().Sub(start)
	fmt.Printf("confluent: %v records, %v\n", count, elapsed)
}

func tlsConfig() *tls.Config {
	certPEM, err := ioutil.ReadFile(certFile)
	check(err)

	caPEM, err := ioutil.ReadFile(caFile)
	check(err)

	keyPEM, err := ioutil.ReadFile(keyFile)
	check(err)

	if certPEM == nil || caPEM == nil || keyPEM == nil {
		panic("tls configuration not available")
	}

	cert, err := tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	check(err)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(caPEM))

	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
}

func saramago() {
	// Code from the sarama-cluster same page with minimal changes for TLS and broker/topic
	//

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsConfig()

	// init consumer
	groupID := ksuid.New().String()
	consumer, err := cluster.NewConsumer(strings.Split(brokers, ","), groupID, []string{topic}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for n := range consumer.Notifications() {
			log.Printf("Notification: %v\n", n)
		}
	}()

	var start time.Time
	var count int

	// consume messages, watch signals
loop:
	for {
		select {
		case _, ok := <-consumer.Messages():
			if !ok {
				panic("messages channel unexpectedly closed")
			}

			count++

			if count == 1 {
				start = time.Now()
			}

			if count == records {
				break loop
			}

		case <-signals:
			return
		}
	}
	elapsed := time.Now().Sub(start)
	fmt.Printf("sarama-cluster: %v records, %v\n", count, elapsed)
}
