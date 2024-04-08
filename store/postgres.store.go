package store

import (
	"fmt"
	"log"

	"github.com/mayankr5/url_shortner/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

const (
	host     = "localhost"
	port     = 5432
	user     = "mayank"
	password = "123"
	dbname   = "url_shortner"
)

func Connect() error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	fmt.Println("Postgres Connnected!")
	db.AutoMigrate(&model.User{}, &model.AuthToken{}, &model.UserURLs{})

	DB = Dbinstance{
		Db: db,
	}
	return nil
}
