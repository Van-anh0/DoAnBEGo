package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/praslar/lib/common"
)

type RoomService struct {
	repo repo.PGInterface
}

type RoomInterface interface {
	Create(ctx context.Context, ob model.RoomRequest) (rs *model.Room, err error)
	Update(ctx context.Context, ob model.RoomRequest) (rs *model.Room, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Room, err error)
	GetList(ctx context.Context, req model.RoomParams) (rs *model.RoomResponse, err error)
}

func NewRoomService(repo repo.PGInterface) RoomInterface {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(ctx context.Context, req model.RoomRequest) (rs *model.Room, err error) {

	ob := &model.Room{}
	common.Sync(req, ob)

	if err := s.repo.CreateRoom(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *RoomService) Update(ctx context.Context, req model.RoomRequest) (rs *model.Room, err error) {

	ob := &model.Room{}
	common.Sync(req, ob)

	if err := s.repo.UpdateRoom(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *RoomService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteRoom(ctx, id)
}

func (s *RoomService) GetOne(ctx context.Context, id string) (rs *model.Room, err error) {

	ob, err := s.repo.GetOneRoom(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *RoomService) GetList(ctx context.Context, req model.RoomParams) (rs *model.RoomResponse, err error) {

	ob, err := s.repo.GetListRoom(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
