package main

import (
	"log"

	"github.com/alejandrotiz/gamerpy/bd"
	"github.com/alejandrotiz/gamerpy/handlres"
)

func main() {

	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlres.Manejadores()

}
