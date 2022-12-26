package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type CommentService struct {
	repo repo.PGInterface
}

type CommentInterface interface {
	Create(ctx context.Context, ob model.CommentRequest) (rs *model.Comment, err error)
	Update(ctx context.Context, ob model.CommentRequest) (rs *model.Comment, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Comment, err error)
	GetList(ctx context.Context, req model.CommentParams) (rs *model.CommentResponse, err error)
}

func NewCommentService(repo repo.PGInterface) CommentInterface {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(ctx context.Context, req model.CommentRequest) (rs *model.Comment, err error) {

	ob := &model.Comment{}
	common.Sync(req, ob)

	if err := s.repo.CreateComment(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CommentService) Update(ctx context.Context, req model.CommentRequest) (rs *model.Comment, err error) {
	ob, err := s.repo.GetOneComment(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateComment(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CommentService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteComment(ctx, id)
}

func (s *CommentService) GetOne(ctx context.Context, id string) (rs *model.Comment, err error) {

	ob, err := s.repo.GetOneComment(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CommentService) GetList(ctx context.Context, req model.CommentParams) (rs *model.CommentResponse, err error) {

	ob, err := s.repo.GetListComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
