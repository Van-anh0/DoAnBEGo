package route

//
//import (
//	"github.com/caarlos0/env/v6"
//	swaggerFiles "github.com/swaggo/files"
//	ginSwagger "github.com/swaggo/gin-swagger"
//	"gitlab.com/jfcore/common/ginext"
//	"gitlab.com/jfcore/common/service"
//	"iot/conf"
//	_ "iot/docs"
//	"iot/pkg/handlers"
//	"iot/pkg/repo"
//	service2 "iot/pkg/service"
//)
//
//type extraSetting struct {
//	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
//}
//
//type Service struct {
//	*service.BaseApp
//	setting *extraSetting
//}
//
//// NewService
//// @title Swagger Example API
//// @version 1.0
//// @description This is a sample server celler server.
//// @termsOfService http://swagger.io/terms/
//// @contact.name API Support
//// @contact.url http://www.swagger.io/support
//// @contact.email support@swagger.io
//// @license.name Apache 2.0
//// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//// @host localhost:8001
//// @BasePath /api/v1
//func NewService() *Service {
//	s := &Service{
//		service.NewApp("iot", "v1.0"),
//		&extraSetting{},
//	}
//	_ = env.Parse(s.setting)
//
//	db := s.GetDB()
//	repoPG := repo.NewPGRepo(db)
//	if s.setting.DbDebugEnable {
//		db = db.Debug()
//	}
//
//	deviceService := service2.NewDeviceService(repoPG)
//	device := handlers.NewDeviceHandlers(deviceService)
//
//	deviceMetaService := service2.NewDeviceMetaService(repoPG)
//	deviceMeta := handlers.NewDeviceMetaHandlers(deviceMetaService)
//
//	deviceSensorService := service2.NewDeviceSensorService(repoPG)
//	deviceSensor := handlers.NewDeviceSensorHandlers(deviceSensorService)
//
//	dsdService := service2.NewDsdService(repoPG)
//	dsd := handlers.NewDsdHandlers(dsdService)
//
//	sdmService := service2.NewSdmService(repoPG)
//	sdm := handlers.NewSdmHandlers(sdmService)
//
//	userService := service2.NewUserService(repoPG)
//	user := handlers.NewUserHandlers(userService)
//
//	roleService := service2.NewRoleService(repoPG)
//	role := handlers.NewRoleHandlers(roleService)
//
//	urrService := service2.NewUrrService(repoPG)
//	urr := handlers.NewUrrHandlers(urrService)
//
//	udrService := service2.NewUdrService(repoPG)
//	udr := handlers.NewUdrHandlers(udrService)
//
//	webhookService := service2.NewWebhookService(repoPG)
//	webhook := handlers.NewWebhookHandlers(webhookService)
//
//	if conf.GetEnv().EnvName == "dev" {
//		s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//	}
//
//	v1Api := s.Router.Group("/api/v1")
//
//	// Webhook
//	v1Api.POST("/webhook/wanesy/:id/up", ginext.WrapHandler(webhook.Webhook))
//
//	// Device
//	v1Api.PUT("/device/create", ginext.WrapHandler(device.Create))
//	v1Api.PUT("/device/update/:id", ginext.WrapHandler(device.Update))
//	v1Api.DELETE("/device/delete/:id", ginext.WrapHandler(device.Delete))
//	v1Api.GET("/device/get-one/:id", ginext.WrapHandler(device.GetOne))
//	v1Api.GET("/device/get-list", ginext.WrapHandler(device.GetList))
//
//	// Device Meta
//	v1Api.PUT("/device-meta/create", ginext.WrapHandler(deviceMeta.Create))
//	v1Api.PUT("/device-meta/update/:id", ginext.WrapHandler(deviceMeta.Update))
//	v1Api.DELETE("/device-meta/delete/:id", ginext.WrapHandler(deviceMeta.Delete))
//	v1Api.GET("/device-meta/get-one/:id", ginext.WrapHandler(deviceMeta.GetOne))
//	v1Api.GET("/device-meta/get-list", ginext.WrapHandler(deviceMeta.GetList))
//
//	// Device Sensor
//	v1Api.PUT("/device-sensor/create", ginext.WrapHandler(deviceSensor.Create))
//	v1Api.PUT("/device-sensor/update/:id", ginext.WrapHandler(deviceSensor.Update))
//	v1Api.DELETE("/device-sensor/delete/:id", ginext.WrapHandler(deviceSensor.Delete))
//	v1Api.GET("/device-sensor/get-one/:id", ginext.WrapHandler(deviceSensor.GetOne))
//	v1Api.GET("/device-sensor/get-list", ginext.WrapHandler(deviceSensor.GetList))
//
//	// Device Sensor Data
//	v1Api.PUT("/dsd/create", ginext.WrapHandler(dsd.Create))
//	v1Api.PUT("/dsd/update/:id", ginext.WrapHandler(dsd.Update))
//	v1Api.DELETE("/dsd/delete/:id", ginext.WrapHandler(dsd.Delete))
//	v1Api.GET("/dsd/get-one/:id", ginext.WrapHandler(dsd.GetOne))
//	v1Api.GET("/dsd/get-list", ginext.WrapHandler(dsd.GetList))
//
//	// Stream Data Method
//	v1Api.PUT("/sdm/create", ginext.WrapHandler(sdm.Create))
//	v1Api.PUT("/sdm/update/:id", ginext.WrapHandler(sdm.Update))
//	v1Api.DELETE("/sdm/delete/:id", ginext.WrapHandler(sdm.Delete))
//	v1Api.GET("/sdm/get-one/:id", ginext.WrapHandler(sdm.GetOne))
//	v1Api.GET("/sdm/get-list", ginext.WrapHandler(sdm.GetList))
//
//	// User
//	v1Api.PUT("/user/create", ginext.WrapHandler(user.Create))
//	v1Api.PUT("/user/update/:id", ginext.WrapHandler(user.Update))
//	v1Api.DELETE("/user/delete/:id", ginext.WrapHandler(user.Delete))
//	v1Api.GET("/user/get-one/:id", ginext.WrapHandler(user.GetOne))
//
//	// Role
//	v1Api.PUT("/role/create", ginext.WrapHandler(role.Create))
//	v1Api.PUT("/role/update/:id", ginext.WrapHandler(role.Update))
//	v1Api.DELETE("/role/delete/:id", ginext.WrapHandler(role.Delete))
//	v1Api.GET("/role/get-one/:id", ginext.WrapHandler(role.GetOne))
//	v1Api.GET("/role/get-list", ginext.WrapHandler(role.GetList))
//
//	// User Role Relationship
//	v1Api.PUT("/urr/create", ginext.WrapHandler(urr.Create))
//	v1Api.PUT("/urr/update/:id", ginext.WrapHandler(urr.Update))
//	v1Api.DELETE("/urr/delete/:id", ginext.WrapHandler(urr.Delete))
//	v1Api.GET("/urr/get-one/:id", ginext.WrapHandler(urr.GetOne))
//	v1Api.GET("/urr/get-list", ginext.WrapHandler(urr.GetList))
//
//	// User Device Relationship
//	v1Api.PUT("/udr/create", ginext.WrapHandler(udr.Create))
//	v1Api.PUT("/udr/update/:id", ginext.WrapHandler(udr.Update))
//	v1Api.DELETE("/udr/delete/:id", ginext.WrapHandler(udr.Delete))
//	v1Api.GET("/udr/get-one/:id", ginext.WrapHandler(udr.GetOne))
//	v1Api.GET("/udr/get-list", ginext.WrapHandler(udr.GetList))
//
//	//handleMQTT := handlers.NewMQTTHandlers()
//	//handleMQTT.ListenMQTT()
//
//	// Migrate
//	migrateHandler := handlers.NewMigrationHandler(db)
//	v1Api.POST("/internal/migrate", migrateHandler.Migrate)
//
//	return s
//}
