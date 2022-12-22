package handlers

import (
	"doan/pkg/model"
	"doan/pkg/utils"
	"github.com/praslar/cloud0/ginext"
	"github.com/praslar/cloud0/logger"
	"net/http"
)

// Login
// @Tags Login
// @Accept  json
// @Produce  json
// @Param data body model.BlacklistParam true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/auth/login [post]
func (h *UserHandlers) Login(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.LoginRequest{}
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: error parse")
		return nil, ginext.NewError(http.StatusBadRequest, "Yêu cầu không hợp lệ")
	}

	data, err := h.service.Login(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, data), nil
}
