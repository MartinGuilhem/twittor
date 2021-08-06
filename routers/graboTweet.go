package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "An error occur when trying to instert registry "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Insert Tweet failed ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
