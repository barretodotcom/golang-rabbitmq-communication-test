package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ConsumeMessage() {
	channel := GetChannel()

	go func() {
		for {
			result, _ := channel.Consume("golang_teste", "", true, false, false, false, amqp.Table{})
			resultString := string((<-result).Body)
			fmt.Print(resultString)
		}
	}()
}
