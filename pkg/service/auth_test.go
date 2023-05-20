package service

import (
	"context"
	"doan/mocks"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	// Mock expectation
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	mockRepo := mocks.NewMockPGInterface(ctrl)
	mockRepo.EXPECT().GetOneUserByEmail(context.Background(), "duchieu.ctk41@gmail.com").Return(&model.User{
		Email:    "duchieu.ctk41@gmail.com",
		Password: "password",
	}, nil)

	// Setup
	// phải đặt mock bên ngoài
	s := &UserService{
		repo: mockRepo,
	}
	type args struct {
		ctx context.Context
		req model.LoginRequest
	}
	cases := []struct {
		name string
		args args
		want *model.User
	}{
		// TODO: Add test cases.
		{
			name: "happy flow",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Email:    "duchieu.ctk41@gmail.com",
					Password: "password",
				},
			},
			want: &model.User{
				Email:    "duchieu.ctk41@gmail.com",
				Password: "password",
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if err != nil {
				t.Errorf("Login() error = %v", err)
				return
			}
			if got.Email != tt.want.Email {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CheckPassword(t *testing.T) {
	type fields struct {
		repo repo.PGInterface
	}
	type args struct {
		passwordRequest string
		password        string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "happy flow",
			args: args{
				passwordRequest: "password",
				password:        "password",
			},
			want: true,
		},
		{
			name: "bad flow",
			args: args{
				passwordRequest: "password1",
				password:        "password",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repo: tt.fields.repo,
			}
			if got := s.CheckPassword(tt.args.passwordRequest, tt.args.password); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
