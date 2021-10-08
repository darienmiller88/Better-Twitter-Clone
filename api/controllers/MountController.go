package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type MountController struct{

}

func (m *MountController) Init(route *gin.RouterGroup, db *gorm.DB){
	userController := UserController{}
	tweetController := TweetController{}

	userController.Init(route.Group("/users"), db)
	tweetController.Init(route.Group("/tweets"), db)
}