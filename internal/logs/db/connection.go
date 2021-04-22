package db

import "github.com/jinzhu/gorm"

var connection *gorm.DB

func GetConnection() *gorm.DB {
	if connection == nil {
		return createConnection()
	}

	return connection
}
