package repo

import (
	"context"
	"doan/pkg/model"
	"gorm.io/gorm"
)

type PGInterface interface {
	GetRepo() *gorm.DB
	Transaction(ctx context.Context, f func(rp PGInterface) error) error
	DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc)

	// user
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
	GetOneUser(ctx context.Context, id string) (*model.User, error)

	// movie theater
	CreateMovieTheater(ctx context.Context, user *model.MovieTheater) error
	UpdateMovieTheater(ctx context.Context, user *model.MovieTheater) error
	DeleteMovieTheater(ctx context.Context, id string) error
	GetOneMovieTheater(ctx context.Context, id string) (*model.MovieTheater, error)

	// room
	CreateRoom(ctx context.Context, user *model.Room) error
	UpdateRoom(ctx context.Context, user *model.Room) error
	DeleteRoom(ctx context.Context, id string) error
	GetOneRoom(ctx context.Context, id string) (*model.Room, error)

	// seat
	CreateSeat(ctx context.Context, user *model.Seat) error
	UpdateSeat(ctx context.Context, user *model.Seat) error
	DeleteSeat(ctx context.Context, id string) error
	GetOneSeat(ctx context.Context, id string) (*model.Seat, error)

	// order
	CreateOrder(ctx context.Context, user *model.Order) error
	UpdateOrder(ctx context.Context, user *model.Order) error
	DeleteOrder(ctx context.Context, id string) error
	GetOneOrder(ctx context.Context, id string) (*model.Order, error)

	// ticket
	CreateTicket(ctx context.Context, user *model.Ticket) error
	UpdateTicket(ctx context.Context, user *model.Ticket) error
	DeleteTicket(ctx context.Context, id string) error
	GetOneTicket(ctx context.Context, id string) (*model.Ticket, error)

	// showtime
	CreateShowtime(ctx context.Context, user *model.Showtime) error
	UpdateShowtime(ctx context.Context, user *model.Showtime) error
	DeleteShowtime(ctx context.Context, id string) error
	GetOneShowtime(ctx context.Context, id string) (*model.Showtime, error)
}
