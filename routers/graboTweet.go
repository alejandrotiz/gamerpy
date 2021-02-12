package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alejandrotiz/gamerpy/bd"
	"github.com/alejandrotiz/gamerpy/models"
)

/*GraboTweet obtiene los datos e inserta en la bd*/
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
		http.Error(w, "Ocurrio un error al insertar tweet"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se inserto tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
