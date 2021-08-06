package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Upload image failed "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Copy image failed "+err.Error(), http.StatusBadRequest)
		return
	}
	var usuario models.Usuario
	usuario.Banner = IDUsuario + "." + extension
	status, err := bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Copy banner in DB failed "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
