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
			err := channel.Publish("", "golang_teste", false, false, msg)
			if err != nil {
				log.Fatal("Erro ao publicar mensagem: \n", err.Error())
			}
		}
	}()
	ConsumeMessage()
}
