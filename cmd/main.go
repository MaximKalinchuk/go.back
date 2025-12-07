package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	goback "go.back"
	"go.back/configs"
	"go.back/internal/handler"
	"go.back/internal/middleware"
	"go.back/internal/repository"
	"go.back/internal/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("Ошибка чтения конфигурационных файлов: %s", err)
	}

	err := gotenv.Load()

	if err != nil {
		logrus.Fatalf("Ошибка загрузки env: %s", err.Error())
	}

	db, err := configs.NewPostgresDB(configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Ошибка инициализации базы данных: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	middleware := middleware.NewMiddleware(services.Authorization)
	handlers := handler.NewHandler(services, middleware)

	server := new(goback.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Ошибка старта сервера: %s", err.Error())
	}

}
