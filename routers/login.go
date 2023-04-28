package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luis501angel/go-course/db"
	"github.com/luis501angel/go-course/jwt"
	"github.com/luis501angel/go-course/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(u.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	document, exist := db.TryLogin(u.Email, u.Password)

	if !exist {
		http.Error(w, "Usuario y/o contraseña inválidos", 400)
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
