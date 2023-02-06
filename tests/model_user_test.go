package tests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/zineb-ada/cyrkl/api/models"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		ID:             1,
		Name:           "nico",
		Lastname:       "abd",
		Email:          "nic_ab@gmail.com",
		Telephone:      "0798456678",
		Password:       "password",
	}
	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Name, savedUser.Name)
	assert.Equal(t, newUser.Lastname, savedUser.Lastname)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Telephone, savedUser.Telephone)
}

func TestFindUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Name, user.Name)
	assert.Equal(t, foundUser.Lastname, user.Lastname)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Urlphoto, user.Urlphoto)
	assert.Equal(t, foundUser.Telephone, user.Telephone)
	assert.Equal(t, foundUser.Position, user.Position)
	assert.Equal(t, foundUser.Positionsought, user.Positionsought)
	assert.Equal(t, foundUser.Industry, user.Industry)
	assert.Equal(t, foundUser.Industrysought, user.Industrysought)
}

func TestUpdateAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userUpdate := models.User{
		ID:             1,
		Name:           "up",
		Lastname:       "date",
		Email:          "up_date@gmail.com",
		Urlphoto:       "thisisanurl",
		Telephone:      "0788456678",
		Password:       "password",
		Position:       "commercial",
		Positionsought: pq.StringArray{"pos1", "pos2", "pos3", "pos4"},
		Industry:       "boum",
		Industrysought: pq.StringArray{"ind1", "ind2", "ind3", "ind4"},
	}
	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Name, userUpdate.Name)
	assert.Equal(t, updatedUser.Lastname, userUpdate.Lastname)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Urlphoto, userUpdate.Urlphoto)
	assert.Equal(t, updatedUser.Telephone, userUpdate.Telephone)
	assert.Equal(t, updatedUser.Position, userUpdate.Position)
	assert.Equal(t, updatedUser.Positionsought, userUpdate.Positionsought)
	assert.Equal(t, updatedUser.Industry, userUpdate.Industry)
	assert.Equal(t, updatedUser.Industrysought, userUpdate.Industrysought)
}

func TestDeleteAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
