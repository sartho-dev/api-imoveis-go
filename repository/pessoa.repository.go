package repository

import (
	"apiGo/config"
	"apiGo/model"
	"database/sql"
)

func InsertRepository(p model.Pessoa) (int , error){

	result, err := config.Connect().Exec("insert into pessoas (nome, email, senha) values (?, ?, ?)", p.Nome, p.Email, p.Senha)

	
	if err != nil{
		return 0, err
	}
	lastId, _ := result.LastInsertId()

	return int(lastId), nil
	
}



func FindByEmail(db *sql.DB, email string) (model.Pessoa, error){

	var p model.Pessoa

	row := db.QueryRow("select id, nome, email, senha from pessoas where email = ?", email)

	err := row.Scan(&p.Id, &p.Nome, &p.Email, &p.Senha)


	if err != nil{
		return model.Pessoa{}, err
	}

	return p, nil
	
	
}