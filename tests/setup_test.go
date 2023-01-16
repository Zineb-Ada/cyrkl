package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/zineb-ada/cyrkl/api/controllers"
	"github.com/zineb-ada/cyrkl/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TEST_DB_DRIVER")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Name:           "toto",
		Lastname:       "titi",
		Email:          "toto_titi@gmail.com",
		Urlphoto:       "thisisanurl",
		Telephone:      "0876756788",
		Password:       "password",
		Position:       "directeur",
		Positionsought: []string{"pos1", "pos2", "pos3", "pos4"},
		Industry:       "finance",
		Industrysought: []string{"ind1", "ind2", "ind3", "ind4"},
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {

	var err error
	if err != nil {
		return nil, err
	}
	users := []models.User{
		models.User{
			Name:           "nico",
			Lastname:       "abd",
			Email:          "nico_abd@gmail.com",
			Urlphoto:       "thisisanurl",
			Telephone:      "0798456678",
			Password:       "password",
			Position:       "directrice",
			Positionsought: []string{"pos1", "pos2", "pos3", "pos4"},
			Industry:       "marketing",
			Industrysought: []string{"ind1", "ind2", "ind3", "ind4"},
		},
		models.User{
			Name:           "babou",
			Lastname:       "jam",
			Email:          "babou_jam@gmail.com",
			Urlphoto:       "thisisanurl",
			Telephone:      "0876435212",
			Password:       "password",
			Position:       "sans",
			Positionsought: []string{"pos1", "pos2", "pos3", "pos4"},
			Industry:       "immobilier",
			Industrysought: []string{"ind1", "ind2", "ind3", "ind4"},
		},
	}
	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}
