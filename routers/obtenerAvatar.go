package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/martinguilhem/twittor/bd"
)

func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID param is required", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Copy image failed", http.StatusBadRequest)
		return
	}

}
