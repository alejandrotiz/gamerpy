package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC es el objeto de conexion a la bd*/
var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://gamerpy:gamerpy123@gamerpy.p7ako.mongodb.net/gamerpy?retryWrites=true&w=majority")

/*ConectarBD conecta con la base de datos mongodb*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("Conexión exitosa con la bd")
	return client
}

/*ChequeoConexion ping a la base de datos*/
func ChequeoConexion() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
