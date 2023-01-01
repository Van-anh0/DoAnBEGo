package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type MovieRankService struct {
	repo repo.PGInterface
}

type MovieRankInterface interface {
	Create(ctx context.Context, ob model.MovieRankRequest) (rs *model.MovieRank, err error)
	Update(ctx context.Context, ob model.MovieRankRequest) (rs *model.MovieRank, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.MovieRank, err error)
	GetList(ctx context.Context, req model.MovieRankParams) (rs *model.MovieRankResponse, err error)
}

func NewMovieRankService(repo repo.PGInterface) MovieRankInterface {
	return &MovieRankService{repo: repo}
}

func (s *MovieRankService) Create(ctx context.Context, req model.MovieRankRequest) (rs *model.MovieRank, err error) {

	ob := &model.MovieRank{}
	common.Sync(req, ob)

	if err := s.repo.CreateMovieRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MovieRankService) Update(ctx context.Context, req model.MovieRankRequest) (rs *model.MovieRank, err error) {
	ob, err := s.repo.GetOneMovieRank(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateMovieRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MovieRankService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteMovieRank(ctx, id)
}

func (s *MovieRankService) GetOne(ctx context.Context, id string) (rs *model.MovieRank, err error) {

	ob, err := s.repo.GetOneMovieRank(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MovieRankService) GetList(ctx context.Context, req model.MovieRankParams) (rs *model.MovieRankResponse, err error) {

	ob, err := s.repo.GetListMovieRank(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
