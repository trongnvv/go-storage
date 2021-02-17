package controllers

import (
	"encoding/json"
	"gostorage/helpers"
	"gostorage/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseControllerInterface interface {
}

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c BaseController) decodeRequestBody(w http.ResponseWriter, r *http.Request, req interface{}) error {
	return json.NewDecoder(r.Body).Decode(req)
}

func (c BaseController) respond(
	w http.ResponseWriter,
	data interface{},
	status int,
	message string,
) {
	helpers.Respond(w, helpers.BaseResponseBody{Data: data, Status: status, Message: message})
}

func (c BaseController) defaultInsertDB() *models.BaseSchema {
	return &models.BaseSchema{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
