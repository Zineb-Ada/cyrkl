package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/zineb-ada/cyrkl/api/models"
)

// ID Name Lastname Email Urlphoto Telephone Password EmailVerifiedAt CreatedAt UpdatedAt

var users = []models.User{
	models.User{
		Name:           "laurie",
		Lastname:       "clu",
		Email:          "laurie-clu@gmail.com",
		Urlphoto:       "thisisanurl",
		Telephone:      "0876756788",
		Password:       "password",
		Position:       "work_position",
		Positionsought: []string{"pos1", "pos2", "pos3", "pos4"},
		Industry:       "industry",
		Industrysought: []string{"ind1", "ind2", "ind3", "ind4"},
	},
	models.User{
		Name:           "nono",
		Lastname:       "rag",
		Email:          "nono-rag@gmail.com",
		Urlphoto:       "thisisanurl",
		Telephone:      "0875436788",
		Password:       "password",
		Position:       "work_position",
		Positionsought: []string{"pos1", "pos2", "pos3", "pos4"},
		Industry:       "industry",
		Industrysought: []string{"ind1", "ind2", "ind3", "ind4"},
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
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
