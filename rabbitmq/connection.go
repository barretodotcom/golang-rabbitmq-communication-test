package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func GetChannel() *amqp.Channel {
	amqpServerUrl := os.Getenv("AMQP_SERVER_URL")

	conn, err := amqp.Dial(amqpServerUrl)

	if err != nil {
		fmt.Println("Erro ao conectar ao Rabbitmq:")
		log.Fatal(err.Error())
	}
	channelRabbitmq, err := conn.Channel()

	_, err = channelRabbitmq.QueueDeclare("golang_teste", true, false, false, false, nil)

	if err != nil {
		fmt.Println("Erro ao declarar a Fila(Queue): ")
		log.Fatal(err.Error())
	}

	return channelRabbitmq
}
