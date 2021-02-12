package models

/*Tweet modelo de tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
