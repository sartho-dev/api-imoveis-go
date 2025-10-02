package main

import (
	"apiGo/config"
	"apiGo/controller"
	"apiGo/middlewares"
	"log"
	"net/http"
)



func main(){
	db := config.Connect()
	defer db.Close()

	http.HandleFunc("/", controller.Handler)

	http.HandleFunc("/create/user", controller.Create)

	http.HandleFunc("/login/user", controller.Login)

	http.HandleFunc("/create/imovel", middlewares.AuthMiddleware(controller.CreateImovel))


	log.Fatal(http.ListenAndServe(":8080", nil))


}

