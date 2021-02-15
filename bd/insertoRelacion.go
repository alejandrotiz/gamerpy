package bd

import (
	"context"
	"time"

	"github.com/alejandrotiz/gamerpy/models"
)

/*InsertoRelacion graba la relacion en la bd*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("gamerpy")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
