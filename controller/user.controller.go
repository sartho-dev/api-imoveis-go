package controller

import (
	"apiGo/model"
	"apiGo/service"
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	answer := map[string]string{
		"mensagem": "Olá do back",
	}

	json.NewEncoder(w).Encode(answer)
	
	
}

func Create(w http.ResponseWriter, r * http.Request){
	
	var p model.Pessoa

	decod := json.NewDecoder(r.Body)

	err := decod.Decode(&p)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		
	}

	result, err := service.CreateService(p)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")


	json.NewEncoder(w).Encode(result)


}

func Login(w http.ResponseWriter, r * http.Request){

	var login model.Login

	decode := json.NewDecoder(r.Body)

	err := decode.Decode(&login)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	p, token , err := service.LoginService(login.Email, login.Senha)
	
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string] interface{}{
		"user": p,
		"token": token,
	})

}