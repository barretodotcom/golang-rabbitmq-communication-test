package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func SendMessage() {
	channel := GetChannel()

	msg := amqp.Publishing{
		Body: []byte("{message}"),
	}

	go func() {
		for {
			err := channel.Publish("", "golang_test", false, false, msg)
			if err != nil {
				log.Fatal("Error while publishing message: \n", err.Error())
			}
		}
	}()
	ConsumeMessage()
}
