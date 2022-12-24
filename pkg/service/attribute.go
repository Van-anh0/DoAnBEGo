package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type AttributeService struct {
	repo repo.PGInterface
}

type AttributeInterface interface {
	Create(ctx context.Context, ob model.AttributeRequest) (rs *model.Attribute, err error)
	Update(ctx context.Context, ob model.AttributeRequest) (rs *model.Attribute, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Attribute, err error)
	GetList(ctx context.Context, req model.AttributeParams) (rs *model.AttributeResponse, err error)
}

func NewAttributeService(repo repo.PGInterface) AttributeInterface {
	return &AttributeService{repo: repo}
}

func (s *AttributeService) Create(ctx context.Context, req model.AttributeRequest) (rs *model.Attribute, err error) {

	ob := &model.Attribute{}
	common.Sync(req, ob)

	if err := s.repo.CreateAttribute(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *AttributeService) Update(ctx context.Context, req model.AttributeRequest) (rs *model.Attribute, err error) {
	ob, err := s.repo.GetOneAttribute(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateAttribute(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *AttributeService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteAttribute(ctx, id)
}

func (s *AttributeService) GetOne(ctx context.Context, id string) (rs *model.Attribute, err error) {

	ob, err := s.repo.GetOneAttribute(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *AttributeService) GetList(ctx context.Context, req model.AttributeParams) (rs *model.AttributeResponse, err error) {

	ob, err := s.repo.GetListAttribute(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
