package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type OrderItemService struct {
	repo repo.PGInterface
}

type OrderItemInterface interface {
	Create(ctx context.Context, ob model.OrderItemRequest) (rs *model.OrderItem, err error)
	Update(ctx context.Context, ob model.OrderItemRequest) (rs *model.OrderItem, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.OrderItem, err error)
	GetList(ctx context.Context, req model.OrderItemParams) (rs *model.OrderItemResponse, err error)
}

func NewOrderItemService(repo repo.PGInterface) OrderItemInterface {
	return &OrderItemService{repo: repo}
}

func (s *OrderItemService) Create(ctx context.Context, req model.OrderItemRequest) (rs *model.OrderItem, err error) {

	ob := &model.OrderItem{}
	common.Sync(req, ob)

	if err := s.repo.CreateOrderItem(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *OrderItemService) Update(ctx context.Context, req model.OrderItemRequest) (rs *model.OrderItem, err error) {
	ob, err := s.repo.GetOneOrderItem(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateOrderItem(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *OrderItemService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteOrderItem(ctx, id)
}

func (s *OrderItemService) GetOne(ctx context.Context, id string) (rs *model.OrderItem, err error) {

	ob, err := s.repo.GetOneOrderItem(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *OrderItemService) GetList(ctx context.Context, req model.OrderItemParams) (rs *model.OrderItemResponse, err error) {

	ob, err := s.repo.GetListOrderItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
