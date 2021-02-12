package bd

import (
	"context"
	"time"

	"github.com/alejandrotiz/gamerpy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro actualiza en la bd el registro*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("gamerpy")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellido) > 0 {
		registro["apellidos"] = u.Apellido
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	/*if len(u.Banner) > 0 {
		registro["banner"] = u.Nombre
	}*/
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	println("el obj ID es " + ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}
