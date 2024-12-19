package mock

import (
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/stretchr/testify/mock"
)

type AuthUsecaseMock struct {
	mock.Mock
}

func (a *AuthUsecaseMock) RefreshToken(token string) (string, error) {
	args := a.Called(token)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), nil
}

func (a *AuthUsecaseMock) Register(request *model.RegisterRequest) {
	panic("implement me")
}
func (a *AuthUsecaseMock) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	panic("implement me")
}
