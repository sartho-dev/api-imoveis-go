package repository

import (
	"apiGo/config"
	"apiGo/model"
)

func InsertImovelRepository(im model.Imovel) (int, error) {
	
	db := config.Connect()
	defer db.Close()

	query := `
	INSERT INTO imoveis (
		tipo, rua, numero, bairro, cidade, estado, cep, pais,
		area, quartos, banheiros, suites, vagas, andar,
		valor, situacao, disponivel, idPessoa
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	result, err := db.Exec(query,
		im.Tipo, im.Rua, im.Numero, im.Bairro, im.Cidade, im.Estado, im.Cep,
		im.Pais, im.Area, im.Quartos, im.Banheiros, im.Suites, im.Vagas, im.Andar,
		im.Valor, im.Situacao, im.Disponivel, im.IdPessoa,
	)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}
