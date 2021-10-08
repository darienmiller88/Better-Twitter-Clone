package models

import(
	"gorm.io/gorm"
)

//This Model will serve as the foundation for the Like, Retweet and QuoteTweet models. Since they all 
//reference the "User" model in exactly the same fashion, through the username column via foreign key, Simply
//have each model embed this model into the struct. This will establish a many - to - many relationship
//between the "users" and "tweets" table, with a user having many tweets, and a tweet having likes, 
//retweets, and quote tweets from many users.
type Base struct{
	gorm.Model
	UsernameRef string `gorm:"type:string;"`
	TweedIDRef  string `gorm:"type:string;"`

	User        User   `gorm:"foreignKey:UsernameRef; references:Username; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tweet       Tweet  `gorm:"foreignKey:TweedIDRef;  references:ID;       constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}