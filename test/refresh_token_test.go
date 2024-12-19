package test

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	controller "github.com/adityasuryadi/messenger/internal/auth/delivery/http"
// 	"github.com/adityasuryadi/messenger/internal/auth/delivery/http/route"
// 	mockObject "github.com/adityasuryadi/messenger/internal/auth/mock"
// 	"github.com/adityasuryadi/messenger/internal/auth/repository"
// 	"github.com/adityasuryadi/messenger/internal/auth/usecase"
// 	"github.com/adityasuryadi/messenger/pkg"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func setUpRoute() *http.ServeMux {
// 	db, _ := SetupTestDB()
// 	validation := pkg.NewValidation(db)

// 	userRepository := repository.NewUserRepository(db)
// 	refreshTokenRepository := repository.NewRefreshTokenRepository(db)
// 	authUsecase := usecase.NewAuthUseCase(userRepository, refreshTokenRepository)
// 	authController := controller.NewAuthController(validation, authUsecase)

// 	route := route.NewRouter(authController)

// 	return route
// }

// func TestRefreshTokenController(t *testing.T) {
// 	router := setUpRoute()
// 	t.Run("Test Refresh Token token not exist in cookie", func(t *testing.T) {
// 		request := httptest.NewRequest("POST", "/api/refresh-token", nil)
// 		recorder := httptest.NewRecorder()

// 		router.ServeHTTP(recorder, request)
// 		response := recorder.Result()

// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, int(401), response.StatusCode)
// 		assert.Equal(t, float64(401), responseBody["code"])
// 		assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
// 	})

// 	t.Run("Test Refresh Token Success", func(t *testing.T) {
// 		request := httptest.NewRequest("POST", "/api/refresh-token", nil)
// 		recorder := httptest.NewRecorder()

// 		refreshToken := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjZhZGViZTE3LTBlNzgtNDI5My1iMzUwLWY4YjM1YmZjOWU5MyIsImV4cCI6MTczNTY5ODMxMCwiaWF0IjoxNzMzMTA2MzEwLCJqdGkiOiI2YWRlYmUxNy0wZTc4LTQyOTMtYjM1MC1mOGIzNWJmYzllOTMifQ.F6DpqVit2BAFXrChN0cNPq7H2wm5iwnN0xSUEAR3FM8`

// 		request.AddCookie(&http.Cookie{
// 			Name:  "refresh_token",
// 			Value: refreshToken,
// 			Path:  "/api",
// 		})

// 		router.ServeHTTP(recorder, request)
// 		response := recorder.Result()

// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)
// 		data := responseBody["data"].(map[string]interface{})

// 		assert.Equal(t, int(200), response.StatusCode)
// 		assert.Equal(t, float64(200), responseBody["code"])
// 		assert.Equal(t, "OK", responseBody["status"])
// 		assert.NotEqual(t, "", data["access_token"])

// 	})

// 	t.Run("Test Refresh Token avaiable", func(t *testing.T) {
// 		db, _ := SetupTestDB()
// 		validation := pkg.NewValidation(db)

// 		authUsecase := new(mockObject.AuthUsecaseMock)
// 		authController := controller.NewAuthController(validation, authUsecase)
// 		authUsecase.On("RefreshToken", mock.Anything).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjZhZGViZTE3LTBlNzgtNDI5My1iMzUwLWY4YjM1YmZjOWU5MyIsImV4cCI6MTczNTY5ODMxMCwiaWF0IjoxNzMzMTA2MzEwLCJqdGkiOiI2YWRlYmUxNy0wZTc4LTQyOTMtYjM1MC1mOGIzNWJmYzllOTMifQ.F6DpqVit2BAFXrChN0cNPq7H2wm5iwnN0xSUEAR3FM8", nil).Once()

// 		router := route.NewRouter(authController)
// 		request := httptest.NewRequest("POST", "/api/refresh-token", nil)
// 		request.AddCookie(&http.Cookie{
// 			Name:  "refresh_token",
// 			Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjZhZGViZTE3LTBlNzgtNDI5My1iMzUwLWY4YjM1YmZjOWU5MyIsImV4cCI6MTczNTY5ODMxMCwiaWF0IjoxNzMzMTA2MzEwLCJqdGkiOiI2YWRlYmUxNy0wZTc4LTQyOTMtYjM1MC1mOGIzNWJmYzllOTMifQ.F6DpqVit2BAFXrChN0cNPq7H2wm5iwnN0xSUEAR3FM8",
// 			Path:  "/api",
// 		})
// 		recorder := httptest.NewRecorder()
// 		router.ServeHTTP(recorder, request)
// 		response := recorder.Result()

// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, int(200), response.StatusCode)
// 		assert.Equal(t, float64(200), responseBody["code"])
// 		assert.Equal(t, "OK", responseBody["status"])

// 	})

// 	t.Run("Test Refresh Token Expired", func(t *testing.T) {
// 		db, _ := SetupTestDB()
// 		validation := pkg.NewValidation(db)

// 		authUsecase := new(mockObject.AuthUsecaseMock)
// 		authController := controller.NewAuthController(validation, authUsecase)
// 		authUsecase.On("RefreshToken", mock.Anything).Return(nil, nil).Once()

// 		router := route.NewRouter(authController)
// 		request := httptest.NewRequest("POST", "/api/refresh-token", nil)
// 		request.AddCookie(&http.Cookie{
// 			Name:  "refresh_token",
// 			Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjZhZGViZTE3LTBlNzgtNDI5My1iMzUwLWY4YjM1YmZjOWU5MyIsImV4cCI6MTczNTY5ODMxMCwiaWF0IjoxNzMzMTA2MzEwLCJqdGkiOiI2YWRlYmUxNy0wZTc4LTQyOTMtYjM1MC1mOGIzNWJmYzllOTMifQ.F6DpqVit2BAFXrChN0cNPq7H2wm5iwnN0xSUEAR3FM8",
// 			Path:  "/api",
// 		})
// 		recorder := httptest.NewRecorder()
// 		router.ServeHTTP(recorder, request)
// 		response := recorder.Result()

// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, int(200), response.StatusCode)
// 		assert.Equal(t, float64(200), responseBody["code"])
// 		assert.Equal(t, "OK", responseBody["status"])

// 	})
// }
