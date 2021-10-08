package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type TweetController struct{
	db *gorm.DB
}

func (t *TweetController) Init(routerGroup *gin.RouterGroup, db *gorm.DB){
	
}