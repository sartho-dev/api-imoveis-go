package service

import (
	"apiGo/model"
	"apiGo/repository"
)

func CreateImovelService(im model.Imovel) (model.Imovel, error){


	idImovel, err := repository.InsertImovelRepository(im)

	if err != nil{
		return model.Imovel{}, err
	}

	im.Id = idImovel

	return  im, nil
}

func FilterImovelService(filter model.Filtro) ([]model.Imovel, error){

	return repository.FilterImovelRepository(filter)

}

func DeleteImovelService(im model.DeletarImovel) (int, error){

	return repository.DeleteImovelRepository(im.Id)

}