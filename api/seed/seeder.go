package seed

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zineb-ada/cyrkl/api/models"
)

var users = []models.User{
	models.User{
		Name:      "lara",
		Lastname:  "clu",
		Email:     "lar-clu@gmail.com",
		Telephone: "0870956788",
		Password:  "123456",
	},
	models.User{
		Name:      "nani",
		Lastname:  "rail",
		Email:     "nan-rail@gmail.com",
		Telephone: "0875138788",
		Password:  "123456",
	},
}

var slots = []models.Slot{
	models.Slot{
		UserID:       2,
		Dateandhours: time.Date(2023, 2, 10, 12, 0, 0, 0, time.Local),
		Lieu:         "Café Flora",
	},
	models.Slot{
		UserID:       1,
		Dateandhours: time.Date(2024, 12, 15, 12, 30, 0, 0, time.Local),
		Lieu:         "Flunch",
	},
}

var invitations = []models.Invitation{
	models.Invitation{
		Statut:    "In progress",
		InviterID: 1,
		InvitedID: 2,
		SlotID:    2,
	},
	models.Invitation{
		Statut:    "In progress",
		InviterID: 2,
		InvitedID: 1,
		SlotID:    1,
	},
}

var usersalgo = []models.UsersAlgo{
	models.UsersAlgo{
		UserID:    2,
		UsersList: 3,
	},
	models.UsersAlgo{
		UserID:    1,
		UsersList: 3,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Slot{}, &models.Invitation{}, models.UsersAlgo{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Slot{}, &models.Invitation{}, models.UsersAlgo{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Slot{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Invitation{}).AddForeignKey("inviter_id", "users(id)", "cascade", "cascade").AddForeignKey("invited_id", "users(id)", "cascade", "cascade").AddForeignKey("slot_id", "slots(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Invitation{}).AddForeignKey("inviter_id", "users(id)", "cascade", "cascade").AddForeignKey("invited_id", "users(id)", "cascade", "cascade").AddForeignKey("slot_id", "slots(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.UsersAlgo{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		slots[i].UserID = users[i].ID
		err = db.Debug().Model(&models.Slot{}).Create(&slots[i]).Error
		if err != nil {
			log.Fatalf("cannot seed slot table: %v", err)
		}
		usersalgo[i].UserID = users[i].ID
		err = db.Debug().Model(&models.UsersAlgo{}).Create(&usersalgo[i]).Error
		if err != nil {
			log.Fatalf("cannot seed useralogo table: %v", err)
		}
	}
	for i, _ := range users {
		switch i {
		case 0:
			invitations[i].InviterID = users[0].ID
			invitations[i].InvitedID = users[1].ID
			invitations[i].SlotID = slots[0].ID
		case 1:
			invitations[i].InviterID = users[1].ID
			invitations[i].InvitedID = users[0].ID
			invitations[i].SlotID = slots[1].ID
		}
		err = db.Debug().Model(&models.Invitation{}).Create(&invitations[i]).Error
		if err != nil {
			log.Fatalf("cannot seed invitation table: %v", err)
		}
	}
}

// err := db.Debug().DropTableIfExists(&models.InvitationsReceived{}).Error
// if err != nil {
// 	log.Fatalf("cannot drop table: %v", err)
// }
// err = db.Debug().AutoMigrate(&models.InvitationsReceived{}).Error
// if err != nil {
// 	log.Fatalf("cannot migrate table: %v", err)
// }
// err := db.Debug().Model(&models.Calendar{}).AddForeignKey("user_sender_id", "users(id)", "cascade", "cascade").Error
// if err != nil {
// 	log.Fatalf("attaching foreign key error: %v", err)
// }
// for i, _ := range invitationsReceived {
// 	err := db.Create(&invitationsReceived[i]).Error
// 	if err != nil {
// 		log.Fatalf("cannot seed invitation received: %v", err)
// 	}
// }
// for i, _ := range users {
// 	err := db.Create(&users[i]).Error
// 	if err != nil {
// 		log.Fatalf("cannot create user: %v", err)
// 	}
// 	calendars[i].UserID = users[i].ID

// 	err = db.Create(&calendars[i]).Error
// 	if err != nil {
// 		log.Fatalf("cannot seed calendar: %v", err)
// 	}
// }
// db.Create(&users[i])
// 	calendars[i].UserID = users[i].ID

// 	err := db.Debug().Model(&models.Calendar{}).Create(&calendars[i]).Error
// 	if err != nil {
// 		log.Fatalf("cannot seed calendar table: %v", err)
// 	}
// }
// 	db.Create(&users[i])
// }
