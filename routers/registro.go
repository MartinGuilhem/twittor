package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/models"
)

/* Registro method creates a user registry in DB */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid input "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "User Email is a required field", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password should have at least 6 characters", 400)
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "User email already exist", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error occurred while trying to register user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "User registration failed", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
