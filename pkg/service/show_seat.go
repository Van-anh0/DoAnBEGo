package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type ShowSeatService struct {
	repo repo.PGInterface
}

type ShowSeatInterface interface {
	Create(ctx context.Context, ob model.ShowSeatRequest) (rs *model.ShowSeat, err error)
	Update(ctx context.Context, ob model.ShowSeatRequest) (rs *model.ShowSeat, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.ShowSeat, err error)
	GetList(ctx context.Context, req model.ShowSeatParams) (rs *model.ShowSeatResponse, err error)
}

func NewShowSeatService(repo repo.PGInterface) ShowSeatInterface {
	return &ShowSeatService{repo: repo}
}

func (s *ShowSeatService) Create(ctx context.Context, req model.ShowSeatRequest) (rs *model.ShowSeat, err error) {

	ob := &model.ShowSeat{}
	common.Sync(req, ob)

	if err := s.repo.CreateShowSeat(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ShowSeatService) Update(ctx context.Context, req model.ShowSeatRequest) (rs *model.ShowSeat, err error) {
	ob, err := s.repo.GetOneShowSeat(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateShowSeat(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ShowSeatService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteShowSeat(ctx, id)
}

func (s *ShowSeatService) GetOne(ctx context.Context, id string) (rs *model.ShowSeat, err error) {

	ob, err := s.repo.GetOneShowSeat(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ShowSeatService) GetList(ctx context.Context, req model.ShowSeatParams) (rs *model.ShowSeatResponse, err error) {

	ob, err := s.repo.GetListShowSeat(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
