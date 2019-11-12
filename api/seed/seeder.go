package seed

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mehdou92/vote-api/api/models"
)

var users = []models.User{
	models.User{
		FirstName:   "Steven",
		LastName:    "victor",
		Email:       "steven@gmail.com",
		Password:    "password",
		DateOfBirth: time.Date(1959, 2, 8, 12, 0, 0, 0, time.UTC),
	},
	models.User{
		FirstName:   "Martin",
		LastName:    "Duck",
		Email:       "luther@gmail.com",
		Password:    "password",
		DateOfBirth: time.Date(1993, 2, 8, 12, 0, 0, 0, time.UTC),
	},
}

var votes = []models.Vote{
	models.Vote{
		Title:   "Title 1",
		Desc: "Hello world 1",
	},
	models.Vote{
		Title:   "Title 2",
		Desc: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Vote{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Vote{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Vote{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		votes[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Vote{}).Create(&votes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed votes table: %v", err)
		}
	}
}
