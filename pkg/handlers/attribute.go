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

type AttributeHandlers struct {
	service service.AttributeInterface
}

func NewAttributeHandlers(service service.AttributeInterface) *AttributeHandlers {
	return &AttributeHandlers{service: service}
}

// Create
// @Tags Create
// @Accept  json
// @Produce  json
// @Param data body model.AttributeRequest true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/attribute/create [post]
func (h *AttributeHandlers) Create(r *ginext.Request) (*ginext.Response, error) {
	req := model.AttributeRequest{}
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
// @Param data body model.AttributeRequest true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/attribute/update/:id [put]
func (h *AttributeHandlers) Update(r *ginext.Request) (*ginext.Response, error) {
	id := utils.GetIdFromUri(r.GinCtx)
	if id == nil {
		return nil, ginext.NewError(http.StatusForbidden, "Wrong ID")
	}

	req := model.AttributeRequest{}
	r.MustBind(&req)
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
// @Router /api/v1/attribute/delete/:id [delete]
func (h *AttributeHandlers) Delete(r *ginext.Request) (*ginext.Response, error) {
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
// @Router /api/v1/attribute/get-one/:id [get]
func (h *AttributeHandlers) GetOne(r *ginext.Request) (*ginext.Response, error) {

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
// @Router /api/v1/attribute/get-list [get]
func (h *AttributeHandlers) GetList(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.AttributeParams{}
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
