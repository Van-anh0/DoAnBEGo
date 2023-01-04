package handlers

import (
	"doan/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/praslar/cloud0/logger"
	"gorm.io/gorm"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{db: db}
}

func (h *MigrationHandler) BaseMigrate(ctx *gin.Context, tx *gorm.DB) error {
	//log := logger.WithCtx(ctx, "BaseMigrate")
	//if err := tx.Exec(`
	//		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	//	`).Error; err != nil {
	//	log.Errorf(err.Error())
	//}

	models := []interface{}{
		&model.Metadata{},
		&model.User{},
		&model.Cinema{},
		&model.Room{},
		&model.Seat{},
		&model.Movie{},
		&model.Showtime{},
		&model.ShowSeat{},
		&model.Order{},
		&model.OrderItem{},
		&model.Product{},
		&model.UserRank{},
		&model.MovieRank{},
		&model.Promotion{},
		&model.Category{},
		&model.CategoryHasProduct{},
		&model.MovieComment{},
	}

	for _, m := range models {
		err := h.db.AutoMigrate(m)
		if err != nil {
			_ = ctx.Error(err)
		}
	}

	//if err := tx.Exec(`
	//	ALTER TABLE answer ADD CONSTRAINT uc_answer_key UNIQUE (question_id, content, type);
	//`).Error; err != nil {
	//	log.Warn(err)
	//}

	return nil
}

func (h *MigrationHandler) Migrate(ctx *gin.Context) {
	log := logger.WithCtx(ctx, "Migrate")
	// put your migrations at the end of the list
	migrate := gormigrate.New(h.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20221229163504",
			Migrate: func(tx *gorm.DB) error {
				log.Info("Migrate 20221229163504 - BaseMigrate")
				if err := h.BaseMigrate(ctx, tx); err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "20230102002023",
			Migrate: func(tx *gorm.DB) error {
				log.Info("Migrate 20230102002023 - BaseMigrate")
				err := h.db.AutoMigrate(&model.Order{})
				if err != nil {
					_ = ctx.Error(err)
				}
				return nil
			},
		},
		{
			ID: "20230103211916",
			Migrate: func(tx *gorm.DB) error {
				log.Info("Migrate 20230103211916 - Add column ticket")
				err := h.db.AutoMigrate(&model.Movie{})
				if err != nil {
					_ = ctx.Error(err)
				}
				return nil
			},
		},
	})
	err := migrate.Migrate()
	if err != nil {
		log.Errorf(err.Error())
	}
}
