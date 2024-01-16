package controller

import (
	"encoding/json"
	"fmt"
	"market/models"
	"market/storage"
	_ "market/storage/postgres"
	"net/http"
)

type Controller struct {
	Store storage.IStorage
}

func New(store storage.IStorage) Controller {
	return Controller{
		Store: store,
	}
}

func handleResponse(w http.ResponseWriter, statuscode int, data interface{}) {
	resp := models.Response{}

	switch code := statuscode; {
	case code < 400:
		resp.Description = "succes"
	case code < 500:
		resp.Description = "bad request"
	default:
		resp.Description = "internal server error"
	}

	resp.StatusCode = statuscode
	resp.Data = data

	js, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error is while marshalling json", err.Error())
		return
	}

	w.WriteHeader(statuscode)
	w.Write(js)
}
