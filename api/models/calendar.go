package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Calendar struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint32    `gorm:"index" json:"user_id"`
	Dateandhours time.Time `gorm:"size:255" json:"dateandhours"`
	Lieu         string    `gorm:"size:255" json:"lieu"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Userc        User      `gorm:"ForeignKey:UserID" json:"user"`
}

func (c *Calendar) PrepareCalendar() {
	c.ID = 0
	// c.Dateandhours, _ = time.Parse("2006-01-02T15:04:05Z", c.Dateandhours.Format("2006-01-02T15:04:05Z"))
	c.Lieu = html.EscapeString(strings.TrimSpace(c.Lieu))
	c.Userc = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}
func (c *Calendar) ValidateCalendar() error {

	if c.Dateandhours == (time.Time{}) {
		return errors.New("Required Date and Hours")
	}
	if c.Lieu == "" {
		return errors.New("Required Lieu")
	}
	if c.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}
func (c *Calendar) SaveDate(db *gorm.DB) (*Calendar, error) {
	var err error
	err = db.Debug().Model(&Calendar{}).Create(&c).Error
	if err != nil {
		return &Calendar{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.Userc).Error
		if err != nil {
			return &Calendar{}, err
		}
	}
	return c, nil
}

func (c *Calendar) FindAllCalendar(db *gorm.DB) (*[]Calendar, error) {
	var err error
	calendar := []Calendar{}
	err = db.Debug().Model(&Calendar{}).Limit(100).Find(&calendar).Error
	if err != nil {
		return &[]Calendar{}, err
	}
	if len(calendar) > 0 {
		for i, _ := range calendar {
			err := db.Debug().Model(&User{}).Where("id = ?", calendar[i].UserID).Take(&calendar[i].Userc).Error
			if err != nil {
				return &[]Calendar{}, err
			}
		}
	}
	return &calendar, nil
}

func (c *Calendar) FindDateByID(db *gorm.DB, pid uint64) (*Calendar, error) {
	var err error
	err = db.Debug().Model(&Calendar{}).Where("id = ?", pid).Take(&c).Error
	if err != nil {
		return &Calendar{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.Userc).Error
		if err != nil {
			return &Calendar{}, err
		}
	}
	return c, nil
}

func (c *Calendar) FindCalendarsByUserID(db *gorm.DB, userID uint64) (*[]Calendar, error) {
	var err error
	calendars := []Calendar{}
	err = db.Debug().Model(&Calendar{}).Where("user_id = ?", userID).Find(&calendars).Error
	if err != nil {
		return nil, err
	}
	if len(calendars) > 0 {
		for i, _ := range calendars {
			err = db.Debug().Model(&User{}).Where("id = ?", calendars[i].UserID).Take(&calendars[i].Userc).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return &calendars, nil
}

func (c *Calendar) UpdateADate(db *gorm.DB, cid uint64) (*Calendar, error) {
	var err error
	var oldCalendar Calendar
	err = db.Debug().Model(&Calendar{}).Where("id = ?", cid).Find(&oldCalendar).Error
	if err != nil {
		return &Calendar{}, err
	}
	oldCalendar.Dateandhours = c.Dateandhours
	oldCalendar.Lieu = c.Lieu
	oldCalendar.UpdatedAt = time.Now()
	err = db.Debug().Model(&oldCalendar).Update(oldCalendar).Error
	if err != nil {
		return &Calendar{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.Userc).Error
		if err != nil {
			return &Calendar{}, err
		}
	}
	return c, nil
}

func (c *Calendar) DeleteADate(db *gorm.DB, cid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Calendar{}).Where("id = ? and user_id = ?", cid, uid).Take(&Calendar{}).Delete(&Calendar{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
