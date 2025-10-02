package service

import (
	"apiGo/model"
	"apiGo/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreatePessoa(p model.Pessoa) (model.Pessoa, error){

	bytes, err := bcrypt.GenerateFromPassword([]byte(p.Senha), 10)


	if err != nil{
		panic(err)
	}

	p.Senha = string(bytes)
	
	id, err := repository.InsertPessoa(p)

	if err != nil{
		return  model.Pessoa{}, err
	}

	p.Id = id

	return p, err

}