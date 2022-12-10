package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	APPNAME = "doan backend golang"
)

func main() {
	logrus.Info(APPNAME)

	fmt.Print("test source golang")
}
