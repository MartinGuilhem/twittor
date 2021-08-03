package main

import (
	"log"

	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
