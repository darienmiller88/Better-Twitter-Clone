package models

import "time"

type Reply2 struct{
	ID 		  int 		`pg:",pk"`
	R         int 		
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt time.Time `pg:",soft_delete"`
 
	Replier   ReplyBase	`pg:"rel:has-one"`
	// Relied   *User  `pg:"fk:username"`
}