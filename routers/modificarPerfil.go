package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid input "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "An error occur when trying to update data "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Update Data failed ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
