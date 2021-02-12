package routers

import (
	"encoding/json"
	"net/http"

	"github.com/alejandrotiz/gamerpy/bd"
	"github.com/alejandrotiz/gamerpy/models"
)

/*ModificarPerfil valida los datos a modificar*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, error := bd.ModificoRegistro(t, IDUsuario)
	if error != nil {
		http.Error(w, "ocurrio un error al intentar modificar"+error.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro modificar el registro de usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
