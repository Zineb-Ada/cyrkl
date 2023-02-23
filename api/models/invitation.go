package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// type InvitationStatus int

// const (
//     InProgress InvitationStatus = iota
//     Accepted
//     Refused
// )

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

func (in *Invitation) SaveInvitation(db *gorm.DB, uid uint64) (*Invitation, error) {
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

func (in *Invitation) FindInvitsReceivedByInvitedIDWithStatus(db *gorm.DB, userID uint64, status string) (*[]Invitation, error) {
	var err error
	invitations := []Invitation{}
	err = db.Debug().Model(&Invitation{}).Where("invited_id = ? AND statut = ?", userID, status).Find(&invitations).Error
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

func (in *Invitation) FindInvitsSendedByInviterIDWithStatus(db *gorm.DB, userID uint64, status string) (*[]Invitation, error) {
	var err error
	invitations := []Invitation{}
	err = db.Debug().Model(&Invitation{}).Where("inviter_id = ? AND statut = ?", userID, status).Find(&invitations).Error
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
	var err error
	var oldinvit Invitation
	err = db.Debug().Model(&Invitation{}).Where("id = ?", inid).Find(&Invitation{}).Error
	if db.Error != nil {
		return &Invitation{}, db.Error
	}
	// if oldinvit.invitedID
	if len(in.Statut) > 0 {
		oldinvit.Statut = in.Statut
		oldinvit.UpdatedAt = time.Now()
		err = db.Debug().Model(&Invitation{}).Where("id = ?", inid).Updates(oldinvit).Error
		if err != nil {
			return &Invitation{}, err
		}
	}
	err = db.Debug().Model(&Invitation{}).Where("id = ?", inid).Take(in).Error
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
