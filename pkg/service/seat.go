package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/praslar/lib/common"
)

type SeatService struct {
	repo repo.PGInterface
}

type SeatInterface interface {
	Create(ctx context.Context, ob model.SeatRequest) (rs *model.Seat, err error)
	Update(ctx context.Context, ob model.SeatRequest) (rs *model.Seat, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Seat, err error)
}

func NewSeatService(repo repo.PGInterface) SeatInterface {
	return &SeatService{repo: repo}
}

func (s *SeatService) Create(ctx context.Context, req model.SeatRequest) (rs *model.Seat, err error) {

	ob := &model.Seat{}
	common.Sync(req, ob)

	if err := s.repo.CreateSeat(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SeatService) Update(ctx context.Context, req model.SeatRequest) (rs *model.Seat, err error) {

	ob := &model.Seat{}
	common.Sync(req, ob)

	if err := s.repo.UpdateSeat(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SeatService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteSeat(ctx, id)
}

func (s *SeatService) GetOne(ctx context.Context, id string) (rs *model.Seat, err error) {

	ob, err := s.repo.GetOneSeat(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
