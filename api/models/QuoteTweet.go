package models

import(
)

//This model will keep track of the number of quote tweets a particular tweet has by storing the name of the
//user quote tweeting, and the id of the tweet they quote tweeted.
type QuoteTweet struct{
	Base `gorm:"embedded"`
}