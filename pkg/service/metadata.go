package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/praslar/lib/common"
)

type MetadataService struct {
	repo repo.PGInterface
}

type MetadataInterface interface {
	Create(ctx context.Context, ob model.MetadataRequest) (rs *model.Metadata, err error)
	Update(ctx context.Context, ob model.MetadataRequest) (rs *model.Metadata, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Metadata, err error)
}

func NewMetadataService(repo repo.PGInterface) MetadataInterface {
	return &MetadataService{repo: repo}
}

func (s *MetadataService) Create(ctx context.Context, req model.MetadataRequest) (rs *model.Metadata, err error) {

	ob := &model.Metadata{}
	common.Sync(req, ob)

	if err := s.repo.CreateMetadata(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MetadataService) Update(ctx context.Context, req model.MetadataRequest) (rs *model.Metadata, err error) {

	ob := &model.Metadata{}
	common.Sync(req, ob)

	if err := s.repo.UpdateMetadata(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *MetadataService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteMetadata(ctx, id)
}

func (s *MetadataService) GetOne(ctx context.Context, id string) (rs *model.Metadata, err error) {

	ob, err := s.repo.GetOneMetadata(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
