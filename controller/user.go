package controller

import (
	"encoding/json"
	"fmt"
	"market/models"
	"market/pkg/check"
	"net/http"
)

func (c Controller) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUser(w, r)
	case http.MethodGet:
		values := r.URL.Query()

		if _, ok := values["id"]; !ok {
			c.GetUserList(w, r)
		} else {
			c.GetUserByID(w, r)
		}
	case http.MethodPut:
		c.UpdateUser(w, r)
	case http.MethodDelete:
		c.DeleteUser(w, r)
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := models.CreateUser{}

	if err := json.NewDecoder(r.Body).Decode(&CreateUser); err != nil {
		fmt.Println("error while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if !check.PhoneNumber(CreateUser.Phone) {
		fmt.Println("phone number format is not correct", CreateUser.Phone)
		handleResponse(w, http.StatusBadRequest, CreateUser.Phone)
		return
	}

	user, err := c.Store.UserStorage().Create(CreateUser)
	if err != nil {
		fmt.Println("error while creating", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	id := models.PrimaryKey{ID: user.ID}
	resp, err := c.Store.UserStorage().GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusCreated, resp)

}

func (c Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	userID := models.PrimaryKey{ID: id}

	//user := models.User{}

	user, err := c.Store.UserStorage().GetByID(userID)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, user)
}

func (c Controller) GetUserList(w http.ResponseWriter, r *http.Request) {
	//userRequest := models.GetListRequest{}
	users, err := c.Store.UserStorage().GetList()

	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, users)
}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.UpdateUSer{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if !check.PhoneNumber(user.Phone) {
		fmt.Println("phone number format is not correct", user.Phone)
		handleResponse(w, http.StatusBadRequest, user.Phone)
		return
	}

	if _, err := c.Store.UserStorage().Update(user); err != nil {
		fmt.Println("error is while updating user", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	id := models.PrimaryKey{ID: user.ID}

	resp, err := c.Store.UserStorage().GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	userID := models.PrimaryKey{ID: id}

	if err := c.Store.UserStorage().Delete(userID); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}
