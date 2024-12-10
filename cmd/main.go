package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bowodev/go-library-management/config"
	"github.com/bowodev/go-library-management/internal/controller"
	"github.com/bowodev/go-library-management/internal/dto"
	govalidator "github.com/bowodev/go-library-management/internal/go_validator"
	"github.com/bowodev/go-library-management/internal/postgres"
	"github.com/bowodev/go-library-management/internal/redis"
	"github.com/bowodev/go-library-management/internal/repository"
	"github.com/bowodev/go-library-management/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/root")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Panicf("failed to read config, error: %v", err)
	}

	cfg := config.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("failed to parse config, error: %v", err)
	}

	db := postgres.New(cfg)
	bookCache, r := redis.New[dto.Book](cfg, "book", 24*time.Hour)
	defer r.Close()

	validator := govalidator.New()

	repository := repository.New(db)
	bookCreator := usecase.NewCreateBook(repository, bookCache, validator)
	bookController := controller.NewBookController(bookCreator)

	app := fiber.New()
	app.Use(recover.New())

	bookGroup := app.Group("/v1/books")
	bookGroup.Post("", bookController.Create)

	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("go-library-management was successful shutdown.")
}
