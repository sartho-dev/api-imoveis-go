package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB{
	connect_string := "root:novasenha@tcp(localhost:3306)/imovelGoBanco"

	db, err := sql.Open("mysql", connect_string)

	if err != nil{
		log.Fatal("Erro ao abrir conexão")
	}

	err = db.Ping()

	if err != nil{
		log.Fatal("Erro ao conectar com banco", err)
	}
	return db
}