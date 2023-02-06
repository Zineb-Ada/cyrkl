package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/zineb-ada/cyrkl/api/models"
)

// ID Name Lastname Email Urlphoto Telephone Password EmailVerifiedAt CreatedAt UpdatedAt

var users = []models.User{
	models.User{
		Name:      "laurie",
		Lastname:  "clu",
		Email:     "laurie-clu@gmail.com",
		Telephone: "0876756788",
		Password:  "password",
	},
	models.User{
		Name:      "nono",
		Lastname:  "rag",
		Email:     "nono-rag@gmail.com",
		Telephone: "0875436788",
		Password:  "password",
	},
}

func Load(db *gorm.DB) {
	for i, _ := range users {
		db.Create(&users[i])
	}
}
