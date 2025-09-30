package model

type Pessoa struct{
	Id int `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Imovel struct{
	Id int `json:"ind"`
	Tipo string `json:"tipo"`
	Situacao string `json:"situacao"`
	IdPessoa int `json:"id_pessoa"`
}