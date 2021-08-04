package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinguilhem/twittor/bd"
)

/* VerPerfil extacts profile values */
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send Id param", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "an error occurred while trying to find the profile "+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
