package models

import (
	"github.com/jinzhu/gorm"
)

type UsersAlgo struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32 `gorm:"index" json:"user_id"`
	UsersList uint32 `gorm:"type:integer" json:"users_list"`
	UserA     User   `gorm:"ForeignKey:UserID" json:"UserA"`
}

func (ua *UsersAlgo) GetUsersAlgoByUser(db *gorm.DB) (*[]UsersAlgo, error) {
	var err error
	users := []UsersAlgo{}
	err = db.Debug().Model(&UsersAlgo{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]UsersAlgo{}, err
	}
	return &users, err
}
