package main

import (
	"fmt"
	"net/http"

	"golang_rabbitmq/rabbitmq"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	r := mux.NewRouter()
	rabbitmq.SendMessage()
	http.ListenAndServe(":3000", r)
}
