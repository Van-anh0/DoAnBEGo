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
	GetListUser(ctx context.Context, req model.UserParams) (*model.UserResponse, error)

	// movie theater
	CreateMovieTheater(ctx context.Context, user *model.MovieTheater) error
	UpdateMovieTheater(ctx context.Context, user *model.MovieTheater) error
	DeleteMovieTheater(ctx context.Context, id string) error
	GetOneMovieTheater(ctx context.Context, id string) (*model.MovieTheater, error)
	GetListMovieTheater(ctx context.Context, req model.MovieTheaterParams) (*model.MovieTheaterResponse, error)

	// room
	CreateRoom(ctx context.Context, user *model.Room) error
	UpdateRoom(ctx context.Context, user *model.Room) error
	DeleteRoom(ctx context.Context, id string) error
	GetOneRoom(ctx context.Context, id string) (*model.Room, error)
	GetListRoom(ctx context.Context, req model.RoomParams) (*model.RoomResponse, error)

	// seat
	CreateSeat(ctx context.Context, user *model.Seat) error
	UpdateSeat(ctx context.Context, user *model.Seat) error
	DeleteSeat(ctx context.Context, id string) error
	GetOneSeat(ctx context.Context, id string) (*model.Seat, error)
	GetListSeat(ctx context.Context, req model.SeatParams) (*model.SeatResponse, error)

	// order
	CreateOrder(ctx context.Context, user *model.Order) error
	UpdateOrder(ctx context.Context, user *model.Order) error
	DeleteOrder(ctx context.Context, id string) error
	GetOneOrder(ctx context.Context, id string) (*model.Order, error)
	GetListOrder(ctx context.Context, req model.OrderParams) (*model.OrderResponse, error)

	// ticket
	CreateTicket(ctx context.Context, user *model.Ticket) error
	UpdateTicket(ctx context.Context, user *model.Ticket) error
	DeleteTicket(ctx context.Context, id string) error
	GetOneTicket(ctx context.Context, id string) (*model.Ticket, error)
	GetListTicket(ctx context.Context, req model.TicketParams) (*model.TicketResponse, error)

	// showtime
	CreateShowtime(ctx context.Context, user *model.Showtime) error
	UpdateShowtime(ctx context.Context, user *model.Showtime) error
	DeleteShowtime(ctx context.Context, id string) error
	GetOneShowtime(ctx context.Context, id string) (*model.Showtime, error)
	GetListShowtime(ctx context.Context, req model.ShowtimeParams) (*model.ShowtimeResponse, error)

	// movie
	CreateMovie(ctx context.Context, user *model.Movie) error
	UpdateMovie(ctx context.Context, user *model.Movie) error
	DeleteMovie(ctx context.Context, id string) error
	GetOneMovie(ctx context.Context, id string) (*model.Movie, error)
	GetListMovie(ctx context.Context, req model.MovieParams) (*model.MovieResponse, error)

	// metadata
	CreateMetadata(ctx context.Context, user *model.Metadata) error
	UpdateMetadata(ctx context.Context, user *model.Metadata) error
	DeleteMetadata(ctx context.Context, id string) error
	GetOneMetadata(ctx context.Context, id string) (*model.Metadata, error)
	GetListMetadata(ctx context.Context, req model.MetadataParams) (*model.MetadataResponse, error)
}
