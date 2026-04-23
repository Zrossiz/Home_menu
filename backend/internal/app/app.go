package app

import (
	"fmt"
	"net/http"

	"github.com/Zrossiz/Home_menu/backend/internal/config"
	"github.com/Zrossiz/Home_menu/backend/internal/service"
	"github.com/Zrossiz/Home_menu/backend/internal/storage/postgres"
	"github.com/Zrossiz/Home_menu/backend/internal/transport/http/handler"
	"github.com/Zrossiz/Home_menu/backend/internal/transport/http/router"
	"github.com/Zrossiz/Home_menu/backend/pkg/logger"
	"github.com/rs/cors"
)

func Start() {
	cfg := config.GetConfig()

	log, err := logger.New(cfg.Logger.Level)
	if err != nil {
		panic(fmt.Errorf("error init logger: %v", err))
	}

	dbConn, err := postgres.Connect(cfg.DB.DBURI)
	if err != nil {
		panic(err)
	}

	postgresStore := postgres.New(dbConn)
	serv := service.New(service.Storage{
		CategoryPostgres:   &postgresStore.Category,
		DishPostgres:       &postgresStore.Dish,
		AttachmentPostgres: &postgresStore.Attachment,
	}, log)
	httpHandler := handler.New(handler.Service{
		Category:   &serv.Category,
		Dish:       &serv.Dish,
		Attachment: &serv.Attachment,
	}, log)
	httpRouter := router.New(router.Handler{
		Category:   &httpHandler.Category,
		Dish:       &httpHandler.Dish,
		Attachment: &httpHandler.Attachment,
	})

	handlerWithCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.Application.ClientURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(httpRouter)

	log.Sugar().Infof("start http server on address: %v", cfg.Application.Port)
	err = http.ListenAndServe(cfg.Application.Port, handlerWithCORS)
	if err != nil {
		panic(fmt.Errorf("error init http server: %v", err))
	}
}
