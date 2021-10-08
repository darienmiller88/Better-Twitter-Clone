package models

import (
	"gorm.io/gorm"
)

//This model defines a tweet, which will reference the user table, and allow tweets of up to 280 characters
type Tweet struct{
	gorm.Model
	TweetContent   string `json:"tweet_content" gorm:"type:string; size:280; not null"`
	UsernameRef    string `gorm:"type:string; size:20;"`
	NumLikes       uint   `gorm:"type:int; default:0"`
	NumRetweets    uint   `gorm:"type:int; default:0"`
	NumQuoteTweets uint   `gorm:"type:int; default:0"`
	IsAReply       uint   `gorm:"type:bit;"`

	//Tweet will reference the "Username" column of the "users" table as a foreign key, establishing a
	//to many relationship with that table. 
	User           User   `gorm:"foreignKey:UsernameRef; references:Username; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}