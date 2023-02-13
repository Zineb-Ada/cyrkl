package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type InvitationsReceived struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserSenderID   uint32    `gorm:"index" json:"user_sender_id"`
	UserReceiverID uint32    `gorm:"index" json:"user_receiver_id"`
	DateID         uint32    `gorm:"index" json:"date_id"`
	InProgress     bool      `gorm:"column:in_progress" json:"in_progress"`
	Accepted       bool      `gorm:"column:accepted" json:"accepted"`
	Refused        bool      `gorm:"column:refused" json:"refused"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Usersend       User      `gorm:"ForeignKey:UserSenderID" json:"usersend"`
	Userreceiv     User      `gorm:"ForeignKey:UserReceiverID" json:"userreceiv"`
	Date           Calendar  `gorm:"ForeignKey:DateID" json:"date"`
}

func (i *InvitationsReceived) PrepareReceivedInvit() {
	i.ID = 0
	i.InProgress = true
	i.Accepted = false
	i.Refused = false
	i.Usersend = User{}
	i.Userreceiv = User{}
	i.Date = Calendar{}
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
}

func (ir *InvitationsReceived) SaveReceivedInvitation(db *gorm.DB) (*InvitationsReceived, error) {
	var err error
	err = db.Debug().Model(&InvitationsReceived{}).Create(&ir).Error
	if err != nil {
		return &InvitationsReceived{}, err
	}
	if ir.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", ir.UserSenderID).Take(&ir.Usersend).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", ir.UserReceiverID).Take(&ir.Userreceiv).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		err = db.Debug().Model(&Calendar{}).Where("id = ?", ir.DateID).Take(&ir.Date).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", ir.Date.UserID).Take(&ir.Date.Userc).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		if ir.UserReceiverID != ir.Date.UserID {
			return &InvitationsReceived{}, errors.New("UserReceiverID and UserID are differente")
		}
	}
	return ir, nil
}

func (i *InvitationsReceived) FindAllInvitationsReceived(db *gorm.DB) (*[]InvitationsReceived, error) {
	var err error
	invitationsReceived := []InvitationsReceived{}
	err = db.Debug().Model(&InvitationsReceived{}).Limit(100).Find(&invitationsReceived).Error
	if err != nil {
		return &[]InvitationsReceived{}, err
	}
	if len(invitationsReceived) > 0 {
		for j, _ := range invitationsReceived {
			err = db.Debug().Model(&User{}).Where("id = ?", invitationsReceived[j].UserSenderID).Take(&invitationsReceived[j].Usersend).Error
			if err != nil {
				return &[]InvitationsReceived{}, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitationsReceived[j].UserReceiverID).Take(&invitationsReceived[j].Userreceiv).Error
			if err != nil {
				return &[]InvitationsReceived{}, err
			}
			err = db.Debug().Model(&Calendar{}).Where("id = ?", invitationsReceived[j].DateID).Take(&invitationsReceived[j].Date).Error
			if err != nil {
				return &[]InvitationsReceived{}, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitationsReceived[j].Date.UserID).Take(&invitationsReceived[j].Date.Userc).Error
			if err != nil {
				return &[]InvitationsReceived{}, err
			}
		}
	}
	return &invitationsReceived, nil
}

func (ir *InvitationsReceived) FindRdInvitationByID(db *gorm.DB, irid uint64) (*InvitationsReceived, error) {
	var err error
	err = db.Debug().Model(&InvitationsReceived{}).Where("id = ?", irid).Take(&ir).Error
	if err != nil {
		return &InvitationsReceived{}, err
	}
	if ir.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", ir.UserSenderID).Take(&ir.Usersend).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", ir.UserReceiverID).Take(&ir.Userreceiv).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		err = db.Debug().Model(&Calendar{}).Where("id = ?", ir.DateID).Take(&ir.Date).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", ir.Date.UserID).Take(&ir.Date.Userc).Error
		if err != nil {
			return &InvitationsReceived{}, err
		}
	}
	return ir, nil
}

// cette fonction elle sert les deux 
func (ir *InvitationsReceived) FindRdInvitByUserID(db *gorm.DB, userID uint64) (*[]InvitationsReceived, error) {
	var err error
	invitationreceived := []InvitationsReceived{}
	err = db.Debug().Model(&InvitationsReceived{}).Where("user_receiver_id = ?", userID).Find(&invitationreceived).Error
	// elle peut voir les deux cas user1 et user2 mais ca ne sert a rien
	// err = db.Debug().Model(&InvitationsReceived{}).Where("user_sender_id = ? OR user_receiver_id = ?", userID, userID).Find(&invitationreceived).Error
	if err != nil {
		return nil, err
	}
	if len(invitationreceived) > 0 {
		for i, _ := range invitationreceived {
			err = db.Debug().Model(&User{}).Where("id = ?", invitationreceived[i].UserSenderID).Take(&invitationreceived[i].Usersend).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitationreceived[i].UserReceiverID).Take(&invitationreceived[i].Userreceiv).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&Calendar{}).Where("id = ?", invitationreceived[i].DateID).Take(&invitationreceived[i].Date).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitationreceived[i].Date.UserID).Take(&invitationreceived[i].Date.Userc).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return &invitationreceived, nil
}

