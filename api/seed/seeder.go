package seed

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zineb-ada/cyrkl/api/models"
)

var users = []models.User{
	models.User{
		Name:           "Lara",
		Lastname:       "Clure",
		Email:          "lar-clu@gmail.com",
		Urlphoto:       "https://p5.storage.canalblog.com/50/89/1228697/127715041.jpg",
		Telephone:      "0870956701",
		Password:       "123456",
		Position:       "Entrepreneur",
		Positionsought: []string{"Mentors", "Managers", "Investisseurs", "Développeur"},
		Industry:       "Conseil",
		Industrysought: []string{"Santé", "Energie", "Marketing", "Industrie"},
	},
	models.User{
		Name:           "Christophe",
		Lastname:       "Leblanc",
		Email:          "chris-leblanc@gmail.com",
		Urlphoto:       "https://studiolecarre.com/wp-content/uploads/2020/07/181105-1748-Data-SrcSet.jpg",
		Telephone:      "0870956702",
		Password:       "123456",
		Position:       "CTO",
		Positionsought: []string{"Managers", "Experts sectoriels", "Conseillers", "Universitaire"},
		Industry:       "Energie",
		Industrysought: []string{"Energie", "Education", "Conseil", "Santé"},
	},
	models.User{
		Name:           "Nani",
		Lastname:       "Rail",
		Email:          "nani-rail@gmail.com",
		Urlphoto:       "https://www.portraitprofessionnel.fr/wp-content/uploads/2022/02/Studio_Photo_Paris.jpg",
		Telephone:      "0870956703",
		Password:       "123456",
		Position:       "Conseil",
		Positionsought: []string{"Développeur", "Investisseurs", "Managers", "Mentors"},
		Industry:       "Entrepreneur",
		Industrysought: []string{"Industrie", "Marketing", "Energie", "Santé"},
	},
	models.User{
		Name:           "Julien",
		Lastname:       "De balzac",
		Email:          "julien-de@gmail.com",
		Urlphoto:       "https://1120corporate-1278.kxcdn.com/wp-content/uploads/portrait-professionnel-frederic-mazel-inextenso-mediterranee.jpg",
		Telephone:      "0870956704",
		Password:       "123456",
		Position:       "Manager",
		Positionsought: []string{"Entrepreneurs", "Investisseurs", "Experts sectoriels", "Consultant"},
		Industry:       "Marketing",
		Industrysought: []string{"Education", "Finance", "Marketing", "Conseil"},
	},
	models.User{
		Name:           "Lara",
		Lastname:       "Elle",
		Email:          "lara-elle@gmail.com",
		Urlphoto:       "https://thumbs.dreamstime.com/b/femme-d-affaires-noire-professionnelle-119443628.jpg",
		Telephone:      "0870956705",
		Password:       "123456",
		Position:       "Conseil",
		Positionsought: []string{"Développeur", "Investisseurs", "Managers", "Mentors"},
		Industry:       "Entrepreneur",
		Industrysought: []string{"Industrie", "Marketing", "Energie", "Santé"},
	},
	models.User{
		Name:           "Maxime",
		Lastname:       "Jores",
		Email:          "max-jores@gmail.com",
		Urlphoto:       "https://www.studioah.fr/wp-content/uploads/2020/11/photo_de_CV_studio_nantes_%C2%A9studioah_portrait_Professionnelle-4-2-768x768.jpg",
		Telephone:      "0870956706",
		Password:       "123456",
		Position:       "Entrepreneur",
		Positionsought: []string{"Mentors", "Managers", "Investisseurs", "Développeur"},
		Industry:       "Conseil",
		Industrysought: []string{"Santé", "Energie", "Marketing", "Industrie"},
	},
	models.User{
		Name:           "Justine",
		Lastname:       "Droite",
		Email:          "justine-droite@gmail.com",
		Urlphoto:       "https://media.istockphoto.com/id/1163294201/fr/photo/femme-daffaires-confiante-de-sourire-posant-avec-des-bras-pliés.jpg?s=612x612&w=0&k=20&c=mS7gnRXKrl6dRhJx3g4qwMr9UVcffa5vp5lCsnF0YKc=",
		Telephone:      "0870956707",
		Password:       "123456",
		Position:       "CTO",
		Positionsought: []string{"Managers", "Experts sectoriels", "Conseillers", "Universitaire"},
		Industry:       "Energie",
		Industrysought: []string{"Energie", "Education", "Conseil", "Santé"},
	},
	models.User{
		Name:           "Reda",
		Lastname:       "Zniber",
		Email:          "reda-zniber@gmail.com",
		Urlphoto:       "https://www.fredericvigier.com/wp-content/uploads/2021/01/directeur-communication-photo-cv.jpg",
		Telephone:      "0870956708",
		Password:       "123456",
		Position:       "Conseil",
		Positionsought: []string{"Développeur", "Investisseurs", "Managers", "Mentors"},
		Industry:       "Entrepreneur",
		Industrysought: []string{"Industrie", "Marketing", "Energie", "Santé"},
	},
	models.User{
		Name:           "Noemie",
		Lastname:       "Ragout",
		Email:          "noemie-rag@gmail.com",
		Urlphoto:       "https://studiolecarre.com/wp-content/uploads/2022/10/2021-03-11_10-19-23-portrait-entreprise-femme-corporate-2.webp",
		Telephone:      "0870956709",
		Password:       "123456",
		Position:       "Entrepreneur",
		Positionsought: []string{"Mentors", "Managers", "Investisseurs", "Développeur"},
		Industry:       "Conseil",
		Industrysought: []string{"Santé", "Energie", "Marketing", "Industrie"},
	},
	models.User{
		Name:           "Kevin",
		Lastname:       "Guedj",
		Email:          "kevin-guedj@gmail.com",
		Urlphoto:       "https://www.reportages-metiers.fr/wp-content/uploads/2022/01/Portrait_Patrick_KMC_Conseils_BD_1-681x1024.jpg",
		Telephone:      "0870956710",
		Password:       "123456",
		Position:       "Entrepreneur",
		Positionsought: []string{"Mentors", "Managers", "Investisseurs", "Développeur"},
		Industry:       "Conseil",
		Industrysought: []string{"Santé", "Energie", "Marketing", "Industrie"},
	},
}

