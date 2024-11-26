package test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adityasuryadi/messenger/internal/delivery/http/route"
	controller "github.com/adityasuryadi/messenger/internal/handler/http"
	"github.com/adityasuryadi/messenger/internal/repository"
	"github.com/adityasuryadi/messenger/internal/usecase"
	"github.com/adityasuryadi/messenger/pkg"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SetupTestDB() (*mongo.Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err != nil {
		panic(err)
		return nil, err
	}
	client.Database("messenger")

	// client.Database("messenger").Collection("user").DeleteMany(context.TODO(), bson.M{})

	return client.Database("messenger"), nil
}

func TestRegisterSuccess(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya@mail.com",
		"fullname": "aditya",
		"password": "test1234",
		"password_confirmation": "test1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, int(200), response.StatusCode)
	assert.Equal(t, float64(200), responseBody["code"])
	assert.Equal(t, "OK", responseBody["status"])
}

func TestFullnameEmpty(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya2@mail.com",
		"fullname": "",
		"password": "test1234",
		"password_confirmation": "test1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "field fullname tidak boleh kosong", errData["fullname"].([]interface{})[0])
}

func TestPasswordEmpty(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya2@mail.com",
		"fullname": "aditya",
		"password": "",
		"password_confirmation": "test1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "field password tidak boleh kosong", errData["password"].([]interface{})[0])
}

func TestPasswordMin(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya2@mail.com",
		"fullname": "aditya",
		"password": "1234",
		"password_confirmation": "1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "minimal 6 karakter", errData["password"].([]interface{})[0])
}

func TestPasswordConfirmationEmpty(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya2@mail.com",
		"fullname": "aditya",
		"password": "test1234",
		"password_confirmation": ""
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "field password_confirmation tidak boleh kosong", errData["password_confirmation"].([]interface{})[0])
}

func TestPasswordConfirmationMin(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya2@mail.com",
		"fullname": "aditya",
		"password": "test1234",
		"password_confirmation": "1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "minimal 6 karakter", errData["password_confirmation"].([]interface{})[0])
}

func TestPasswordConfirmationNotMatch(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya2@mail.com",
		"fullname": "aditya",
		"password": "test1234",
		"password_confirmation": "test12343333"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "field tidak sama dengan Password", errData["password_confirmation"].([]interface{})[0])
}

func TestEmailRequiredFailed(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "",
		"fullname": "test1234",
		"password": "test1234",
		"password_confirmation": "test1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "field email tidak boleh kosong", errData["email"].([]interface{})[0])
}

func TestRegisterEmailAlreadyExists(t *testing.T) {
	db, _ := SetupTestDB()
	validation := pkg.NewValidation(db)

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	route := route.NewRouter(authController)

	requestBody := strings.NewReader(`{
		"email": "aditya@mail.com",
		"fullname": "test1234",
		"password": "test1234",
		"password_confirmation": "test1234"
	}`)
	request := httptest.NewRequest("POST", "/api/register", requestBody)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	errData := responseBody["error"].(map[string]interface{})

	assert.Equal(t, int(400), response.StatusCode)
	assert.Equal(t, float64(400), responseBody["code"])
	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
	assert.Equal(t, "email not avaiable", errData["email"].([]interface{})[0])
}
