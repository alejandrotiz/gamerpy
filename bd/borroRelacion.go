package bd

import (
	"context"
	"time"

	"github.com/alejandrotiz/gamerpy/models"
)

/*BorroRelacion borra relacion de la bd*/
func BorroRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("gamerpy")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
