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
		valor, situacao, disponivel, id_pessoa
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


func FilterImovelRepository(filter model.Filtro) ([]model.Imovel, error){
	db := config.Connect()

	defer db.Close()

	args := []interface{}{}


	query := `select * from imoveis where 1=1 `

	if filter.Situacao != ""{
		query = query + "and situacao = ?"
		args = append(args, filter.Situacao)
	}
	if filter.Tipo != "" {
		query = query + "and tipo = ?"
		args = append(args, filter.Tipo)
	}
	if filter.Estado != ""{
		query = query + "and estado = ?"
		args = append(args, filter.Estado)
	}
	if filter.Cidade != ""{
		query = query + "and cidade = ?"
		args = append(args, filter.Cidade)
	}
	if filter.De >= 0 && filter.Ate >= 0{ 
		query = query + "and valor between ? and ?"
		args = append(args, filter.De, filter.Ate)
	}
	if filter.Quartos != 0{
		query = query + "and quartos = ?"
		args = append(args, filter.Quartos)
	}
	if filter.Vagas != 0{
		query = query + "and vagas = ?"
		args = append(args, filter.Vagas)
	}
	if filter.Banheiros != 0{
		query = query + "and banheiros = ?"
		args = append(args, filter.Banheiros)
	}


	rows, err := db.Query(query, args...)

	if err != nil{
		return []model.Imovel{}, err
	}

	defer rows.Close()

	var imoveis []model.Imovel

	for rows.Next(){
		var imovel model.Imovel

		err := rows.Scan(
			&imovel.Id, 
			&imovel.Situacao, 
			&imovel.Tipo, 
			&imovel.Estado, 
			&imovel.Cidade, 
			&imovel.Valor, 
			&imovel.Quartos, 
			&imovel.Vagas,
			&imovel.Banheiros,
		)

		if err != nil{
			return nil, err
		}
		imoveis = append(imoveis, imovel)
	}

	return imoveis, nil
	

	
}



func DeleteImovelRepository(id int) (int, error){
	db := config.Connect()

	defer db.Close()

	result, err := db.Exec("delete from imoveis where id = ? ", id)

	if err != nil{
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil

}