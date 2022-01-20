package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rsource-open-source/rsource-api/utils"
	"github.com/zeimedee/go-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

var creds = utils.CredentialsInstance

func ConnectDb() {

	port := strconv.Itoa(creds.Port)

	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", creds.Host, creds.Username, creds.Password, creds.Database, port, creds.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Book{})

	DB = Dbinstance{
		Db: db,
	}
}