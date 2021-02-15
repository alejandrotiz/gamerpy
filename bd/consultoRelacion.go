package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/alejandrotiz/gamerpy/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion se consulta relacion en la bd*/
func ConsultoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("gamerpy")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
