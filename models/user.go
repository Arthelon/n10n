package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int64 `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) Save() error {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passBytes)
	err = db.Save(&user).Error
	return err
}

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}


func GetUserByUsername(username string) (*User, error) {
	user := User{}
	err := db.Where("username = ?", username).First(&user, &User{}).Error
	return &user, err
}

func GetUserById(id int64) (*User, error) {
	user := User{}
	err := db.Where("id = ?", id).First(&user, &User{}).Error
	return &user, err
}