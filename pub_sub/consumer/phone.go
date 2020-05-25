package main

import (
	"encoding/json"
	"fmt"
	"pub_sub/components"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
		"group.id":          "phone",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"foo", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		value := string(msg.Value)
		var rb components.Request_body
		json.Unmarshal([]byte(value), &rb)

		mb := rb.Message_body
		tid := rb.Transaction_id
		email := rb.Email
		mobile := rb.Phone
		key := rb.Key

		if err == nil {
			fmt.Printf("Message on %s: key %s \n", msg.TopicPartition, key)
			fmt.Printf("%s: Transaction ID: %s, Message: %s\n Email: %s\n Mobile: %s\n", msg.TopicPartition, tid, mb, email, mobile)
			components.Mail(email, tid, mb)
			//components.SMS(mobile, tid, mb)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	//c.Close()
}
