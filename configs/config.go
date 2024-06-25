package configs

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    db  *gorm.DB
    err error
)

func ConnectDB() {
    dsn := "host=localhost port=5433 user=postgres dbname=online_store password=saipul1452 sslmode=disable"
    db, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Database connected!")
}

func GetDB() *gorm.DB {
    return db
}
