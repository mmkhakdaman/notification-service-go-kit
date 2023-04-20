package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"notification_service/config"
	"notification_service/models"
	"testing"
)

var db *gorm.DB

func LoadDB() {
	dbConfig := config.GetDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Notification{})
	if err != nil {
		return
	}
}

func CreateTestGormDB(t *testing.T) (*gorm.DB, func()) {
	// sqllite
	var err error
	db, err = gorm.Open(
		sqlite.Open("test.db"),
		&gorm.Config{},
	)
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Notification{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	cleanup := func() {
		db.Migrator().DropTable(&models.Notification{})
	}

	return db, cleanup
}

func GetDB() *gorm.DB {
	return db
}
