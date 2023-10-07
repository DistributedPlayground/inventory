package message

import (
	"encoding/json"
	"fmt"

	"github.com/DistributedPlayground/inventory/database"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Message interface {
	Listen() error
}

type Product struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type message struct {
	kc    *kafka.Consumer
	redis database.RedisStore
	topic string
}

func NewMessage(kc *kafka.Consumer, redis database.RedisStore) Message {
	return &message{kc: kc, redis: redis, topic: "product"}
}
func (m message) Listen() error {

	// subscribe to a topic
	err := m.kc.Subscribe(m.topic, nil)
	if err != nil {
		panic("Failed to subscribe to topic: " + err.Error())
	}

	// continuously poll for new messages
	for {
		msg, err := m.kc.ReadMessage(-1)
		if err != nil {
			// handle error
			fmt.Printf("Error reading message: %s\n", err)
			continue
		}

		product := Product{}
		err = json.Unmarshal(msg.Value, &product)
		if err != nil {
			fmt.Printf("Error unmarshalling message: %s\n", err)
			continue
		}

		// process the message
		fmt.Printf("Received message: %s\n", string(msg.Value))
		for _, header := range msg.Headers {
			if string(header.Key) == "MessageType" {
				switch string(header.Value) {
				case "Create":
					m.redis.Set(product.Id, product.Quantity, 0)
				case "Update":
					m.redis.Set(product.Id, product.Quantity, 0)
				case "Delete":
					m.redis.Delete(product.Id)
				default:
					fmt.Println("Unknown")
				}
			}
		}
	}
}
