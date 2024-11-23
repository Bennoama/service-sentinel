package db

import (
	"database/sql"
	"log"
	"service-sentinel/samplers"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var sqlDB *sql.DB

func connectToDB () (error) {
	dsn := "host=localhost user=postgres password=bgr5znTj dbname=serviceSentinel port=5432 sslmode=disable"
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
	httpErr := gormDB.AutoMigrate(&samplers.HttpSampler{})
	if httpErr != nil {
		return httpErr
	}

	pingErr := gormDB.AutoMigrate(&samplers.PingSampler{})
	if pingErr != nil {
		return pingErr
	}
	return nil
}


func Init () (error) {
	connectionErr := connectToDB()
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