package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	goback "go.back"
	"go.back/configs"
	"go.back/internal/handler"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("Error reading config file: %s", err)
	}

	handlers := handler.NewHandler()

	server := new(goback.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}

}
