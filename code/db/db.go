package db

import (
	"fmt"
	"log"
	"serveos-datasync/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(cfg config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)

	var err error
	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		DB, err = sqlx.Connect("mysql", dsn)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %s", attempts, maxAttempts, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully")
}
