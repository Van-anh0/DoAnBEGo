package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type TicketService struct {
	repo repo.PGInterface
}

type TicketInterface interface {
	Create(ctx context.Context, ob model.TicketRequest) (rs *model.OrderItem, err error)
	Update(ctx context.Context, ob model.TicketRequest) (rs *model.OrderItem, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.OrderItem, err error)
	GetList(ctx context.Context, req model.TicketParams) (rs *model.TicketResponse, err error)
}

func NewTicketService(repo repo.PGInterface) TicketInterface {
	return &TicketService{repo: repo}
}

func (s *TicketService) Create(ctx context.Context, req model.TicketRequest) (rs *model.OrderItem, err error) {

	ob := &model.OrderItem{}
	common.Sync(req, ob)

	if err := s.repo.CreateTicket(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *TicketService) Update(ctx context.Context, req model.TicketRequest) (rs *model.OrderItem, err error) {
	ob, err := s.repo.GetOneTicket(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateTicket(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *TicketService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteTicket(ctx, id)
}

func (s *TicketService) GetOne(ctx context.Context, id string) (rs *model.OrderItem, err error) {

	ob, err := s.repo.GetOneTicket(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *TicketService) GetList(ctx context.Context, req model.TicketParams) (rs *model.TicketResponse, err error) {

	ob, err := s.repo.GetListTicket(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
