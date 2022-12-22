package service

import (
	"context"
	"doan/pkg/model"
	"github.com/praslar/cloud0/ginext"
	"net/http"
)

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (rs *model.User, err error) {

	ob, err := s.repo.GetOneUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !s.CheckPassword(req.Password, ob.Password) {
		return nil, ginext.NewError(http.StatusUnauthorized, "Password is incorrect")
	}
	return ob, nil
}

func (s *UserService) CheckPassword(passwordRequest, password string) bool {
	if password == passwordRequest {
		return true
	}
	return false
}
