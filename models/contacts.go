package models

import (
	"fmt"
	u "go-jwt-api/utils"

	"github.com/jinzhu/gorm"
)

// Contact struct for contact
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
}

// Validate contact data
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Name is required"), false
	}

	if contact.Email == "" {
		return u.Message(false, "Email is required"), false
	}

	if contact.UserID <= 0 {
		return u.Message(false, "User not found"), false
	}

	return u.Message(true, "success"), true
}

// Create a new contact
func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

// GetContact for user
func GetContact(id uint) *Contact {
	contact := &Contact{}

	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

// GetContacts for user
func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)

	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
