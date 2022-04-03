package config

import (
	"context"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
)

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	fmt.Println(l)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		// assign the logger to the reader
		// Logger: l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		SendMail(msg.Value);
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}


func InitKafka() {
	// create a new context
	fmt.Println("Kafka Started...");
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	 
	consume(ctx)
}