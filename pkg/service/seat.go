package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/utils"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type SeatService struct {
	repo repo.PGInterface
}

type SeatInterface interface {
	Create(ctx context.Context, ob model.SeatRequest) (rs *model.Seat, err error)
	Update(ctx context.Context, ob model.SeatRequest) (rs *model.Seat, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Seat, err error)
	GetList(ctx context.Context, req model.SeatParams) (rs *model.ListSeatResponse, err error)
	AdminGetList(ctx context.Context, req model.SeatParams) (rs *model.ListSeatResponse, err error)
}

func NewSeatService(repo repo.PGInterface) SeatInterface {
	return &SeatService{repo: repo}
}

func (s *SeatService) Create(ctx context.Context, req model.SeatRequest) (rs *model.Seat, err error) {

	ob := &model.Seat{}
	common.Sync(req, ob)

	if err := s.repo.CreateSeat(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SeatService) Update(ctx context.Context, req model.SeatRequest) (rs *model.Seat, err error) {
	ob, err := s.repo.GetOneSeat(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateSeat(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SeatService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteSeat(ctx, id)
}

func (s *SeatService) GetOne(ctx context.Context, id string) (rs *model.Seat, err error) {

	ob, err := s.repo.GetOneSeat(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *SeatService) GetList(ctx context.Context, req model.SeatParams) (rs *model.ListSeatResponse, err error) {

	ob, err := s.repo.GetListSeat(ctx, req)
	if err != nil {
		return nil, err
	}

	listId := make([]string, 0)
	for _, v := range ob.Data {
		listId = append(listId, v.ID)
	}

	// Get list show seat
	showSeat, err := s.repo.GetListShowSeat(ctx, model.ShowSeatParams{
		SeatId:     listId,
		ShowtimeId: req.ShowtimeId,
	})
	if err != nil {
		return nil, err
	}

	// combine ob and showSeat to get status with map
	mapShowSeat := make(map[string]string)
	for _, v := range showSeat.Data {
		mapShowSeat[v.SeatId] = utils.SEAT_BOOKED
	}
	for i, v := range ob.Data {
		if _, ok := mapShowSeat[v.ID]; ok {
			ob.Data[i].Status = mapShowSeat[v.ID]
		} else {
			ob.Data[i].Status = utils.SEAT_AVAILABLE
		}
	}

	return ob, nil
}

func (s *SeatService) AdminGetList(ctx context.Context, req model.SeatParams) (rs *model.ListSeatResponse, err error) {

	ob, err := s.repo.AdminGetListSeat(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
