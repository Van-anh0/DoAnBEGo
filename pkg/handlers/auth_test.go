package handlers

import (
	"bytes"
	"context"
	"doan/mocks"
	"doan/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/praslar/cloud0/ginext"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandlers_Login(t *testing.T) {
	// create mock
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	mockService := mocks.NewMockUserInterface(ctrl)
	reqData := model.LoginRequest{
		Email:    "duchieu.ctk41@gmail.com",
		Password: "password",
	}
	mockService.EXPECT().Login(context.Background(), reqData).Return(&model.User{}, nil)

	// Setup
	h := &UserHandlers{
		service: mockService,
	}

	type args struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	cases := []struct {
		name       string
		args       args
		wantStatus int
	}{
		// TODO: Add test cases.
		{
			name: "happy flow",
			args: args{
				Email:    "duchieu.ctk41@gmail.com",
				Password: "password",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "bad flow: email invalid",
			args: args{
				Email:    "duchieu",
				Password: "password",
			},
			wantStatus: http.StatusUnauthorized,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			reqBody, err := json.Marshal(tt.args)
			if err != nil {
				fmt.Println(err)
				return
			}
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(resp)

			ginCtx.Request = req
			r := &ginext.Request{GinCtx: ginCtx}
			h.Login(r)

			// Verify
			if resp.Code != tt.wantStatus {
				t.Errorf("Login() got = %v, want %v", resp.Code, tt.wantStatus)
			}
		})
	}
}
