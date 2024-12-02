package db

import (
	"database/sql"
	"fmt"
	"log"
	"service-sentinel/monitors"
	"service-sentinel/utils"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var sqlDB *sql.DB

func connectToDB (dbName string) (error) {
	password := utils.GetDBPassword()
	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=%s port=5432 sslmode=disable", password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
		return err
    }
	gormDB = db
	sqlDB, err = gormDB.DB()
	if err != nil {
		return err
	}
	return nil
}

func createTables () (error) {
	httpErr := gormDB.AutoMigrate(&monitors.HttpMonitor{})
	if httpErr != nil {
		return httpErr
	}

	pingErr := gormDB.AutoMigrate(&monitors.PingMonitor{})
	if pingErr != nil {
		return pingErr
	}
	return nil
}

func Init (dbName string) (error) {
	connectionErr := connectToDB(dbName)
	if connectionErr != nil {
		log.Fatal("ERROR in connecting to db:", connectionErr)
		panic(connectionErr)
	}
	
	createTablesErr := createTables()
	if createTablesErr != nil {
		log.Fatal("ERROR in creating tables:", createTablesErr)
		panic(createTablesErr)
	}

	return nil
}

func ShutDown () (error) {
	return sqlDB.Close()
}
