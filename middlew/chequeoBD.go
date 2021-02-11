package middlew

import (
	"net/http"

	"github.com/alejandrotiz/gamerpy/bd"
)

/*ChequeoBD controla si existe conexion a la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(rw, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
