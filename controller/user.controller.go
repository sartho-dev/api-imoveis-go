package controller

import (
	"apiGo/model"
	"database/sql"
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

func Create(db *sql.DB, w http.ResponseWriter, r * http.Request){
	var p model.Pessoa

	decod := json.NewDecoder(r.Body)

	err := decod.Decode(&p)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		
	}

	result, err := db.Exec("insert into pessoas (nome, email, senha) VALUES (?,?,?)", 
	p.Nome, p.Email, p.Senha)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()

	p.Id = int(id)

	w.Header().Set("Content-Type", "application/json")


	json.NewEncoder(w).Encode(p)


}