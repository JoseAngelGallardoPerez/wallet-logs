package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/Confialink/wallet-logs/internal/logs/config"
)

// CreateConnection creates connection with database
func createConnection() *gorm.DB {
	cfg := config.GetConfig()
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Schema)
	db, err := gorm.Open(
		"mysql",
		connectionString,
	)

	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
		return nil
	}

	if cfg.Db.IsDebugMode {
		return db.Debug()
	}

	return db
}
