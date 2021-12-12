package models

import "time"

type ReplyBase struct{
	ID 		  int 		`pg:",pk"`
	R         int 		`pg:",unique"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt time.Time `pg:",soft_delete"`
 
	// Relied   *User  `pg:"fk:username"`
}