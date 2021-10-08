package models

import (
	"gorm.io/gorm"
)

//The message model will define private messages sent from one user to another, and vice versa.
type Message struct{
	gorm.Model
	MessageContent   string `json:"message_content" gorm:"type:string; not null"`
	UsernameSender   string `json:"sender_name"     gorm:"type:string;"`
	UsernameReceiver string `json:"receiver_name"   gorm:"type:string;"`

	//Both the sender and receiver columns will reference the username of the "users" table. This will be
	//a one - to - many relationship, with the user having many tweets.
	Sender           User   `gorm:"foreignKey:UsernameSender;   references:Username; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Receiver         User   `gorm:"foreignKey:UsernameReceiver; references:Username; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}