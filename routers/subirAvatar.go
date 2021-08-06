package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	// fmt.Println(handler)
	fmt.Println(file)
	fmt.Println(err)
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

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
	usuario.Avatar = IDUsuario + "." + extension
	status, err := bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Copy avatar in DB failed "+err.Error(), http.StatusBadRequest)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
