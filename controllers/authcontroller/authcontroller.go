package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/xvbnm48/go-session-jwt/config"
	"github.com/xvbnm48/go-session-jwt/helper"
	"github.com/xvbnm48/go-session-jwt/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	// take data username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "username atau password tidak ditemukan",
			}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// check password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{
			"message": "username atau password tidak ditemukan",
		}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// generate token
	expTime := time.Now().Add(time.Hour * 24)
	claims := &config.JwtClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-session-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// algorithm for signing token
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// response token with json
	// response := map[string]string{
	// 	"token": token,
	// }

	// helper.ResponseJSON(w, http.StatusOK, response)

	// response token with cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	response := map[string]string{
		"message": "login berhasil",
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{
		"message": "user berhasil ditambahkan",
	}
	helper.ResponseJSON(w, http.StatusCreated, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {}
