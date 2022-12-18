package seed

import (
	"log"

	"github.com/Zineb-Ada/cyrkl/api/models"
	"github.com/jinzhu/gorm"
)

// ID Name Lastname Email Urlphoto Telephone Password EmailVerifiedAt CreatedAt UpdatedAt

var users = []models.User{
	models.User{
		Name:      "laurie",
		Lastname:  "clu",
		Email:     "laurie-clu@gmail.com",
		Urlphoto:  "thisisanurl",
		Telephone: "0876756788",
		Password:  "password",
	},
	models.User{
		Name:      "nono",
		Lastname:  "rag",
		Email:     "nono-rag@gmail.com",
		Urlphoto:  "thisisanurl",
		Telephone: "0876756788",
		Password:  "password",
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

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		// posts[i].AuthorID = users[i].ID

		// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}
}
