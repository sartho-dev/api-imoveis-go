package main

import (
	"apiGo/config"
	"apiGo/controller"
	"log"
	"net/http"
)



func main(){
	db := config.Connect()
	defer db.Close()

	http.HandleFunc("/", controller.Handler)

	http.HandleFunc("/create", controller.Create)



	log.Fatal(http.ListenAndServe(":8080", nil))

}

