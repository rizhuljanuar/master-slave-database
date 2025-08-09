package handlers

import (

	"database/sql"

	"apps/repositories"
	"apps/models"

)

type ItemHandler struct {
	repo *repositories.ItemRepository

}


func NewItemHandler(repo *repositories.ItemRepository) *ItemHandler {
	return &ItemHandler{repo:repo}
}


func (i *ItemHandler)FindByID(id string) models.Response {
	var (
		rsp = &models.Response{}
	)

	item, err := i.repo.FindByID(id)

	if err == sql.ErrNoRows {

		rsp.WithCode(404)
		rsp.WithMessage(`Item tidak ditemukan`)

		return *rsp
	}
	if err != nil {
		rsp.WithCode(500)
		rsp.WithMessage(`Gagal mengambil data`)

		return *rsp
	}
	rsp.WithCode(200)
	rsp.WithMessage(`success`)
	rsp.WithData(item)

	return *rsp
}


func (i *ItemHandler)FindAll() models.Response {
	var (
		rsp = &models.Response{}
	)

	items, err := i.repo.FindAll()

	if err == sql.ErrNoRows  || items == nil{

		rsp.WithCode(404)
		rsp.WithMessage(`Item tidak tersedia`)

		return *rsp
	}
	if err != nil {
		rsp.WithCode(500)
		rsp.WithMessage(`Gagal mengambil data`)

		return *rsp
	}
	rsp.WithCode(200)
	rsp.WithMessage(`success`)
	rsp.WithData(items)

	return *rsp
}

func (i *ItemHandler)Update(id string, item models.Item) models.Response {
	var (
		rsp = &models.Response{}
	)

	rowsAffected, err := i.repo.Update(id, &item)
	if err != nil {
		rsp.WithCode(500)
		rsp.WithMessage(`Gagal update item`)

		return *rsp
	}

	_ = rowsAffected

	itemU, err := i.repo.FindByID(id)
	if err == sql.ErrNoRows {

		rsp.WithCode(404)
		rsp.WithMessage(`Item tidak ditemukan`)

		return *rsp
	}
	if err != nil {
		rsp.WithCode(500)
		rsp.WithMessage(`Gagal update item`)

		return *rsp
	}


	itemRes := models.Item{
		ID: itemU.ID,
		Title: item.Title,
		Description: item.Description,
		Quantity: item.Quantity,
		Price: item.Price,
		CreatedAt: itemU.CreatedAt,
		UpdatedAt: itemU.UpdatedAt,
	}



	rsp.WithCode(200)
	rsp.WithMessage(`Item berhasil di update`)
	rsp.WithData(itemRes)


	return *rsp
}

func (i *ItemHandler)Store(item models.Item) models.Response {
	var (
		rsp = &models.Response{}
	)


	err := i.repo.Create(&item)
	
	
	if err != nil {
		rsp.WithCode(500)
		rsp.WithMessage(`Gagal membuat item`)

		return *rsp
	}


	rsp.WithCode(200)
	rsp.WithMessage(`success`)
	rsp.WithData(item)


	return *rsp
}



func (i *ItemHandler)Delete(id string) models.Response {
	var (
		rsp = &models.Response{}
	)

	rowsAffected, err := i.repo.Delete(id)
	if err != nil {
		rsp.WithCode(500)
		rsp.WithMessage(`Gagal menghapus item`)

		return *rsp
	}

	if rowsAffected == 0 {
		rsp.WithCode(404)
		rsp.WithMessage(`Item tidak ditemukan`)

		return *rsp
	}

	rsp.WithCode(200)
	rsp.WithMessage(`Item berhasil dihapus`)


	return *rsp
}