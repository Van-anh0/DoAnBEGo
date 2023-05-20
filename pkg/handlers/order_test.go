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
	"github.com/praslar/lib/pointer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrderHandlers_Create(t *testing.T) {
	// create mock
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	mockService := mocks.NewMockOrderInterface(ctrl)
	reqData := model.OrderRequest{
		UserID:     pointer.String("123e4567-e89b-12d3-a456-426614174000"),
		ShowtimeId: pointer.String("123e4567-e89b-12d3-a456-426614174000"),
		OrderItem: &[]model.OrderItem{
			{
				ProductId: "123e4567-e89b-12d3-a456-426614174000",
				Quantity:  1,
			},
		},
		ShowSeat: &[]model.ShowSeat{
			{
				SeatId:     "123e4567-e89b-12d3-a456-426614174000",
				ShowtimeId: "123e4567-e89b-12d3-a456-426614174000",
			},
		},
	}
	mockService.EXPECT().Create(context.Background(), reqData).Return(&model.Order{}, nil)

	// Setup
	h := &OrderHandlers{
		service: mockService,
	}

	type args struct {
		UserID     *string            `json:"user_id"`
		ShowtimeId *string            `json:"showtime_id"`
		OrderItem  *[]model.OrderItem `json:"order_item"`
		ShowSeat   *[]model.ShowSeat  `json:"show_seat"`
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		// TODO: Add test cases.
		{
			name: "happy flow",
			args: args{
				UserID:     pointer.String("123e4567-e89b-12d3-a456-426614174000"),
				ShowtimeId: pointer.String("123e4567-e89b-12d3-a456-426614174000"),
				OrderItem: &[]model.OrderItem{
					{
						ProductId: "123e4567-e89b-12d3-a456-426614174000",
						Quantity:  1,
					},
				},
				ShowSeat: &[]model.ShowSeat{
					{
						SeatId:     "123e4567-e89b-12d3-a456-426614174000",
						ShowtimeId: "123e4567-e89b-12d3-a456-426614174000",
					},
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "bad flow: email invalid",
			args: args{
				ShowtimeId: pointer.String("123e4567-e89b-12d3-a456-426614174000"),
				OrderItem: &[]model.OrderItem{
					{
						ProductId: "123e4567-e89b-12d3-a456-426614174000",
						Quantity:  1,
					},
				},
				ShowSeat: &[]model.ShowSeat{
					{
						SeatId:     "123e4567-e89b-12d3-a456-426614174000",
						ShowtimeId: "123e4567-e89b-12d3-a456-426614174000",
					},
				},
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			reqBody, err := json.Marshal(tt.args)
			if err != nil {
				fmt.Println(err)
				return
			}
			req, _ := http.NewRequest("POST", "/order/create", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(resp)

			ginCtx.Request = req
			r := &ginext.Request{GinCtx: ginCtx}
			h.Create(r)

			// Verify
			if resp.Code != tt.wantStatus {
				t.Errorf("Login() got = %v, want %v", resp.Code, tt.wantStatus)
			}
		})
	}
}
