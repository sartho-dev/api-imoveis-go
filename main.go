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

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		controller.Create(db, w, r)
	})



	log.Fatal(http.ListenAndServe(":8080", nil))

}

