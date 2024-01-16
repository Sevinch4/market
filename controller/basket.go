package controller

import (
	"encoding/json"
	"fmt"
	"market/models"
	"net/http"
)

func (c Controller) Basket(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateBasket(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			c.GetBasketList(w, r)
		} else {
			c.GetBasketByID(w, r)
		}
	case http.MethodPut:
		c.UpdateBasket(w, r)
	case http.MethodDelete:
		c.DeleteBasket(w, r)
	}
}

func (c Controller) CreateBasket(w http.ResponseWriter, r *http.Request) {
	createBasket := models.CreateBasket{}

	if err := json.NewDecoder(r.Body).Decode(&createBasket); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	basket, err := c.Store.Basket().CreateBasket(createBasket)
	if err != nil {
		fmt.Println("error is while creating basket", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	id := models.PrimaryKey{ID: basket.ID}
	res, err := c.Store.Basket().GetBasketByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}
	handleResponse(w, http.StatusCreated, res)
}

func (c Controller) GetBasketByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	idBasket := models.PrimaryKey{ID: id}

	fmt.Println("return id", idBasket)
	basket, err := c.Store.Basket().GetBasketByID(idBasket)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, basket)
}

func (c Controller) GetBasketList(w http.ResponseWriter, r *http.Request) {
	baskets, err := c.Store.Basket().GetBasketList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, baskets)
}

func (c Controller) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	updatedBasket := models.UpdateBasket{}

	if err := json.NewDecoder(r.Body).Decode(&updatedBasket); err != nil {
		fmt.Println("error is while decoding ", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if updatedBasket.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, updatedBasket.ID)
		return
	}

	updatedBasket.ID = id

	if _, err := c.Store.Basket().UpdateBasket(updatedBasket); err != nil {
		fmt.Println("error is while updating basket", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	ids := models.PrimaryKey{ID: id}
	res, err := c.Store.Basket().GetBasketByID(ids)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, res)
}

func (c Controller) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	basketID := models.PrimaryKey{ID: id}
	if err := c.Store.Basket().DeleteBasket(basketID); err != nil {
		fmt.Println("error is while deleting basket", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, nil)
}
