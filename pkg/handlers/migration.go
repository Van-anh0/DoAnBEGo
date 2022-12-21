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
		&model.User{},
		&model.Metadata{},
		&model.MovieTheater{},
		&model.Room{},
		&model.Movie{},
		&model.Seat{},
		&model.Showtime{},
		&model.Order{},
		&model.Ticket{},
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
	migrate := gormigrate.New(h.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20221212085554",
			Migrate: func(tx *gorm.DB) error {
				if err := h.BaseMigrate(ctx, tx); err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "20221216000202",
			Migrate: func(tx *gorm.DB) error {
				if err := h.db.AutoMigrate(&model.Order{}, &model.Ticket{}); err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "20221221143528",
			Migrate: func(tx *gorm.DB) error {
				if err := h.db.AutoMigrate(&model.Seat{}, &model.Showtime{}); err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "20221216091924",
			Migrate: func(tx *gorm.DB) error {
				if err := h.db.AutoMigrate(&model.Showtime{}); err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "20221216110030",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Exec(`
					ALTER TABLE orders DROP COLUMN slot_id;
				`).Error; err != nil {
					log.Warn(err)
				}
				return nil
			},
		},
		{
			ID: "20221221173006",
			Migrate: func(tx *gorm.DB) error {
				if err := h.db.AutoMigrate(&model.Showtime{}); err != nil {
					return err
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
