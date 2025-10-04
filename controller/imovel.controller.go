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

func FilterImovel(w http.ResponseWriter, r *http.Request){
	var filter model.Filtro

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&filter)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imoveis, err := service.FilterImovelService(filter)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(imoveis)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func DeleteImovel(w http.ResponseWriter, r *http.Request){

	var id model.DeletarImovel

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&id)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rowsAffected, err :=service.DeleteImovelService(id)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(rowsAffected)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}