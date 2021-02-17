package controllers

import (
	"context"
	"fmt"
	"gostorage/models"

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

func (c StorageController) Upload(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	// tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer tempFile.Close()

	// // read all of the contents of our uploaded file into a
	// // byte array
	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // write this byte array to our temporary file
	// tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
