package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type UserRankService struct {
	repo repo.PGInterface
}

type UserRankInterface interface {
	Create(ctx context.Context, ob model.UserRankRequest) (rs *model.UserRank, err error)
	Update(ctx context.Context, ob model.UserRankRequest) (rs *model.UserRank, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.UserRank, err error)
	GetList(ctx context.Context, req model.UserRankParams) (rs *model.UserRankResponse, err error)
}

func NewUserRankService(repo repo.PGInterface) UserRankInterface {
	return &UserRankService{repo: repo}
}

func (s *UserRankService) Create(ctx context.Context, req model.UserRankRequest) (rs *model.UserRank, err error) {

	ob := &model.UserRank{}
	common.Sync(req, ob)

	if err := s.repo.CreateUserRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *UserRankService) Update(ctx context.Context, req model.UserRankRequest) (rs *model.UserRank, err error) {
	ob, err := s.repo.GetOneUserRank(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateUserRank(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *UserRankService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteUserRank(ctx, id)
}

func (s *UserRankService) GetOne(ctx context.Context, id string) (rs *model.UserRank, err error) {

	ob, err := s.repo.GetOneUserRank(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *UserRankService) GetList(ctx context.Context, req model.UserRankParams) (rs *model.UserRankResponse, err error) {

	ob, err := s.repo.GetListUserRank(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
