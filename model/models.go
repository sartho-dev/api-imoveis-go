package model

type Pessoa struct{
	Id int `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Role string `json:"role"`
}

type Imovel struct{
	Id int `json:"ind"`
	Tipo string `json:"tipo"`
	Rua string `json:"rua"`
	Numero string `json:"numero"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`	
	Estado string `json:"estado"`
	Cep string `json:"cep"`
	Pais string `json:"pais"`
	Area int `json:"area"`
	Quartos int `json:"quartos"`
	Banheiros int `json:"banheiros"`
	Suites int `json:"suites"`
	Vagas int `json:"vagas"`
	Andar int `json:"andar"`
	Valor int `json:"valor"`
	Situacao string `json:"situacao"`
	Disponivel bool `json:"disponivel"`
	IdPessoa int `json:"id_pessoa"`
}

type Login struct{
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type TokenResponse struct{
	Token string `json:"token"`
}