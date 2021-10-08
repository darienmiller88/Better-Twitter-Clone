package models

import(
)

//This model will keep track of the number of likes a particular tweet has by storing the name of the
//user who liked the tweet, and the id of the tweet they liked.
type Like struct{
	Base `gorm:"embedded"`
}