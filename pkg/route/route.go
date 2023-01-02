package route

import (
	"doan/conf"
	"doan/pkg/handlers"
	"doan/pkg/repo"
	service2 "doan/pkg/service"
	"github.com/caarlos0/env/v6"
	"github.com/praslar/cloud0/ginext"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*conf.BaseApp
	setting *extraSetting
}

// NewService
// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api/v1
func NewService() *Service {
	s := &Service{
		conf.NewApp("doan", "v1.0"),
		&extraSetting{},
	}
	_ = env.Parse(s.setting)

	db := s.GetDB()
	if s.setting.DbDebugEnable {
		db = db.Debug()
	}
	repoPG := repo.NewPGRepo(db)

	userService := service2.NewUserService(repoPG)
	user := handlers.NewUserHandlers(userService)

	CinemaService := service2.NewCinemaService(repoPG)
	Cinema := handlers.NewCinemaHandlers(CinemaService)

	movieService := service2.NewMovieService(repoPG)
	movie := handlers.NewMovieHandlers(movieService)

	roomService := service2.NewRoomService(repoPG)
	room := handlers.NewRoomHandlers(roomService)

	seatService := service2.NewSeatService(repoPG)
	seat := handlers.NewSeatHandlers(seatService)

	showSeatService := service2.NewShowSeatService(repoPG)
	showSeat := handlers.NewShowSeatHandlers(showSeatService)

	showtimeService := service2.NewShowtimeService(repoPG)
	show := handlers.NewShowtimeHandlers(showtimeService)

	orderService := service2.NewOrderService(repoPG)
	order := handlers.NewOrderHandlers(orderService)

	orderItemService := service2.NewOrderItemService(repoPG)
	orderItem := handlers.NewOrderItemHandlers(orderItemService)

	metadataService := service2.NewMetadataService(repoPG)
	metadata := handlers.NewMetadataHandlers(metadataService)

	productService := service2.NewProductService(repoPG)
	product := handlers.NewProductHandlers(productService)

	categoryService := service2.NewCategoryService(repoPG)
	category := handlers.NewCategoryHandlers(categoryService)

	categoryHasProductService := service2.NewCategoryHasProductService(repoPG)
	chp := handlers.NewCategoryHasProductHandlers(categoryHasProductService)

	commentService := service2.NewCommentService(repoPG)
	comment := handlers.NewCommentHandlers(commentService)

	promotionService := service2.NewPromotionService(repoPG)
	promotion := handlers.NewPromotionHandlers(promotionService)

	rankService := service2.NewRankService(repoPG)
	rank := handlers.NewRankHandlers(rankService)

	MovieRankService := service2.NewMovieRankService(repoPG)
	MovieRank := handlers.NewMovieRankHandlers(MovieRankService)

	if conf.GetEnv().EnvName == "dev" {
		s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1Api := s.Router.Group("/api/v1")

	// Auth
	v1Api.POST("/auth/login", ginext.WrapHandler(user.Login))
	v1Api.POST("/auth/register", ginext.WrapHandler(user.Register))

	// User
	v1Api.POST("/user/create", ginext.WrapHandler(user.Create))
	v1Api.PUT("/user/update/:id", ginext.WrapHandler(user.Update))
	v1Api.DELETE("/user/delete/:id", ginext.WrapHandler(user.Delete))
	v1Api.GET("/user/get-one/:id", ginext.WrapHandler(user.GetOne))
	v1Api.GET("/user/get-list", ginext.WrapHandler(user.GetList))

	// Movie Theater
	v1Api.POST("/cinema/create", ginext.WrapHandler(Cinema.Create))
	v1Api.PUT("/cinema/update/:id", ginext.WrapHandler(Cinema.Update))
	v1Api.DELETE("/cinema/delete/:id", ginext.WrapHandler(Cinema.Delete))
	v1Api.GET("/cinema/get-one/:id", ginext.WrapHandler(Cinema.GetOne))
	v1Api.GET("/cinema/get-list", ginext.WrapHandler(Cinema.GetList))

	// Movie
	v1Api.POST("/movie/create", ginext.WrapHandler(movie.Create))
	v1Api.PUT("/movie/update/:id", ginext.WrapHandler(movie.Update))
	v1Api.DELETE("/movie/delete/:id", ginext.WrapHandler(movie.Delete))
	v1Api.GET("/movie/get-one/:id", ginext.WrapHandler(movie.GetOne))
	v1Api.GET("/movie/get-list", ginext.WrapHandler(movie.GetList))

	// Room
	v1Api.POST("/room/create", ginext.WrapHandler(room.Create))
	v1Api.PUT("/room/update/:id", ginext.WrapHandler(room.Update))
	v1Api.DELETE("/room/delete/:id", ginext.WrapHandler(room.Delete))
	v1Api.GET("/room/get-one/:id", ginext.WrapHandler(room.GetOne))
	v1Api.GET("/room/get-list", ginext.WrapHandler(room.GetList))

	// Seat
	v1Api.POST("/seat/create", ginext.WrapHandler(seat.Create))
	v1Api.PUT("/seat/update/:id", ginext.WrapHandler(seat.Update))
	v1Api.DELETE("/seat/delete/:id", ginext.WrapHandler(seat.Delete))
	v1Api.GET("/seat/get-one/:id", ginext.WrapHandler(seat.GetOne))
	v1Api.GET("/seat/get-list", ginext.WrapHandler(seat.GetList))

	// Showtime
	v1Api.POST("/show/create", ginext.WrapHandler(show.Create))
	v1Api.PUT("/show/update/:id", ginext.WrapHandler(show.Update))
	v1Api.DELETE("/show/delete/:id", ginext.WrapHandler(show.Delete))
	v1Api.GET("/show/get-one/:id", ginext.WrapHandler(show.GetOne))
	v1Api.GET("/show/get-list", ginext.WrapHandler(show.GetList))
	v1Api.GET("/show/get-list-group", ginext.WrapHandler(show.GetListGroup))

	// Showtime
	v1Api.POST("/show-seat/create", ginext.WrapHandler(showSeat.Create))
	v1Api.PUT("/show-seat/update/:id", ginext.WrapHandler(showSeat.Update))
	v1Api.DELETE("/show-seat/delete/:id", ginext.WrapHandler(showSeat.Delete))
	v1Api.GET("/show-seat/get-one/:id", ginext.WrapHandler(showSeat.GetOne))
	v1Api.GET("/show-seat/get-list", ginext.WrapHandler(showSeat.GetList))

	// Metadata
	v1Api.POST("/metadata/create", ginext.WrapHandler(metadata.Create))
	v1Api.PUT("/metadata/update/:id", ginext.WrapHandler(metadata.Update))
	v1Api.DELETE("/metadata/delete/:id", ginext.WrapHandler(metadata.Delete))
	v1Api.GET("/metadata/get-one/:id", ginext.WrapHandler(metadata.GetOne))
	v1Api.GET("/metadata/get-list", ginext.WrapHandler(metadata.GetList))

	// Order
	v1Api.POST("/order/create", ginext.WrapHandler(order.Create))
	v1Api.PUT("/order/update/:id", ginext.WrapHandler(order.Update))
	v1Api.DELETE("/order/delete/:id", ginext.WrapHandler(order.Delete))
	v1Api.GET("/order/get-one/:id", ginext.WrapHandler(order.GetOne))
	v1Api.GET("/order/get-list", ginext.WrapHandler(order.GetList))

	// orderItem
	v1Api.POST("/order-item/create", ginext.WrapHandler(orderItem.Create))
	v1Api.PUT("/order-item/update/:id", ginext.WrapHandler(orderItem.Update))
	v1Api.DELETE("/order-item/delete/:id", ginext.WrapHandler(orderItem.Delete))
	v1Api.GET("/order-item/get-one/:id", ginext.WrapHandler(orderItem.GetOne))
	v1Api.GET("/order-item/get-list", ginext.WrapHandler(orderItem.GetList))

	// Product
	v1Api.POST("/product/create", ginext.WrapHandler(product.Create))
	v1Api.PUT("/product/update/:id", ginext.WrapHandler(product.Update))
	v1Api.DELETE("/product/delete/:id", ginext.WrapHandler(product.Delete))
	v1Api.GET("/product/get-one/:id", ginext.WrapHandler(product.GetOne))
	v1Api.GET("/product/get-list", ginext.WrapHandler(product.GetList))

	// category
	v1Api.POST("/category/create", ginext.WrapHandler(category.Create))
	v1Api.PUT("/category/update/:id", ginext.WrapHandler(category.Update))
	v1Api.DELETE("/category/delete/:id", ginext.WrapHandler(category.Delete))
	v1Api.GET("/category/get-one/:id", ginext.WrapHandler(category.GetOne))
	v1Api.GET("/category/get-list", ginext.WrapHandler(category.GetList))

	// chp
	v1Api.POST("/category-has-product/create", ginext.WrapHandler(chp.Create))
	v1Api.PUT("/category-has-product/update/:id", ginext.WrapHandler(chp.Update))
	v1Api.DELETE("/category-has-product/delete/:id", ginext.WrapHandler(chp.Delete))
	v1Api.GET("/category-has-product/get-one/:id", ginext.WrapHandler(chp.GetOne))
	v1Api.GET("/category-has-product/get-list", ginext.WrapHandler(chp.GetList))

	// comment
	v1Api.POST("/comment/create", ginext.WrapHandler(comment.Create))
	v1Api.PUT("/comment/update/:id", ginext.WrapHandler(comment.Update))
	v1Api.DELETE("/comment/delete/:id", ginext.WrapHandler(comment.Delete))
	v1Api.GET("/comment/get-one/:id", ginext.WrapHandler(comment.GetOne))
	v1Api.GET("/comment/get-list", ginext.WrapHandler(comment.GetList))

	// promotion
	v1Api.POST("/promotion/create", ginext.WrapHandler(promotion.Create))
	v1Api.PUT("/promotion/update/:id", ginext.WrapHandler(promotion.Update))
	v1Api.DELETE("/promotion/delete/:id", ginext.WrapHandler(promotion.Delete))
	v1Api.GET("/promotion/get-one/:id", ginext.WrapHandler(promotion.GetOne))
	v1Api.GET("/promotion/get-list", ginext.WrapHandler(promotion.GetList))

	// user_rank
	v1Api.POST("/rank/create", ginext.WrapHandler(rank.Create))
	v1Api.PUT("/rank/update/:id", ginext.WrapHandler(rank.Update))
	v1Api.DELETE("/rank/delete/:id", ginext.WrapHandler(rank.Delete))
	v1Api.GET("/rank/get-one/:id", ginext.WrapHandler(rank.GetOne))
	v1Api.GET("/rank/get-list", ginext.WrapHandler(rank.GetList))

	// product_rank
	v1Api.POST("/product-rank/create", ginext.WrapHandler(MovieRank.Create))
	v1Api.PUT("/product-rank/update/:id", ginext.WrapHandler(MovieRank.Update))
	v1Api.DELETE("/product-rank/delete/:id", ginext.WrapHandler(MovieRank.Delete))
	v1Api.GET("/product-rank/get-one/:id", ginext.WrapHandler(MovieRank.GetOne))
	v1Api.GET("/product-rank/get-list", ginext.WrapHandler(MovieRank.GetList))

	// Migrate
	migrateHandler := handlers.NewMigrationHandler(db)
	v1Api.POST("/internal/migrate", migrateHandler.Migrate)

	return s
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
