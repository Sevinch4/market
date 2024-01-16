package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"market/config"
	"market/controller"
	"market/storage/postgres"
	"net/http"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}

	defer store.Close()

	con := controller.New(store)

	http.HandleFunc("/user", con.User)
	http.HandleFunc("/basket", con.Basket)

	fmt.Println("server runnig....")
	http.ListenAndServe("localhost:8080", nil)
}
