package postgres

import (
	"fmt"
	"log"

	"github.com/bowodev/go-library-management/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timeZone=Asia/Jakarta",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("[FATAL] failed to init db, errors: %v", err)
	}

	return db
}
