package handlers

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/dgraph-io/dgo/v200"
//	"github.com/dgraph-io/dgo/v200/protos/api"
//	"github.com/gin-gonic/gin"
//	"github.com/go-gormigrate/gormigrate/v2"
//	"github.com/jinzhu/gorm/dialects/postgres"
//	"github.com/praslar/cloud0/logger"
//	"gorm.io/gorm"
//	"gorm.io/gorm/clause"
//	"io/ioutil"
//	"joon/pkg/model"
//	"joon/pkg/utils"
//	"strings"
//)
//
//type MigrationHandler struct {
//	db     *gorm.DB
//	dgraph *dgo.Dgraph
//}
//
//func NewMigrationHandler(db *gorm.DB, dgraph *dgo.Dgraph) *MigrationHandler {
//	return &MigrationHandler{db: db, dgraph: dgraph}
//}
//
//func (h *MigrationHandler) BaseMigrate(ctx *gin.Context, tx *gorm.DB) error {
//	log := logger.WithCtx(ctx, "BaseMigrate")
//	if err := tx.Exec(`
//			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
//			DROP type IF EXISTS cvs_status; Create TYPE cvs_status AS ENUM ('match', 'meet', 'hidden','like','waiting', 'unmatch', 'dislike');
//			DROP type IF EXISTS message_type; Create TYPE message_type AS ENUM ('text', 'image', 'video', 'audio', 'post', 'match');
//			DROP type IF EXISTS message_state; Create TYPE message_state AS ENUM ('ready', 'seen');
//			DROP type IF EXISTS user_post_category; Create TYPE user_post_category AS ENUM ('moment', 'prompt', 'profile_answers', 'selfie');
//			DROP type IF EXISTS transaction_status; Create TYPE transaction_status AS ENUM ('success', 'failed', 'cancelled');
//		`).Error; err != nil {
//		log.Errorf(err.Error())
//	}
//
//	models := []interface{}{
//		&model.User{},
//		&model.Otp{},
//		&model.Metadata{},
//		&model.UserAnswer{},
//		&model.Question{},
//		&model.Answer{},
//		&model.BlockAndReport{},
//		&model.Blacklist{},
//		&model.UserSocial{},
//		&model.UserPost{},
//		&model.Media{},
//		&model.NotifyRule{},
//		&model.NotifyInfo{},
//		&model.UserToken{},
//		&model.RefreshToken{},
//		&model.Dislike{},
//		&model.Role{},
//		&model.Transaction{},
//		&model.Screen{},
//		&model.UserLockTime{},
//		// V2
//		&model.Conversation{},
//		&model.Participant{},
//		&model.Message{},
//		&model.Call{},
//		&model.UserPoint{},
//		&model.Mission{},
//	}
//
//	for _, m := range models {
//		err := h.db.AutoMigrate(m)
//		if err != nil {
//			_ = ctx.Error(err)
//		}
//	}
//
//	if err := tx.Exec(`
//		ALTER TABLE conversation ADD CONSTRAINT uc_key UNIQUE (creator_id, user_id);
//		ALTER TABLE user_answer ADD CONSTRAINT uc_question_user_key UNIQUE (question_id, user_id);
//		ALTER TABLE participant ADD CONSTRAINT uc_conversation_user_key UNIQUE (conversation_id, user_id);
//		ALTER TABLE answer ADD CONSTRAINT uc_answer_key UNIQUE (question_id, content, type);
//	`).Error; err != nil {
//		log.Warn(err)
//	}
//
//	return nil
//}
//
//func (h *MigrationHandler) Migrate(ctx *gin.Context) {
//	log := logger.WithCtx(ctx, "BaseMigrate")
//	migrate := gormigrate.New(h.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
//		{
//			ID: "20221128213038",
//			Migrate: func(tx *gorm.DB) error {
//				if err := h.db.AutoMigrate(&model.Track{}); err != nil {
//					return err
//				}
//				return nil
//			},
//		},
//		{
//			ID: "20221128150251",
//			Migrate: func(tx *gorm.DB) error {
//				if err := h.db.AutoMigrate(&model.Link{}); err != nil {
//					return err
//				}
//				return nil
//			},
//		},
//		{
//			ID: "20221122224447",
//			Migrate: func(tx *gorm.DB) error {
//				if err := h.db.AutoMigrate(&model.WhitelistMeta{}); err != nil {
//					return err
//				}
//				return nil
//			},
//		},
//		{
//			ID: "20221114080852",
//			Migrate: func(tx *gorm.DB) error {
//				log.Info("20221114080852 - add column user_id in noti_history table")
//				if err := h.db.AutoMigrate(&model.NotiHistory{}); err != nil {
//					return err
//				}
//				return nil
//			},
//		},
//		{
//			ID: "20221111175610",
//			Migrate: func(tx *gorm.DB) error {
//				log.Info("20221111175610 - add column consumer in noti_history table")
//				if err := h.db.AutoMigrate(&model.NotiHistory{}); err != nil {
//					return err
//				}
//				return nil
//			},
//		},
//		{
//			ID: "20221108231314",
//			Migrate: func(tx *gorm.DB) error {
//				if err := h.db.AutoMigrate(&model.NotiHistory{}); err != nil {
//					return err
//				}
//				return nil
//			},
//		}, {
//			ID: "20220523172948",
//			Migrate: func(tx *gorm.DB) error {
//				err := h.BaseMigrate(ctx, tx)
//				if err != nil {
//					log.Errorf(err.Error())
//				}
//				return err
//			},
//		},
//		{
//			ID: "20221013185315",
//			Migrate: func(tx *gorm.DB) error {
//				if err := h.db.AutoMigrate(&model.User{}); err != nil {
//					return err
//				}
//				// current users
//				users := make([]model.User, 0)
//				if err := h.db.Find(&users).Error; err != nil {
//					return err
//				}
//				question := model.Question{}
//				if err := h.db.Where("key = ?", "gender").Find(&question).Error; err != nil {
//					return err
//				}
//				userIDs := make([]string, 0)
//				for _, user := range users {
//					userIDs = append(userIDs, user.ID.String())
//				}
//				// get users gender
//				answers := make([]model.UserAnswer, 0)
//				if err := h.db.Where("user_id in (?) AND question_id = ?", userIDs, question.ID).Find(&answers).Error; err != nil {
//					return err
//				}
//				for i, user := range users {
//					for _, answer := range answers {
//						if user.ID == answer.UserID {
//							users[i].Gender = strings.ToLower(answer.Content)
//						}
//					}
//				}
//
//				if err := tx.Clauses(clause.OnConflict{
//					Columns:   []clause.Column{{Name: "phone_number"}},
//					DoUpdates: clause.AssignmentColumns([]string{"gender"}),
//				}).CreateInBatches(&users, len(users)).Error; err != nil {
//					return err
//				}
//				return nil
//			},
//		},
//	})
//	err := migrate.Migrate()
//	if err != nil {
//		log.Errorf(err.Error())
//	}
//}
//
//func (h *MigrationHandler) Seeding(ctx *gin.Context) {
//	tx := h.db.Begin()
//	defer tx.Rollback()
//
//	if err := seedingScreen(tx); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	if err := seedingQuestion(tx); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	if err := seedingAnswer(tx); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//	seedingRole(tx)
//
//	if err := seedingLanguage(tx); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	if err := seedingEthnicGroup(tx); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	if err := seedingEthnicOrigin(tx); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	tx.Commit()
//	fmt.Println("Seeding successfully")
//}
//
//func seedingScreen(tx *gorm.DB) error {
//	var screens []model.Screen
//	rawData, err := ioutil.ReadFile("resources/screen.json")
//	if err != nil {
//		return err
//	}
//
//	if err := json.Unmarshal(rawData, &screens); err != nil {
//		return err
//	}
//
//	if err := tx.CreateInBatches(&screens, len(screens)).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func seedingQuestion(tx *gorm.DB) error {
//	var questions []model.Question
//	rawData, err := ioutil.ReadFile("resources/question.json")
//	if err != nil {
//		return err
//	}
//
//	if err := json.Unmarshal(rawData, &questions); err != nil {
//		return err
//	}
//
//	if err := tx.CreateInBatches(&questions, len(questions)).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func seedingAnswer(tx *gorm.DB) error {
//	var answers []model.Answer
//	rawData, err := ioutil.ReadFile("resources/answer.json")
//	if err != nil {
//		return err
//	}
//
//	if err := json.Unmarshal(rawData, &answers); err != nil {
//		return err
//	}
//
//	if err := tx.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "question_id"}, {Name: "content"}, {Name: "type"}},
//		DoNothing: true,
//	}).CreateInBatches(&answers, len(answers)).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func seedingRole(tx *gorm.DB) {
//	var tmp []model.Role
//	// Role for Admin Portal
//	tmp = append(tmp, model.Role{IsActive: true, Name: utils.ROLE_NAME_ADMIN, Value: utils.RoleAdmin, Description: utils.ROLE_NAME_ADMIN, PlatformKey: utils.PLATFORM_ADMIN_PORTAL, Key: utils.PLATFORM_ADMIN_PORTAL + "_admin"})
//	tmp = append(tmp, model.Role{IsActive: true, Name: utils.ROLE_NAME_MEMBERSHIP, Value: utils.RoleMembership, Description: utils.ROLE_NAME_MEMBERSHIP, PlatformKey: utils.PLATFORM_APP, Key: utils.PLATFORM_APP + "_membership"})
//	tmp = append(tmp, model.Role{IsActive: true, Name: utils.ROLE_NAME_NORMAL, Value: utils.RoleNormal, Description: utils.ROLE_NAME_NORMAL, PlatformKey: utils.PLATFORM_APP, Key: utils.PLATFORM_APP + "_normal"})
//	tmp = append(tmp, model.Role{IsActive: true, Name: utils.ROLE_NAME_ANONYMOUS, Value: utils.RoleAnonymous, Description: utils.ROLE_NAME_ANONYMOUS, PlatformKey: utils.PLATFORM_APP, Key: utils.PLATFORM_APP + "_anonymous"})
//	for _, v := range tmp {
//		tx.Where("is_active = ? AND name = ? AND platform_key = ? AND key = ?", v.IsActive, v.Name, v.PlatformKey, v.Key).FirstOrCreate(&v)
//	}
//}
//
//func seedingLanguage(tx *gorm.DB) error {
//	type Lang struct {
//		Iso  string `json:"isoCode"`
//		Name string `json:"name"`
//	}
//	rawData, err := ioutil.ReadFile("resources/languages.json")
//	if err != nil {
//		return err
//	}
//
//	var languages []Lang
//	if err := json.Unmarshal(rawData, &languages); err != nil {
//		return err
//	}
//
//	languageQuestion := model.Question{}
//	if err := tx.Model(&model.Question{}).Where("key = ?", "language").Take(&languageQuestion).Error; err != nil {
//		return err
//	}
//
//	var languageAnswers []model.Answer
//	for i, language := range languages {
//		raw, _ := json.Marshal(language)
//		languageAnswers = append(languageAnswers, model.Answer{
//			QuestionID: languageQuestion.ID,
//			Content:    language.Name,
//			Priority:   i,
//			Key:        strings.ToLower(strings.ReplaceAll(language.Name, " ", "_")),
//			JsonValue:  postgres.Jsonb{RawMessage: raw},
//			Type:       3,
//		})
//	}
//
//	if err := tx.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "question_id"}, {Name: "content"}, {Name: "type"}},
//		DoNothing: true,
//	}).CreateInBatches(&languageAnswers, len(languageAnswers)).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func seedingEthnicGroup(tx *gorm.DB) error {
//	rawData, err := ioutil.ReadFile("resources/ethnics.json")
//	if err != nil {
//		return err
//	}
//
//	var ethnicGroups []string
//	if err := json.Unmarshal(rawData, &ethnicGroups); err != nil {
//		return err
//	}
//
//	ethnicGroupQuestion := model.Question{}
//	if err := tx.Model(&model.Question{}).Where("key = ?", "ethnic_group").Take(&ethnicGroupQuestion).Error; err != nil {
//		return err
//	}
//
//	var ethnicGroupAnswers []model.Answer
//	for i, ethnicGroup := range ethnicGroups {
//		ethnicGroupAnswers = append(ethnicGroupAnswers, model.Answer{
//			QuestionID: ethnicGroupQuestion.ID,
//			Content:    ethnicGroup,
//			Priority:   i,
//			Key:        strings.ToLower(strings.ReplaceAll(ethnicGroup, " ", "_")),
//			Type:       3,
//		})
//	}
//
//	if err := tx.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "question_id"}, {Name: "content"}, {Name: "type"}},
//		DoUpdates: clause.AssignmentColumns([]string{"key"}),
//	}).CreateInBatches(&ethnicGroupAnswers, len(ethnicGroupAnswers)).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func seedingEthnicOrigin(tx *gorm.DB) error {
//	rawData, err := ioutil.ReadFile("resources/countries.json")
//	if err != nil {
//		return err
//	}
//
//	var ethnicOrigins []struct {
//		Nationality string `json:"nationality"`
//		EnShortName string `json:"en_short_name"`
//		NumCode     string `json:"num_code"`
//	}
//
//	if err := json.Unmarshal(rawData, &ethnicOrigins); err != nil {
//		return err
//	}
//
//	ethnicOriginQuestion := model.Question{}
//	if err := tx.Model(&model.Question{}).Where("key = ?", "ethnic_origin").Take(&ethnicOriginQuestion).Error; err != nil {
//		return err
//	}
//
//	var ethnicOriginAnswers []model.Answer
//	for i, ethnicOrigin := range ethnicOrigins {
//		raw, _ := json.Marshal(ethnicOrigin)
//		ethnicOriginAnswers = append(ethnicOriginAnswers, model.Answer{
//			QuestionID: ethnicOriginQuestion.ID,
//			Content:    ethnicOrigin.EnShortName,
//			Priority:   i,
//			Key:        strings.ToLower(strings.ReplaceAll(ethnicOrigin.EnShortName, " ", "_")),
//			JsonValue:  postgres.Jsonb{RawMessage: raw},
//			Type:       3,
//		})
//	}
//
//	if err := tx.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "question_id"}, {Name: "content"}, {Name: "type"}},
//		DoUpdates: clause.AssignmentColumns([]string{"key"}),
//	}).CreateInBatches(&ethnicOriginAnswers, len(ethnicOriginAnswers)).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (h *MigrationHandler) MigrateDgraph(ctx *gin.Context) {
//	err := h.dgraph.Alter(ctx, &api.Operation{
//		Schema: `
//			id: string @index(hash) .
//			name: string @index(exact) .
//			dob: datetime .
//			gender: string @index(exact) .
//			height: float @index(float) .
//			country: string @index(hash) .
//			location: geo @index(geo) .
//			filter: [uid] .
//			tagged: [uid] .
//			reference: [uid] @reverse .
//			type: string .
//			coords: float .
//			type User {
//				id: string
//				name: string
//				dob: datetime
//				gender: bool
//				height: float
//				location: Loc
//				filter: [Tag]
//				tagged: [Tag]
//				reference: [User]
//			}
//			type Tag {
//				id: string
//				name: string
//			}
//			type Loc {
//				type: string
//				coords: float
//			}
//		`,
//	})
//	if err != nil {
//		_ = ctx.Error(err)
//	}
//}
