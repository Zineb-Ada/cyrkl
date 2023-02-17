package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Slot struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint32    `gorm:"index" json:"user_id"`
	Dateandhours time.Time `gorm:"size:255" json:"dateandhours"`
	Lieu         string    `gorm:"size:255" json:"lieu"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Userc        User      `gorm:"ForeignKey:UserID" json:"user"`
}

func (s *Slot) PrepareSlot() {
	s.ID = 0
	// c.Dateandhours, _ = time.Parse("2006-01-02T15:04:05Z", c.Dateandhours.Format("2006-01-02T15:04:05Z"))
	s.Lieu = html.EscapeString(strings.TrimSpace(s.Lieu))
	s.Userc = User{}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}
func (s *Slot) ValidateSlot() error {

	if s.Dateandhours == (time.Time{}) {
		return errors.New("Required Date and Hours")
	}
	if s.Lieu == "" {
		return errors.New("Required Lieu")
	}
	if s.UserID < 1 {
		return errors.New("Required UserID")
	}
	return nil
}
func (s *Slot) SaveSlot(db *gorm.DB) (*Slot, error) {
	var err error
	err = db.Debug().Model(&Slot{}).Create(&s).Error
	if err != nil {
		return &Slot{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.Userc).Error
		if err != nil {
			return &Slot{}, err
		}
	}
	return s, nil
}

func (s *Slot) FindSlots(db *gorm.DB) (*[]Slot, error) {
	var err error
	slots := []Slot{}
	err = db.Debug().Model(&Slot{}).Limit(100).Find(&slots).Error
	if err != nil {
		return &[]Slot{}, err
	}
	if len(slots) > 0 {
		for i, _ := range slots {
			err := db.Debug().Model(&User{}).Where("id = ?", slots[i].UserID).Take(&slots[i].Userc).Error
			if err != nil {
				return &[]Slot{}, err
			}
		}
	}
	return &slots, nil
}

func (s *Slot) FindSlotByID(db *gorm.DB, pid uint64) (*Slot, error) {
	var err error
	err = db.Debug().Model(&Slot{}).Where("id = ?", pid).Take(&s).Error
	if err != nil {
		return &Slot{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.Userc).Error
		if err != nil {
			return &Slot{}, err
		}
	}
	return s, nil
}

func (s *Slot) FindSlotsByUserID(db *gorm.DB, userID uint64) (*[]Slot, error) {
	var err error
	slots := []Slot{}
	err = db.Debug().Model(&Slot{}).Where("user_id = ?", userID).Find(&slots).Error
	if err != nil {
		return nil, err
	}
	if len(slots) > 0 {
		for i, _ := range slots {
			err = db.Debug().Model(&User{}).Where("id = ?", slots[i].UserID).Take(&slots[i].Userc).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return &slots, nil
}

func (s *Slot) UpdateASlot(db *gorm.DB, cid uint64) (*Slot, error) {
	var err error
	var oldCalendar Slot
	err = db.Debug().Model(&Slot{}).Where("id = ?", cid).Find(&oldCalendar).Error
	if err != nil {
		return &Slot{}, err
	}
	oldCalendar.Dateandhours = s.Dateandhours
	oldCalendar.Lieu = s.Lieu
	oldCalendar.UpdatedAt = time.Now()
	err = db.Debug().Model(&oldCalendar).Update(oldCalendar).Error
	if err != nil {
		return &Slot{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.Userc).Error
		if err != nil {
			return &Slot{}, err
		}
	}
	return s, nil
}

func (s *Slot) DeleteASlot(db *gorm.DB, cid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Slot{}).Where("id = ? and user_id = ?", cid, uid).Take(&Slot{}).Delete(&Slot{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
