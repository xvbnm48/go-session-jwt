package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xvbnm48/go-session-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("gagal mengdecode json")
	}

	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("gagal menambahkan user")
	}

	response, _ := json.Marshal(map[string]string{
		"message": "user berhasil ditambahkan",
	})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {}
