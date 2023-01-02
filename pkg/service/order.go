package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/cloud0/ginext"
	"github.com/praslar/lib/common"
	"net/http"
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

	// check length OrderItem is not empty
	//if ob.OrderItem == nil || len(ob.OrderItem) == 0 {
	//	return nil, ginext.NewError(http.StatusBadRequest, "OrderItem is empty")
	//}

	// check length ShowSeat is not empty
	if ob.ShowSeat == nil || len(ob.ShowSeat) == 0 {
		return nil, ginext.NewError(http.StatusBadRequest, "ShowSeat is empty")
	}

	if err := s.repo.CreateOrder(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil

	if ob.OrderItem == nil || len(ob.OrderItem) == 0 {
		for _, v := range ob.OrderItem {
			v.OrderId = ob.ID
		}

		// create list OrderItem
		if err = s.repo.CreateMultiOrderItem(ctx, &ob.OrderItem); err != nil {
			return nil, err
		}
	}

	for _, v := range ob.ShowSeat {
		v.OrderId = ob.ID
	}

	// create list ShowSeat
	if err = s.repo.CreateMultiShowSeat(ctx, &ob.ShowSeat); err != nil {
		return nil, err
	}

	return ob, nil
}

func (s *OrderService) Update(ctx context.Context, req model.OrderRequest) (rs *model.Order, err error) {
	ob, err := s.repo.GetOneOrder(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

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
