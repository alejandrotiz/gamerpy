package routers

import (
	"errors"
	"strings"

	"github.com/alejandrotiz/gamerpy/bd"
	"github.com/alejandrotiz/gamerpy/models"
	"github.com/dgrijalva/jwt-go"
)

/*Email que se usara en todos los endpoints*/
var Email string

/*IDUsuario que se usara en todos los endpoints*/
var IDUsuario string

/*ProcesoToken para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("gamerpydesarrollo")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
			println("se guarda id de usuario" + IDUsuario)
			println("se guarda EMAIL usuario" + Email)
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
