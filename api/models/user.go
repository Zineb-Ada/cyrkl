package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Lastname  string `gorm:"size:255;not null" json:"lastname"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Urlphoto  string `gorm:"size:255" json:urlphoto"`
	Telephone string `gorm:"size:20;unique" json:"telephone"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	// EmailVerifiedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"email_verified_at,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// type User struct {
	// 	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	// 	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	// 	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	// 	Password  string    `gorm:"size:100;not null;" json:"password"`
	// 	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// 	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ID Name Lastname Email Urlphoto Telephone Password EmailVerifiedAt CreatedAt UpdatedAt

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Urlphoto = html.EscapeString(strings.TrimSpace(u.Urlphoto))
	u.Telephone = html.EscapeString(strings.TrimSpace(u.Telephone))
	// u.EmailVerifiedAt = time.Now()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Lastname == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Telephone == "" {
			return errors.New("Required Telphone")
		}
		if u.Urlphoto == "" {
			return errors.New("Required Urlphoto")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "createuser":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Lastname == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Lastname == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Telephone == "" {
			return errors.New("Required Telphone")
		}
		if u.Urlphoto == "" {
			return errors.New("Required Urlphoto")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":      u.Name,
			"lastname":  u.Lastname,
			"email":     u.Email,
			"urlphoto":  u.Urlphoto,
			"telephone": u.Telephone,
			"password":  u.Password,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
