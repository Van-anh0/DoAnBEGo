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
	CreateTicket(ctx context.Context, user *model.OrderItem) error
	UpdateTicket(ctx context.Context, user *model.OrderItem) error
	DeleteTicket(ctx context.Context, id string) error
	GetOneTicket(ctx context.Context, id string) (*model.OrderItem, error)
	GetListTicket(ctx context.Context, req model.TicketParams) (*model.TicketResponse, error)
	CreateMultiTicket(ctx context.Context, ob *[]model.OrderItem) error

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

	// product
	CreateProduct(ctx context.Context, user *model.Product) error
	UpdateProduct(ctx context.Context, user *model.Product) error
	DeleteProduct(ctx context.Context, id string) error
	GetOneProduct(ctx context.Context, id string) (*model.Product, error)
	GetListProduct(ctx context.Context, req model.ProductParams) (*model.ProductResponse, error)

	// category
	CreateCategory(ctx context.Context, user *model.Category) error
	UpdateCategory(ctx context.Context, user *model.Category) error
	DeleteCategory(ctx context.Context, id string) error
	GetOneCategory(ctx context.Context, id string) (*model.Category, error)
	GetListCategory(ctx context.Context, req model.CategoryParams) (*model.CategoryResponse, error)

	// attribute
	CreateAttribute(ctx context.Context, user *model.Attribute) error
	UpdateAttribute(ctx context.Context, user *model.Attribute) error
	DeleteAttribute(ctx context.Context, id string) error
	GetOneAttribute(ctx context.Context, id string) (*model.Attribute, error)
	GetListAttribute(ctx context.Context, req model.AttributeParams) (*model.AttributeResponse, error)

	// comment
	CreateComment(ctx context.Context, user *model.Comment) error
	UpdateComment(ctx context.Context, user *model.Comment) error
	DeleteComment(ctx context.Context, id string) error
	GetOneComment(ctx context.Context, id string) (*model.Comment, error)
	GetListComment(ctx context.Context, req model.CommentParams) (*model.CommentResponse, error)

	// product_rank
	CreateProductRank(ctx context.Context, user *model.ProductRank) error
	UpdateProductRank(ctx context.Context, user *model.ProductRank) error
	DeleteProductRank(ctx context.Context, id string) error
	GetOneProductRank(ctx context.Context, id string) (*model.ProductRank, error)
	GetListProductRank(ctx context.Context, req model.ProductRankParams) (*model.ProductRankResponse, error)

	// promotion
	CreatePromotion(ctx context.Context, user *model.Promotion) error
	UpdatePromotion(ctx context.Context, user *model.Promotion) error
	DeletePromotion(ctx context.Context, id string) error
	GetOnePromotion(ctx context.Context, id string) (*model.Promotion, error)
	GetListPromotion(ctx context.Context, req model.PromotionParams) (*model.PromotionResponse, error)

	// user_rank
	CreateUserRank(ctx context.Context, user *model.UserRank) error
	UpdateUserRank(ctx context.Context, user *model.UserRank) error
	DeleteUserRank(ctx context.Context, id string) error
	GetOneUserRank(ctx context.Context, id string) (*model.UserRank, error)
	GetListUserRank(ctx context.Context, req model.UserRankParams) (*model.UserRankResponse, error)

	// category_has_product
	CreateCategoryHasProduct(ctx context.Context, user *model.CategoryHasProduct) error
	UpdateCategoryHasProduct(ctx context.Context, user *model.CategoryHasProduct) error
	DeleteCategoryHasProduct(ctx context.Context, id string) error
	GetOneCategoryHasProduct(ctx context.Context, id string) (*model.CategoryHasProduct, error)
	GetListCategoryHasProduct(ctx context.Context, req model.CategoryHasProductParams) (*model.CategoryHasProductResponse, error)
}
