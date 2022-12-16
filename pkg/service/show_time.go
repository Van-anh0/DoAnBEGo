package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"github.com/praslar/lib/common"
)

type ShowtimeService struct {
	repo repo.PGInterface
}

type ShowtimeInterface interface {
	Create(ctx context.Context, ob model.ShowtimeRequest) (rs *model.Showtime, err error)
	Update(ctx context.Context, ob model.ShowtimeRequest) (rs *model.Showtime, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Showtime, err error)
	GetList(ctx context.Context, req model.ShowtimeParams) (rs *model.ShowtimeResponse, err error)
	GetListGroupByDay(ctx context.Context, req model.ShowtimeParams) (rs *model.ShowtimeGroupResponse, err error)
}

func NewShowtimeService(repo repo.PGInterface) ShowtimeInterface {
	return &ShowtimeService{repo: repo}
}

func (s *ShowtimeService) Create(ctx context.Context, req model.ShowtimeRequest) (rs *model.Showtime, err error) {

	ob := &model.Showtime{}
	common.Sync(req, ob)

	if err := s.repo.CreateShowtime(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ShowtimeService) Update(ctx context.Context, req model.ShowtimeRequest) (rs *model.Showtime, err error) {

	ob := &model.Showtime{}
	common.Sync(req, ob)

	if err := s.repo.UpdateShowtime(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ShowtimeService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteShowtime(ctx, id)
}

func (s *ShowtimeService) GetOne(ctx context.Context, id string) (rs *model.Showtime, err error) {

	ob, err := s.repo.GetOneShowtime(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ShowtimeService) GetList(ctx context.Context, req model.ShowtimeParams) (rs *model.ShowtimeResponse, err error) {

	ob, err := s.repo.GetListShowtime(ctx, req)
	if err != nil {
		return nil, err
	}

	// change list to object group by day
	var listShowtime = make(map[string][]model.Showtime)
	for _, v := range ob.Data {
		listShowtime[v.Day.String()] = append(listShowtime[v.Day.String()], v)
	}
	return ob, nil
}

func (s *ShowtimeService) GetListGroupByDay(ctx context.Context, req model.ShowtimeParams) (rs *model.ShowtimeGroupResponse, err error) {

	ob, err := s.repo.GetListShowtime(ctx, req)
	if err != nil {
		return nil, err
	}

	// change list to object group by day
	var listShowtime = make(map[string][]model.Showtime)
	for _, v := range ob.Data {
		listShowtime[v.Day.String()] = append(listShowtime[v.Day.String()], v)
	}

	var listShowtimeGroup model.ShowtimeGroupResponse
	listShowtimeGroup.Data = listShowtime
	//listShowtimeGroup.Meta = ob.Meta

	return &listShowtimeGroup, nil
}
