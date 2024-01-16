package storage

import "market/models"

type IStorage interface {
	Close()
	UserStorage() IUserStorage
	Basket() IBasket
}

type IUserStorage interface {
	Create(models.CreateUser) (models.User, error)
	GetByID(models.PrimaryKey) (models.User, error)
	//GetList(models.GetListRequest) (models.UserResponse, error)
	GetList() (models.UserResponse, error)
	Update(models.UpdateUSer) (models.User, error)
	Delete(models.PrimaryKey) error
}

type IBasket interface {
	CreateBasket(models.CreateBasket) (models.Basket, error)
	GetBasketByID(models.PrimaryKey) (models.Basket, error)
	GetBasketList() (models.BasketResponse, error)
	//GetBasketList(models.GetListRequest) (models.BasketResponse, error)
	UpdateBasket(models.UpdateBasket) (models.Basket, error)
	DeleteBasket(key models.PrimaryKey) error
}
