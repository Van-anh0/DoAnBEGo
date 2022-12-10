package handlers

import (
	"doan/pkg/model"
	"doan/pkg/service"
	"doan/pkg/utils"
	"github.com/praslar/cloud0/ginext"
	"github.com/praslar/lib/common"
	"net/http"
)

type RoomHandlers struct {
	service service.RoomInterface
}

func NewRoomHandlers(service service.RoomInterface) *RoomHandlers {
	return &RoomHandlers{service: service}
}

// Create
// @Tags Create
// @Accept  json
// @Produce  json
// @Param data body model.RoomRequest true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/room/create [post]
func (h *RoomHandlers) Create(r *ginext.Request) (*ginext.Response, error) {
	req := model.RoomRequest{}
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
// @Param data body model.RoomRequest true "body data"
// @Success 200 {object} interface{}
// @Router /api/v1/room/update/:id [put]
func (h *RoomHandlers) Update(r *ginext.Request) (*ginext.Response, error) {
	req := model.RoomRequest{}
	r.MustBind(&req)

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
// @Router /api/v1/room/delete/:id [delete]
func (h *RoomHandlers) Delete(r *ginext.Request) (*ginext.Response, error) {
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
// @Router /api/v1/room/get-one/:id [get]
func (h *RoomHandlers) GetOne(r *ginext.Request) (*ginext.Response, error) {

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
