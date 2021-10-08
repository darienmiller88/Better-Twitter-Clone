package models

import(
	"gorm.io/gorm"
)

//This model will create a one - to - many relationship with the Tweet model through the ID column. 
//It will store the ID of the tweet the user replied with, and the ID of the tweet the user replied to.
type Reply struct{
	gorm.Model
	IDReplier int    `gorm:"type:int;"`
	IDReplied int    `gorm:"type:int;"`

	Replier   Tweet `gorm:"foreignKey:IDReplier; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Replied   Tweet `gorm:"foreignKey:IDReplied; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}