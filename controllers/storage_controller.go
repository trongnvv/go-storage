package controllers

import (
	"context"
	"fmt"
	"gostorage/models"
	"io"
	"os"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

type StorageController struct {
	*BaseController
}

type reqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type resRegister struct {
	Token string `json:"token"`
}

func NewStorageController() *StorageController {
	controller := NewBaseController()
	return &StorageController{BaseController: controller}
}

func (c StorageController) Register(w http.ResponseWriter, r *http.Request) {
	var req reqRegister

	if err := c.decodeRequestBody(w, r, &req); err != nil || req.Password == "" || req.Username == "" {
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.FileModel.FindOne(ctx, bson.M{"username": req.Username}).Decode(&models.UserSchema{}); err == nil {
		c.respond(w, nil, http.StatusBadRequest, "Account existed")
		return
	}

	// newUser := models.UserSchema{
	// 	BaseSchema: c.defaultInsertDB(),
	// 	Username:   req.Username,
	// 	Password:   helpers.HashPassword(req.Password),
	// }

}

func registFile(name string, isError ...bool) (*os.File, error) {
	path := "upload"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return nil, err
		}
	}

	dst, err := os.Create("upload/" + name)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (c StorageController) Upload(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("[Upload] Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("[Upload] File Size: %+v\n", handler.Size)

	dst, err := registFile(handler.Filename)
	defer dst.Close()
	if err != nil {
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}

	c.respond(w, nil, http.StatusBadRequest, "Successfully Uploaded File")
}
