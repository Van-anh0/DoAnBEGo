package repo

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/utils"
	"strings"
)

func (r *RepoPG) CreateSeat(ctx context.Context, ob *model.Seat) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Create(ob).Error
}

func (r *RepoPG) UpdateSeat(ctx context.Context, ob *model.Seat) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", ob.ID).Updates(&ob).Error
}

func (r *RepoPG) DeleteSeat(ctx context.Context, id string) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", id).Delete(&model.Seat{}).Error
}

func (r *RepoPG) GetOneSeat(ctx context.Context, id string) (*model.Seat, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.Seat{}
	if err := tx.Where("id = ?", id).Find(&rs).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}

func (r *RepoPG) GetListSeat(ctx context.Context, req model.SeatParams) (*model.ListSeatResponse, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	rs := model.ListSeatResponse{}
	var err error
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	total := new(struct {
		Count int `json:"count"`
	})

	tx = tx.Table("seat")

	//if req.ShowtimeId != "" {
	//	tx = tx.Select("seat.*, CASE WHEN show_seat.showtime_id = ? THEN 'booked' ELSE 'available' END AS status", req.ShowtimeId)
	//	tx = tx.Joins("LEFT JOIN show_seat ON seat.id = show_seat.seat_id")
	//}

	if req.Search != "" {
		tx = tx.Where("unaccent(name) ilike %?%", req.Search)
	}

	if req.Filter != "" {
		filter := strings.Split(req.Filter, ",")
		for i := 0; i < len(filter); i += 2 {
			if i+1 < len(filter) {
				tx = tx.Where(filter[i]+" = ?", filter[i+1])
			}
		}
	}
	switch req.Sort {
	case utils.SORT_CREATED_AT_OLDEST:
		tx = tx.Order("created_at")
	case "name":
		tx = tx.Order("seat.name")
	case "-name":
		tx = tx.Order("seat.name desc")
	case "row":
		tx = tx.Order("seat.row")
	default:
		tx = tx.Order("seat.row,seat.col")
	}
	if err := tx.Find(&rs.Data).Error; err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	if rs.Meta, err = r.GetPaginationInfo("", tx, total.Count, page, pageSize); err != nil {
		return nil, r.ReturnErrorInGetFunc(ctx, err, utils.GetCurrentCaller(r, 0))
	}

	return &rs, nil
}
