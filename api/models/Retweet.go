package models

import(
)

//This model will keep track of the number of retweets a particular tweet has by storing the name of the
//user retweeting, and the id of the tweet they retweeted.
type Retweet struct{
	Base `gorm:"embedded"`
}