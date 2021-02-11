package bd

import "golang.org/x/crypto/bcrypt"

/*EnctriptarPassword encripta pass con costo 8*/
func EnctriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
