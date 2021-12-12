package models

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

//Model to define how a twitter user will look
type User struct {
	gorm.Model
	Username    string `json:"username"     gorm:"type:string; size:20;  unique; not null"`
	DisplayName string `json:"display_name" gorm:"type:string; size:30;  not null"`
	Bio         string `json:"bio"          gorm:"type:string; size:160;"`
	Password    string `json:"password"     gorm:"type:string; not null"`
	Email       string `json:"email"        gorm:"type:string; not null"`
}


func (u User) Validate() error {
	return validation.ValidateStruct(&u,

		//The username field is required, and must be in the range of 5 chars to 20.
		validation.Field(&u.Username, validation.Required, validation.Length(5, 20)),

		//The password field must be at least 6 characters, and contain a lowercase and uppercase character, 
		//as well as at least one number.
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(6, 0),
			validation.Match(regexp.MustCompile("[a-z]")).Error("Password must contain lowercase letters"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("Password must contain uppercase letters"),
			validation.Match(regexp.MustCompile("[0-9]")).Error("Password must contain at least one number"),
		),

		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
