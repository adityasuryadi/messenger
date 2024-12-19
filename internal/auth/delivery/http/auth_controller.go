package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/adityasuryadi/messenger/helper"
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/adityasuryadi/messenger/internal/auth/usecase"
	"github.com/adityasuryadi/messenger/pkg"
)

type AuthController struct {
	AuthUsecase usecase.AuthUseCase
	Validation  *pkg.Validation
}

func NewAuthController(validation *pkg.Validation, authUsecase usecase.AuthUseCase) *AuthController {
	return &AuthController{
		AuthUsecase: authUsecase,
		Validation:  validation,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	// request from body
	request := model.RegisterRequest{}
	helper.ReadRequestBody(r, &request)

	err := c.Validation.ValidateRequest(&request)
	if err != nil {
		errValidation := c.Validation.ErrorJson(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := &model.ErrorResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Error:  errValidation,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	c.AuthUsecase.Register(&request)
	w.Header().Add("Content-Type", "application/json")
	response := &model.SuccessResponse{
		Status: "OK",
		Code:   200,
		Data:   nil,
	}
	helper.WriteResponseBody(w, response)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	request := model.LoginRequest{}
	helper.ReadRequestBody(r, &request)

	err := c.Validation.ValidateRequest(&request)
	if err != nil {
		errValidation := c.Validation.ErrorJson(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := &model.ErrorResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Error:  errValidation,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	response, err := c.AuthUsecase.Login(&request)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		response := &model.ErrorResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Error:  err.Error(),
		}
		helper.WriteResponseBody(w, response)
		return
	}

	successResponse := model.SuccessResponse{
		Status: "OK",
		Code:   200,
		Data:   response,
	}

	// set refresh token to cookie
	cookie := &http.Cookie{}
	cookie.Name = "refresh_token"
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)
	cookie.Value = response.RefreshToken
	http.SetCookie(w, cookie)

	w.Header().Add("Content-Type", "application/json")

	helper.WriteResponseBody(w, successResponse)
}

func (c *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {

	type RefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	refreshToken := r.CookiesNamed("refresh_token")
	if len(refreshToken) == 0 {
		err := errors.New("refresh token not found")
		helper.WriteUnauthorizedResponse(w, err)
		return
	}

	refreshTokenValue := refreshToken[0].Value
	token, err := c.AuthUsecase.RefreshToken(refreshTokenValue)
	if err != nil {
		helper.WriteUnauthorizedResponse(w, err)
		return
	}
	response := &RefreshTokenResponse{
		AccessToken: token,
	}

	helper.WriteOkResponse(w, response)
}

func (c *AuthController) Logout() {

}
