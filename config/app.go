package config

import (
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BootstrapConfig struct {
	Mux *http.ServeMux
	DB  *mongo.Database
}
