package repository

import (
	"apiGo/config"
	"apiGo/model"
)

func InsertPessoa(p model.Pessoa) (int , error){

	result, err := config.Connect().Exec("insert into pessoas (nome, email, senha) values (?, ?, ?)", p.Nome, p.Email, p.Senha)

	
	if err != nil{
		return 0, err
	}
	lastId, _ := result.LastInsertId()

	return int(lastId), nil
	
}
