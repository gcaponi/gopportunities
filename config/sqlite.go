package config

import (
	"os"

	"github.com/gcaponi/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqLite")
	dbPath := "./db/main.db"

	// Check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file no found, crating...")
		
		//Create DB file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil{
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil{
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	
	// Return DB
	return db, nil
}