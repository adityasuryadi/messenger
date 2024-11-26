package config

import (
	"log/slog"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SetupDB(viper *viper.Viper) (*mongo.Database, error) {
	// db := viper.Get("DB")
	// dbHost := viper.Get("DB_HOST")
	// dbPort := viper.Get("DB_PORT")
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err != nil {
		slog.Error("Failed to connect to database", err)
		return nil, err
	}
	client.Database("messenger")

	return client.Database("messenger"), nil
}
