package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"strings"
)

func (r *RepoPG) CreateOrderItem(ctx context.Context, ob *model.OrderItem) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateOrderItem(ctx context.Context, ob *model.OrderItem) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteOrderItem(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.OrderItem{}).Error
}

func (r *RepoPG) GetOneOrderItem(ctx context.Context, id string) (*model.OrderItem, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.OrderItem{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListOrderItem(ctx context.Context, req model.OrderItemParams) (*model.OrderItemResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.OrderItemResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	if req.Search != "" {
		tx = tx.Where("unaccent(name) ilike %?%", req.Search)
	}

	if req.Filter != "" {
		filter := strings.Split(req.Filter, ",")
		for i := 0; i < len(filter); i += 2 {
			if i+1 < len(filter) {
				tx = tx.Where(filter[i]+" = ?", filter[i+1])
			}
		}
	}

	switch req.Sort {
	case utils.SORT_CREATED_AT_OLDEST:
		tx = tx.Order("created_at")
	default:
		tx = tx.Order("created_at desc")
	}
	if err := tx.Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

// createMultiOrderItem is a function to create multiple OrderItems
func (r *RepoPG) CreateMultiOrderItem(ctx context.Context, ob *[]model.OrderItem) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}
