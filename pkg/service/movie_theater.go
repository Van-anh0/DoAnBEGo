package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/praslar/lib/common"
)

type MovieTheaterService struct {
	repo repo.PGInterface
}

type MovieTheaterInterface interface {
	Create(ctx context.Context, ob model.MovieTheaterRequest) (rs *model.MovieTheater, err error)
	Update(ctx context.Context, ob model.MovieTheaterRequest) (rs *model.MovieTheater, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.MovieTheater, err error)
}

func NewMovieTheaterService(repo repo.PGInterface) MovieTheaterInterface {
	return &MovieTheaterService{repo: repo}
}

func (s *MovieTheaterService) Create(ctx context.Context, req model.MovieTheaterRequest) (rs *model.MovieTheater, err error) {

	ob := &model.MovieTheater{}
	common.Sync(req, ob)

	if err := s.repo.CreateMovieTheater(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MovieTheaterService) Update(ctx context.Context, req model.MovieTheaterRequest) (rs *model.MovieTheater, err error) {

	ob := &model.MovieTheater{}
	common.Sync(req, ob)

	if err := s.repo.UpdateMovieTheater(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MovieTheaterService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteMovieTheater(ctx, id)
}

func (s *MovieTheaterService) GetOne(ctx context.Context, id string) (rs *model.MovieTheater, err error) {

	ob, err := s.repo.GetOneMovieTheater(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
