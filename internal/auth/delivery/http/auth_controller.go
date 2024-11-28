package controller

import (
	"encoding/json"
	"net/http"

	"github.com/adityasuryadi/messenger/helper"
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/adityasuryadi/messenger/internal/auth/usecase"
	"github.com/adityasuryadi/messenger/pkg"
)

type AuthController struct {
	AuthUsecase *usecase.AuthUseCase
	Validation  *pkg.Validation
}

func NewAuthController(validation *pkg.Validation, authUsecase *usecase.AuthUseCase) *AuthController {
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

	w.Header().Add("Content-Type", "application/json")
	helper.WriteResponseBody(w, successResponse)
}

func (c *AuthController) RefreshToken() {

}

func (c *AuthController) Logout() {

}
