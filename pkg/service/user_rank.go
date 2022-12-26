package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type RankService struct {
	repo repo.PGInterface
}

type RankInterface interface {
	Create(ctx context.Context, ob model.RankRequest) (rs *model.Rank, err error)
	Update(ctx context.Context, ob model.RankRequest) (rs *model.Rank, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Rank, err error)
	GetList(ctx context.Context, req model.RankParams) (rs *model.RankResponse, err error)
}

func NewRankService(repo repo.PGInterface) RankInterface {
	return &RankService{repo: repo}
}

func (s *RankService) Create(ctx context.Context, req model.RankRequest) (rs *model.Rank, err error) {

	ob := &model.Rank{}
	common.Sync(req, ob)

	if err := s.repo.CreateRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *RankService) Update(ctx context.Context, req model.RankRequest) (rs *model.Rank, err error) {
	ob, err := s.repo.GetOneRank(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *RankService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteRank(ctx, id)
}

func (s *RankService) GetOne(ctx context.Context, id string) (rs *model.Rank, err error) {

	ob, err := s.repo.GetOneRank(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *RankService) GetList(ctx context.Context, req model.RankParams) (rs *model.RankResponse, err error) {

	ob, err := s.repo.GetListRank(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
