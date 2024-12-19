package entity

type Message struct {
	From    string `bson:"from"`
	Message string `bson:"message"`
}
