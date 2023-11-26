package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/elue-dev/todoapi/controllers"
	"github.com/elue-dev/todoapi/helpers"
	"github.com/elue-dev/todoapi/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := r.ParseMultipartForm(10 << 20)
    if err != nil {
		json.NewEncoder(w).Encode(models.ErrResponse{
			Success: false,
			Error: "Please provide username, email and password",
		})
        return
    }

	user.Username = r.FormValue("username")
    user.Email = r.FormValue("email")
    user.Password = r.FormValue("password")


	isValidated := helpers.ValidateSignUpFields(user.Username, user.Email, user.Password)

	if !isValidated {
		json.NewEncoder(w).Encode(models.ErrResponse{
			Success: false,
			Error: "Please provide username, email and password",
		})
	} 

	hashedPassword, err := helpers.HashPassword(user.Password)
	 if err != nil {
		fmt.Println("Could not hash user password", err)
		 return
	 }

	 user.Password = hashedPassword

	 file, _, err := r.FormFile("avatar")
	 if err != nil {
		 log.Fatalf("Failed to get avatar from form: %v", err)
		 return
	 }
	 defer file.Close()
 
	 cld, err := cloudinary.New()
	 if err != nil {
		 log.Fatalf("Failed to initialize Cloudinary: %v", err)
		 return
	 }
 
	 var ctx = context.Background()
	 
	 uploadResult, err := cld.Upload.Upload(
		 ctx,
		 file,
         uploader.UploadParams{PublicID: "avatar"})
 
	 if err != nil {
		 log.Fatalf("Failed to upload file: %v\n", err)
		 return
	 }
  
	 user.Avatar = &uploadResult.SecureURL


	 result, err := controllers.RegisterUser(user)

	 if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrResponse{
			Success: false,
			Error: err.Error(),
		 })
		return
	}
	
	 json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data: result,
	 })
}