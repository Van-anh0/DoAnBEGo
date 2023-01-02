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
	GetOneUserByEmail(ctx context.Context, email string) (*model.User, error)

	// movie theater
	CreateCinema(ctx context.Context, user *model.Cinema) error
	UpdateCinema(ctx context.Context, user *model.Cinema) error
	DeleteCinema(ctx context.Context, id string) error
	GetOneCinema(ctx context.Context, id string) (*model.Cinema, error)
	GetListCinema(ctx context.Context, req model.CinemaParams) (*model.CinemaResponse, error)

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
	GetListSeat(ctx context.Context, req model.SeatParams) (*model.ListSeatResponse, error)

	// order
	CreateOrder(ctx context.Context, user *model.Order) error
	UpdateOrder(ctx context.Context, user *model.Order) error
	DeleteOrder(ctx context.Context, id string) error
	GetOneOrder(ctx context.Context, id string) (*model.Order, error)
	GetListOrder(ctx context.Context, req model.OrderParams) (*model.OrderResponse, error)

	// OrderItem
	CreateOrderItem(ctx context.Context, user *model.OrderItem) error
	UpdateOrderItem(ctx context.Context, user *model.OrderItem) error
	DeleteOrderItem(ctx context.Context, id string) error
	GetOneOrderItem(ctx context.Context, id string) (*model.OrderItem, error)
	GetListOrderItem(ctx context.Context, req model.OrderItemParams) (*model.OrderItemResponse, error)
	CreateMultiOrderItem(ctx context.Context, ob *[]model.OrderItem) error

	// showtime
	CreateShowtime(ctx context.Context, user *model.Showtime) error
	UpdateShowtime(ctx context.Context, user *model.Showtime) error
	DeleteShowtime(ctx context.Context, id string) error
	GetOneShowtime(ctx context.Context, id string) (*model.Showtime, error)
	GetListShowtime(ctx context.Context, req model.ShowParams) (*model.ShowtimeResponse, error)

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

	// product
	CreateProduct(ctx context.Context, user *model.Product) error
	UpdateProduct(ctx context.Context, user *model.Product) error
	DeleteProduct(ctx context.Context, id string) error
	GetOneProduct(ctx context.Context, id string) (*model.Product, error)
	GetListProduct(ctx context.Context, req model.ProductParams) (*model.ListProductResponse, error)

	// category
	CreateCategory(ctx context.Context, user *model.Category) error
	UpdateCategory(ctx context.Context, user *model.Category) error
	DeleteCategory(ctx context.Context, id string) error
	GetOneCategory(ctx context.Context, id string) (*model.Category, error)
	GetListCategory(ctx context.Context, req model.CategoryParams) (*model.CategoryResponse, error)

	// comment
	CreateComment(ctx context.Context, user *model.MovieComment) error
	UpdateComment(ctx context.Context, user *model.MovieComment) error
	DeleteComment(ctx context.Context, id string) error
	GetOneComment(ctx context.Context, id string) (*model.MovieComment, error)
	GetListComment(ctx context.Context, req model.MovieCommentParams) (*model.MovieCommentResponse, error)

	// product_rank
	CreateMovieRank(ctx context.Context, user *model.MovieRank) error
	UpdateMovieRank(ctx context.Context, user *model.MovieRank) error
	DeleteMovieRank(ctx context.Context, id string) error
	GetOneMovieRank(ctx context.Context, id string) (*model.MovieRank, error)
	GetListMovieRank(ctx context.Context, req model.MovieRankParams) (*model.MovieRankResponse, error)

	// promotion
	CreatePromotion(ctx context.Context, user *model.Promotion) error
	UpdatePromotion(ctx context.Context, user *model.Promotion) error
	DeletePromotion(ctx context.Context, id string) error
	GetOnePromotion(ctx context.Context, id string) (*model.Promotion, error)
	GetListPromotion(ctx context.Context, req model.PromotionParams) (*model.PromotionResponse, error)

	// user_rank
	CreateRank(ctx context.Context, user *model.UserRank) error
	UpdateRank(ctx context.Context, user *model.UserRank) error
	DeleteRank(ctx context.Context, id string) error
	GetOneRank(ctx context.Context, id string) (*model.UserRank, error)
	GetListRank(ctx context.Context, req model.RankParams) (*model.RankResponse, error)

	// category_has_product
	CreateCategoryHasProduct(ctx context.Context, user *model.CategoryHasProduct) error
	UpdateCategoryHasProduct(ctx context.Context, user *model.CategoryHasProduct) error
	DeleteCategoryHasProduct(ctx context.Context, id string) error
	GetOneCategoryHasProduct(ctx context.Context, id string) (*model.CategoryHasProduct, error)
	GetListCategoryHasProduct(ctx context.Context, req model.CategoryHasProductParams) (*model.CategoryHasProductResponse, error)

	// ShowSeat
	CreateShowSeat(ctx context.Context, user *model.ShowSeat) error
	UpdateShowSeat(ctx context.Context, user *model.ShowSeat) error
	DeleteShowSeat(ctx context.Context, id string) error
	GetOneShowSeat(ctx context.Context, id string) (*model.ShowSeat, error)
	GetListShowSeat(ctx context.Context, req model.ShowSeatParams) (*model.ShowSeatResponse, error)
	CreateMultiShowSeat(ctx context.Context, ob *[]model.ShowSeat) error
}
