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

	bookByAuthorCache, r1 := redis.New[dto.Books](cfg, "book-by-author", 24*time.Hour)
	defer r1.Close()

	authorCache, r2 := redis.New[dto.Author](cfg, "author", 24*time.Hour)
	defer r2.Close()

	validator := govalidator.New()
	repository := repository.New(db)

	// books
	bookCreator := usecase.NewCreateBook(repository, bookCache, validator)
	bookDeleter := usecase.NewDeleteBook(repository, bookCache)
	bookUpdater := usecase.NewUpdateBook(repository, validator, bookCache)
	bookAllGetter := usecase.NewBookAllGetter(repository)
	bookGetter := usecase.NewGetBook(bookCache, repository)
	byAuthorBookGetter := usecase.NewGetBookByAuthor(repository, bookByAuthorCache)
	bookController := controller.NewBookController(
		bookCreator,
		bookUpdater,
		bookDeleter,
		bookGetter,
		bookAllGetter,
		byAuthorBookGetter,
	)

	// author
	authorCreator := usecase.NewAuthorCreator(repository, authorCache, validator)
	authorDeleter := usecase.NewAuthorDeleter(repository, authorCache, bookCache)
	authorUpdater := usecase.NewAuthorUpdater(repository, authorCache, validator)
	authorGetter := usecase.NewAuthorGetter(repository, authorCache)
	authorAllGetter := usecase.NewAuthorAllGetter(repository)
	authorController := controller.NewAuthorController(
		authorCreator,
		authorUpdater,
		authorDeleter,
		authorGetter,
		authorAllGetter,
	)

	app := fiber.New()
	app.Use(recover.New())

	bookGroup := app.Group("/v1/books")
	bookGroup.Post("", bookController.Create)
	bookGroup.Put("/:id", bookController.Update)
	bookGroup.Delete("/:id", bookController.Delete)
	bookGroup.Get("", bookController.GetAll)
	bookGroup.Get("/:id", bookController.GetById)

	authorGroup := app.Group("/v1/authors")
	authorGroup.Post("", authorController.Create)
	authorGroup.Put("/:id", authorController.Update)
	authorGroup.Delete("/:id", authorController.Delete)
	authorGroup.Get("", authorController.GetAll)
	authorGroup.Get("/:id", authorController.GetById)
	authorGroup.Get("/:authorId/books", bookController.GetByAuthor)

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
