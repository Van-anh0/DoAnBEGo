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
	Create(ctx context.Context, ob model.MovieCommentRequest) (rs *model.MovieComment, err error)
	Update(ctx context.Context, ob model.MovieCommentRequest) (rs *model.MovieComment, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.MovieComment, err error)
	GetList(ctx context.Context, req model.MovieCommentParams) (rs *model.MovieCommentResponse, err error)
}

func NewCommentService(repo repo.PGInterface) CommentInterface {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(ctx context.Context, req model.MovieCommentRequest) (rs *model.MovieComment, err error) {

	ob := &model.MovieComment{}
	common.Sync(req, ob)

	if err := s.repo.CreateComment(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CommentService) Update(ctx context.Context, req model.MovieCommentRequest) (rs *model.MovieComment, err error) {
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

func (s *CommentService) GetOne(ctx context.Context, id string) (rs *model.MovieComment, err error) {

	ob, err := s.repo.GetOneComment(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *CommentService) GetList(ctx context.Context, req model.MovieCommentParams) (rs *model.MovieCommentResponse, err error) {

	ob, err := s.repo.GetListComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
