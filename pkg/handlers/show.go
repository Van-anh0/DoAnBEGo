package handlers

import (
	"doan/pkg/model"
	"doan/pkg/service"
	"doan/pkg/utils"
	"github.com/praslar/cloud0/ginext"
	"github.com/praslar/cloud0/logger"
	"github.com/praslar/lib/common"
	"net/http"
)

type ShowHandlers struct {
	service service.ShowtimeInterface
}

func NewShowtimeHandlers(service service.ShowtimeInterface) *ShowHandlers {
	return &ShowHandlers{service: service}
}

// Create
// @Tags Create
// @Accept  json
// @Produce  json
// @Param data body model.ShowtimeRequest true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/show-time/create [post]
func (h *ShowHandlers) Create(r *ginext.Request) (*ginext.Response, error) {
	req := model.ShowtimeRequest{}
	r.MustBind(&req)

	if err := common.CheckRequireValid(req); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	data, err := h.service.Create(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, data), nil
}

// Update
// @Tags Update
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param data body model.ShowtimeRequest true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/show-time/update/:id [put]
func (h *ShowHandlers) Update(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	id := utils.GetIdFromUri(r.GinCtx)
	if id == nil {
		return nil, ginext.NewError(http.StatusForbidden, "Wrong ID")
	}

	req := model.ShowtimeRequest{}
	if err := r.GinCtx.ShouldBind(&req); err != nil {
		log.WithError(err).Error("error_400: error parse")
		return nil, ginext.NewError(http.StatusBadRequest, "Yêu cầu không hợp lệ")
	}
	req.ID = id

	if err := common.CheckRequireValid(req); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	data, err := h.service.Update(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, data), nil
}

// Delete
// @Tags Delete
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} interface{}
// @Router /api/v1/show-time/delete/:id [delete]
func (h *ShowHandlers) Delete(r *ginext.Request) (*ginext.Response, error) {
	id := utils.ParseIDFromUri(r.GinCtx)
	if id == nil {
		return nil, ginext.NewError(http.StatusForbidden, "Wrong ID")
	}

	if err := h.service.Delete(r.Context(), id.String()); err != nil {
		return nil, err
	}
	return ginext.NewResponse(http.StatusOK), nil
}

// GetOne
// @Tags GetOne
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} interface{}
// @Router /api/v1/show-time/get-one/:id [get]
func (h *ShowHandlers) GetOne(r *ginext.Request) (*ginext.Response, error) {

	id := utils.ParseIDFromUri(r.GinCtx)
	if id == nil {
		return nil, ginext.NewError(http.StatusForbidden, "Wrong ID")
	}

	data, err := h.service.GetOne(r.Context(), id.String())
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, data), nil
}

// GetList
// @Tags GetList
// @Accept  json
// @Produce  json
// @Param data body model.BlacklistParam true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/show/get-list [get]
func (h *ShowHandlers) GetList(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.ShowParams{}
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: error parse")
		return nil, ginext.NewError(http.StatusBadRequest, "Yêu cầu không hợp lệ")
	}

	data, err := h.service.GetList(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return &ginext.Response{Code: http.StatusOK, GeneralBody: &ginext.GeneralBody{
		Data: data.Data,
		Meta: data.Meta,
	}}, nil
}

// GetListGroupDay
// @Tags GetListGroupDay
// @Accept  json
// @Produce  json
// @Param data body model.BlacklistParam true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/show/get-list-group [get]
func (h *ShowHandlers) GetListGroupDay(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.ShowParams{}
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: error parse")
		return nil, ginext.NewError(http.StatusBadRequest, "Yêu cầu không hợp lệ")
	}

	data, err := h.service.GetListGroupByDay(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return &ginext.Response{Code: http.StatusOK, GeneralBody: &ginext.GeneralBody{
		Data: data.Data,
		Meta: data.Meta,
	}}, nil
}

// GetListGroupMovie
// @Tags GetListGroupMovie
// @Accept  json
// @Produce  json
// @Param data body model.BlacklistParam true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/show/get-list-group-movie [get]
func (h *ShowHandlers) GetListGroupMovie(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.ShowParams{}
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: error parse")
		return nil, ginext.NewError(http.StatusBadRequest, "Yêu cầu không hợp lệ")
	}

	data, err := h.service.GetListGroupMovie(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return &ginext.Response{Code: http.StatusOK, GeneralBody: &ginext.GeneralBody{
		Data: data.Data,
		Meta: data.Meta,
	}}, nil
}
