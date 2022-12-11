package main

import (
	"context"
	"doan/conf"
	"doan/pkg/route"
	"doan/pkg/utils"
	"github.com/praslar/cloud0/logger"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	APPNAME = "doan backend golang"
)

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()
	// Dev
	logger.DefaultLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		PadLevelText:     true,
		ForceQuote:       true,
		QuoteEmptyFields: true,
	})

	//if err := utils.InitAwsSession(); err != nil {
	//	logger.Tag("main").Error(err)
	//}
	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
