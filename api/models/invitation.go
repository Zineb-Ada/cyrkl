package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Invitation struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	InviterID uint32 `gorm:"index" json:"inviter_id"`
	InvitedID uint32 `gorm:"index" json:"invited_id"`
	SlotID    uint32 `gorm:"index" json:"slot_id"`
	Statut    string `gorm:"size:255" json:"statut"`
	// Statut    string    `gorm:"column:statut;type:enum('In progress', 'Accepted', 'Refused')" json:"statut"`
	Inviter   User      `gorm:"ForeignKey:InviterID" json:"inviter"`
	Invited   User      `gorm:"ForeignKey:InvitedID" json:"invited"`
	Slotd     Slot      `gorm:"ForeignKey:SlotID" json:"slot"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// const (
//     InvitationStatusInProgress = "In progress"
//     InvitationStatusAccepted   = "Accepted"
//     InvitationStatusRefused    = "Refused"
// )

func (in *Invitation) PrepareInvitation(action string) {
	if strings.ToLower(action) == "update" {
		in.ID = 0
		// in.Inviter = User{}
		// in.Invited = User{}
		// in.Slotd = Slot{}
		in.UpdatedAt = time.Now()
	}
	if strings.ToLower(action) == "create" {
		in.ID = 0
		in.Statut = "In progress"
		in.Inviter = User{}
		in.Invited = User{}
		in.Slotd = Slot{}
		in.CreatedAt = time.Now()
		in.UpdatedAt = time.Now()
	}
}

func (in *Invitation) SaveInvitation(db *gorm.DB) (*Invitation, error) {
	var err error
	err = db.Debug().Model(&Invitation{}).Create(&in).Error
	if err != nil {
		return &Invitation{}, err
	}
	if in.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", in.InviterID).Take(&in.Inviter).Error
		if err != nil {
			return &Invitation{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", in.InvitedID).Take(&in.Invited).Error
		if err != nil {
			return &Invitation{}, err
		}
		err = db.Debug().Model(&Slot{}).Where("id = ?", in.SlotID).Take(&in.Slotd).Error
		if err != nil {
			return &Invitation{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", in.Slotd.UserID).Take(&in.Slotd.Userc).Error
		if err != nil {
			return &Invitation{}, err
		}
		if in.InvitedID != in.Slotd.UserID {
			return &Invitation{}, errors.New("InvitedID and UserID are differente")
		}
	}
	return in, nil
}

func (in *Invitation) FindInvitations(db *gorm.DB) (*[]Invitation, error) {
	var err error
	invitations := []Invitation{}
	err = db.Debug().Model(&Invitation{}).Limit(100).Find(&invitations).Error
	if err != nil {
		return &[]Invitation{}, err
	}
	if len(invitations) > 0 {
		for i, _ := range invitations {
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].InviterID).Take(&invitations[i].Inviter).Error
			if err != nil {
				return &[]Invitation{}, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].InvitedID).Take(&invitations[i].Invited).Error
			if err != nil {
				return &[]Invitation{}, err
			}
			err = db.Debug().Model(&Slot{}).Where("id = ?", invitations[i].SlotID).Take(&invitations[i].Slotd).Error
			if err != nil {
				return &[]Invitation{}, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].Slotd.UserID).Take(&invitations[i].Slotd.Userc).Error
			if err != nil {
				return &[]Invitation{}, err
			}
		}
	}
	return &invitations, nil
}

func (in *Invitation) FindInvitatByID(db *gorm.DB, inid uint64) (*Invitation, error) {
	var err error
	err = db.Debug().Model(&Invitation{}).Where("id = ?", inid).Take(&in).Error
	if err != nil {
		return &Invitation{}, err
	}
	if in.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", in.InviterID).Take(&in.Inviter).Error
		if err != nil {
			return &Invitation{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", in.InvitedID).Take(&in.Invited).Error
		if err != nil {
			return &Invitation{}, err
		}
		err = db.Debug().Model(&Slot{}).Where("id = ?", in.SlotID).Take(&in.Slotd).Error
		if err != nil {
			return &Invitation{}, err
		}
		err = db.Debug().Model(&User{}).Where("id = ?", in.Slotd.UserID).Take(&in.Slotd.Userc).Error
		if err != nil {
			return &Invitation{}, err
		}
	}
	return in, nil
}

// cette fonction est pour Getinvitreceived pour Getinvitsended je fait le contraire
func (in *Invitation) FindInvitsReceivedByInvitedID(db *gorm.DB, userID uint64) (*[]Invitation, error) {
	var err error
	invitations := []Invitation{}
	err = db.Debug().Model(&Invitation{}).Where("invited_id = ?", userID).Find(&invitations).Error
	if err != nil {
		return nil, err
	}
	if len(invitations) > 0 {
		for i, _ := range invitations {
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].InviterID).Take(&invitations[i].Inviter).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].InvitedID).Take(&invitations[i].Invited).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&Slot{}).Where("id = ?", invitations[i].SlotID).Take(&invitations[i].Slotd).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].Slotd.UserID).Take(&invitations[i].Slotd.Userc).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return &invitations, nil
}

func (in *Invitation) FindInvitSendedByInviterID(db *gorm.DB, userID uint64) (*[]Invitation, error) {
	var err error
	invitations := []Invitation{}
	err = db.Debug().Model(&Invitation{}).Where("inviter_id = ?", userID).Find(&invitations).Error
	if err != nil {
		return nil, err
	}
	if len(invitations) > 0 {
		for i, _ := range invitations {
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].InviterID).Take(&invitations[i].Inviter).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].InvitedID).Take(&invitations[i].Invited).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&Slot{}).Where("id = ?", invitations[i].SlotID).Take(&invitations[i].Slotd).Error
			if err != nil {
				return nil, err
			}
			err = db.Debug().Model(&User{}).Where("id = ?", invitations[i].Slotd.UserID).Take(&invitations[i].Slotd.Userc).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return &invitations, nil
}

func (in *Invitation) UpdateInvit(db *gorm.DB, inid uint64) (*Invitation, error) {
	// var err error
	db = db.Debug().Model(&Invitation{}).Where("id = ?", inid).Take(&Invitation{}).UpdateColumns(
		map[string]interface{}{
			"statut":     in.Statut,
			"inviter_id": in.InviterID,
			"invited_id": in.InvitedID,
			"slot_id":    in.SlotID,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Invitation{}, db.Error
	}
	// var oldInvit Invitation
	// err = db.Debug().Model(&Invitation{}).Where("id = ?", inid).Find(&oldInvit).Error
	// if err != nil {
	// 	return &Invitation{}, err
	// }
	// oldInvit.Statut = in.Statut
	// // oldInvit.SlotID = in.SlotID
	// fmt.Printf("statut 159 : %s", in.Statut)
	// fmt.Printf("oldinvit 160 : %s", oldInvit.Statut)
	// oldInvit.UpdatedAt = time.Now()

	// // if oldInvit.InvitedID != oldInvit.Slotd.UserID {
	// // 	return &Invitation{}, errors.New("You are unauthorized for update Invition")
	// // }
	// // if in.SlotID != in.Slotd.ID {
	// // 	return &Invitation{}, errors.New("Incorrect Slot.ID, you are unauthorized for update Invition")
	// // }
	// // && in.SlotID == oldInvit.Slotd.ID
	// err = db.Debug().Model(&User{}).Where("id = ?", in.Slotd.UserID).Take(&in.Slotd.Userc).Error
	// if err != nil {
	// 	return &Invitation{}, err
	// }
	// if in.InvitedID == in.Slotd.UserID {
	// 	err = db.Debug().Model(&oldInvit).Update(oldInvit).Error
	// 	if err != nil {
	// 		return &Invitation{}, errors.New("Incorrect InvitedID, you are unauthorized for update Invition")
	// 	}
	// } else {
	// 	return nil, errors.New("Incorrect InvitedID, you are unauthorized for update Invition")
	// }

	// if in.InvitedID != in.Slotd.UserID || in.InvitedID == in.InviterID {
	// 	return &Invitation{}, errors.New("You are unauthorized for update Invition")
	// }
	// if in.ID != 0 {
	// err = db.Debug().Model(&User{}).Where("id = ?", in.InviterID).Find(&in.Inviter).Error
	// if err != nil {
	// 	return &Invitation{}, err
	// }
	// err = db.Debug().Model(&User{}).Where("id = ?", in.InvitedID).Find(&in.Invited).Error
	// if err != nil {
	// 	return &Invitation{}, err
	// }
	// err = db.Debug().Model(&Slot{}).Where("id = ?", in.SlotID).Find(&in.Slotd).Error
	// if err != nil {
	// 	return &Invitation{}, err
	// }
	// err = db.Debug().Model(&User{}).Where("id = ?", in.Slotd.UserID).Find(&in.Slotd.Userc).Error
	// if err != nil {
	// 	return &Invitation{}, err
	// }
	// }
	return in, nil
}

func (in *Invitation) DeleteInvit(db *gorm.DB, inid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Invitation{}).Where("id = ? and user_id = ?", inid, uid).Take(&Invitation{}).Delete(&Invitation{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Invitation not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// func (in *Invitation)
