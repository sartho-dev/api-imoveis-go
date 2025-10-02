package controller

import (
	"apiGo/model"
	"apiGo/service"
	"encoding/json"
	"net/http"
)

func CreateImovel(w http.ResponseWriter, r *http.Request){

	var imovel model.Imovel

	decode := json.NewDecoder(r.Body)

	err := decode.Decode(&imovel)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imovelResult, err := service.CreateImovelService(imovel)
	
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(imovelResult)
	

}