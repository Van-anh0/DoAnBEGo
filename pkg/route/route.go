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
	repoPG := repo.NewPGRepo(db)
	if s.setting.DbDebugEnable {
		db = db.Debug()
	}

	userService := service2.NewUserService(repoPG)
	user := handlers.NewUserHandlers(userService)

	movieTheaterService := service2.NewMovieTheaterService(repoPG)
	movieTheater := handlers.NewMovieTheaterHandlers(movieTheaterService)

	if conf.GetEnv().EnvName == "dev" {
		s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1Api := s.Router.Group("/api/v1")

	// User
	v1Api.PUT("/user/create", ginext.WrapHandler(user.Create))
	v1Api.PUT("/user/update/:id", ginext.WrapHandler(user.Update))
	v1Api.DELETE("/user/delete/:id", ginext.WrapHandler(user.Delete))
	v1Api.GET("/user/get-one/:id", ginext.WrapHandler(user.GetOne))

	// Movie Theater
	v1Api.PUT("/movie-theater/create", ginext.WrapHandler(movieTheater.Create))
	v1Api.PUT("/movie-theater/update/:id", ginext.WrapHandler(movieTheater.Update))
	v1Api.DELETE("/movie-theater/delete/:id", ginext.WrapHandler(movieTheater.Delete))
	v1Api.GET("/movie-theater/get-one/:id", ginext.WrapHandler(movieTheater.GetOne))

	// Migrate
	migrateHandler := handlers.NewMigrationHandler(db)
	v1Api.POST("/internal/migrate", migrateHandler.Migrate)

	return s
}
