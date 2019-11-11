package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresDB struct {
	db *gorm.DB
}

func new() {

}

func test() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=mehdi dbname=vote password=123456789")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
