package service

import (
	"apiGo/config"
	"apiGo/model"
	"apiGo/repository"
	"apiGo/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateService(p model.Pessoa) (model.Pessoa, error){

	bytes, err := bcrypt.GenerateFromPassword([]byte(p.Senha), 10)


	if err != nil{
		panic(err)
	}

	p.Senha = string(bytes)
	
	id, err := repository.InsertRepository(p)

	if err != nil{
		return  model.Pessoa{}, err
	}

	p.Id = id

	return p, err

}


func LoginService(email string, senha string) (model.Pessoa,model.TokenResponse, error){
	
	db := config.Connect()

	defer db.Close()

    p, err := repository.FindByEmail(db, email)

	if err != nil{
		return model.Pessoa{},  model.TokenResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(p.Senha), []byte(senha))

	if err != nil{
		return model.Pessoa{},  model.TokenResponse{}, err
	}

	if p.Role == ""{
		p.Role = "Usuario"
	}

	token, err := utils.GenerateToken(p.Id, p.Email, p.Role)

	if err != nil{
		return model.Pessoa{},  model.TokenResponse{}, err
	}

	res := model.TokenResponse{
		Token: token,
	}

	
	return p, res, nil
}