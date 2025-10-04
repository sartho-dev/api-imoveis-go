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

	http.HandleFunc("/criar/usuario", controller.Create)

	http.HandleFunc("/login/usuario", controller.Login)

	http.HandleFunc("/criar/imovel", middlewares.AuthMiddleware(controller.CreateImovel))

	http.HandleFunc("/filtrar/imoveis", controller.FilterImovel)	

	http.HandleFunc("/deletar/imovel", middlewares.AuthMiddleware(controller.DeleteImovel))
	
	log.Fatal(http.ListenAndServe(":8080", nil))


}

