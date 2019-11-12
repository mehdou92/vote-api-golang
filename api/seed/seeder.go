package seed

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mehdou92/vote-api/api/models"
)

var users = []models.User{
	models.User{
		FirstName: "Steven",
		LastName: "victor",
		Email:    "steven@gmail.com",
		Password: "password",
		Dateofbirth: time.Date(1959, 2, 8, 12, 0, 0, 0, time.UTC),
	},
	models.User{
		FirstName: "Steven",
		LastName: "victor",
		Email:    "luther@gmail.com",
		Password: "password",
		Dateofbirth: time.Date(2012, 2, 8, 12, 0, 0, 0, time.UTC),
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
