package models

import (
	"errors"
	"html"
	"strings"
	"tulip/backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Role     string `gorm:"size:12;not null" json:"role"`
	Active   bool   `gorm:"not null; default:true" json:"active"`
}

func (u *User) UpdateUser() (*User, error) {
	err := Db.Save(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) SaveUser() (*User, error) {
	err := Db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Role = html.EscapeString(strings.TrimSpace(u.Role))

	return
}

func (u *User) PrepareGive() {
	u.Password = "hidden"
}

func VerifyPassword(pwd, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := User{}
	err = Db.Model(User{}).Where("username=?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID, u.Role)

	if err != nil {
		return "", err
	}
	return token, nil

}

func GetUserByID(user_id uint) (User, error) {
	var u User

	if err := Db.First(&u, user_id).Error; err != nil {
		return u, errors.New("user not found")
	}
	u.PrepareGive()
	return u, nil
}

func GetUsers() ([]User, error) {
	var us []User
	if err := Db.Find(&us).Error; err != nil {
		return us, err
	}
	for i, _ := range us {
		us[i].PrepareGive()
	}
	return us, nil
}
