package main

import (
	"TodoApp"
	"TodoApp/pkg/handler"
	"TodoApp/pkg/repository"
	"TodoApp/pkg/service"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title TodoApp API
// @version 1.0
// @description API Server for TodoApp.

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializating configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error connect to database: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	server := new(TodoApp.Server)
	logrus.Infof("starting the server on %s port", viper.GetString("port"))

	go func() {
		if err = server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("серверу не удалось запуститься: %s", err.Error())
		}
	}()

	logrus.Printf("TodoApp started on port %s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp shutting down")

	if err = server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("ошибка при остановке сервера: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("ошибка при закрытии соединения с БД: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
