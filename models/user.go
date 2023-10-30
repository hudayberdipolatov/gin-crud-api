package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// create a user

func CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// get users

func GetUsers(db *gorm.DB, user *[]User) (err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by id

func GetUser(db *gorm.DB, user *User, id int) (err error) {
	err = db.Where("id = ?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

// update user

func UpdateUser(db *gorm.DB, user *User) (err error) {
	err = db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// delete user
func DeleteUser(db *gorm.DB, user *User, id int) (err error) {
	err = db.Where("id =? ", id).Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
