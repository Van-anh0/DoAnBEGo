package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type CinemaService struct {
	repo repo.PGInterface
}

type CinemaInterface interface {
	Create(ctx context.Context, ob model.CinemaRequest) (rs *model.Cinema, err error)
	Update(ctx context.Context, ob model.CinemaRequest) (rs *model.Cinema, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Cinema, err error)
	GetList(ctx context.Context, req model.CinemaParams) (rs *model.CinemaResponse, err error)
}

func NewCinemaService(repo repo.PGInterface) CinemaInterface {
	return &CinemaService{repo: repo}
}

func (s *CinemaService) Create(ctx context.Context, req model.CinemaRequest) (rs *model.Cinema, err error) {

	ob := &model.Cinema{}
	common.Sync(req, ob)

	if err := s.repo.CreateCinema(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CinemaService) Update(ctx context.Context, req model.CinemaRequest) (rs *model.Cinema, err error) {
	ob, err := s.repo.GetOneCinema(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateCinema(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CinemaService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteCinema(ctx, id)
}

func (s *CinemaService) GetOne(ctx context.Context, id string) (rs *model.Cinema, err error) {

	ob, err := s.repo.GetOneCinema(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CinemaService) GetList(ctx context.Context, req model.CinemaParams) (rs *model.CinemaResponse, err error) {
	ob, err := s.repo.GetListCinema(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
