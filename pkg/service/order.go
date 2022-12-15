package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/praslar/lib/common"
)

type OrderService struct {
	repo repo.PGInterface
}

type OrderInterface interface {
	Create(ctx context.Context, ob model.OrderRequest) (rs *model.Order, err error)
	Update(ctx context.Context, ob model.OrderRequest) (rs *model.Order, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Order, err error)
	GetList(ctx context.Context, req model.OrderParams) (rs *model.OrderResponse, err error)
}

func NewOrderService(repo repo.PGInterface) OrderInterface {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(ctx context.Context, req model.OrderRequest) (rs *model.Order, err error) {

	ob := &model.Order{}
	common.Sync(req, ob)

	if err := s.repo.CreateOrder(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *OrderService) Update(ctx context.Context, req model.OrderRequest) (rs *model.Order, err error) {

	ob := &model.Order{}
	common.Sync(req, ob)

	if err := s.repo.UpdateOrder(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *OrderService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteOrder(ctx, id)
}

func (s *OrderService) GetOne(ctx context.Context, id string) (rs *model.Order, err error) {

	ob, err := s.repo.GetOneOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *OrderService) GetList(ctx context.Context, req model.OrderParams) (rs *model.OrderResponse, err error) {

	ob, err := s.repo.GetListOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
