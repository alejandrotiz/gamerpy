package bd

import (
	"context"
	"log"
	"time"

	"github.com/alejandrotiz/gamerpy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets devuelve los tweets*/
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("gamerpy")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(map[string]int{"fecha": -1})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, error := col.Find(ctx, condicion, opciones)
	if error != nil {
		log.Fatal(error.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}
	return resultados, true

}
