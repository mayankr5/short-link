package store

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mayankr5/url_shortner/config"
	"github.com/mayankr5/url_shortner/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() error {
	var err error
	p := config.Config("PGPORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("PGHOST"), port, config.Config("PGUSER"), config.Config("PGPASSWORD"), config.Config("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		return err
	}

	fmt.Println("Postgres Connnected!")
	db.AutoMigrate(&model.User{}, &model.AuthToken{}, &model.UserURL{})

	DB = Dbinstance{
		Db: db,
	}
	return nil
}