var slots = []models.Slot{
	models.Slot{
		UserID:       1,
		Dateandhours: time.Date(2023, 2, 10, 12, 0, 0, 0, time.Local),
		Lieu:         "Paris 1er",
	},
	models.Slot{
		UserID:       2,
		Dateandhours: time.Date(2023, 12, 15, 12, 30, 0, 0, time.Local),
		Lieu:         "Paris 2e",
	},
	models.Slot{
		UserID:       3,
		Dateandhours: time.Date(2023, 04, 13, 13, 0, 0, 0, time.Local),
		Lieu:         "Paris 3e",
	},
	models.Slot{
		UserID:       4,
		Dateandhours: time.Date(2023, 05, 16, 13, 30, 0, 0, time.Local),
		Lieu:         "Paris 4e",
	},
	models.Slot{
		UserID:       5,
		Dateandhours: time.Date(2023, 04, 26, 12, 0, 0, 0, time.Local),
		Lieu:         "Paris 5e",
	},
	models.Slot{
		UserID:       6,
		Dateandhours: time.Date(2023, 06, 02, 13, 15, 0, 0, time.Local),
		Lieu:         "Paris 18e",
	},
	models.Slot{
		UserID:       7,
		Dateandhours: time.Date(2023, 07, 20, 12, 30, 0, 0, time.Local),
		Lieu:         "Paris 16e",
	},
	models.Slot{
		UserID:       8,
		Dateandhours: time.Date(2023, 06, 01, 13, 30, 0, 0, time.Local),
		Lieu:         "Paris 15e",
	},
	models.Slot{
		UserID:       9,
		Dateandhours: time.Date(2023, 07, 20, 12, 30, 0, 0, time.Local),
		Lieu:         "Paris 11e",
	},
	models.Slot{
		UserID:       10,
		Dateandhours: time.Date(2023, 05, 17, 13, 00, 0, 0, time.Local),
		Lieu:         "Paris 8e",
	},
}

var invitations = []models.Invitation{
	models.Invitation{
		Statut:    "in progress",
		InviterID: 1,
		InvitedID: 2,
		SlotID:    2,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 2,
		InvitedID: 3,
		SlotID:    3,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 3,
		InvitedID: 4,
		SlotID:    4,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 4,
		InvitedID: 5,
		SlotID:    5,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 5,
		InvitedID: 6,
		SlotID:    6,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 6,
		InvitedID: 7,
		SlotID:    7,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 7,
		InvitedID: 8,
		SlotID:    8,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 8,
		InvitedID: 9,
		SlotID:    9,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 9,
		InvitedID: 10,
		SlotID:    10,
	},
	models.Invitation{
		Statut:    "in progress",
		InviterID: 10,
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

	err = db.Debug().Model(&models.Invitation{}).AddForeignKey("inviter_id", "users(id)", "cascade", "cascade").
		AddForeignKey("invited_id", "users(id)", "cascade", "cascade").
		AddForeignKey("slot_id", "slots(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Invitation{}).AddForeignKey("inviter_id", "users(id)", "cascade", "cascade").
		AddForeignKey("invited_id", "users(id)", "cascade", "cascade").
		AddForeignKey("slot_id", "slots(id)", "cascade", "cascade").Error
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
		// usersalgo[i].UserID = users[i].ID
		// err = db.Debug().Model(&models.UsersAlgo{}).Create(&usersalgo[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed useralogo table: %v", err)
		// }
	}
	for i, _ := range users {
		invitations[i].InviterID = users[i].ID
		if i == len(users)-1 {
			invitations[i].InvitedID = users[0].ID
			invitations[i].SlotID = slots[0].ID
		} else {
			invitations[i].InvitedID = users[i+1].ID
			invitations[i].SlotID = slots[i+1].ID
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
