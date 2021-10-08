package controllers

import (
	"TwitterClone/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{
	db *gorm.DB
}

func (u *UserController) Init(routerGroup *gin.RouterGroup, db *gorm.DB){
	u.db = db
	
	routerGroup.GET("/", u.getUsers)
	routerGroup.POST("/", u.postUser)
}

func (u *UserController) getUsers(c *gin.Context, ){
	result := []models.User{}
		
 	u.db.Find(&result)

	c.IndentedJSON(http.StatusOK, result)
}

func (u *UserController) postUser(c *gin.Context){
	user := models.User{}

	if err := c.ShouldBind(&user); err != nil{
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	if err := user.Validate(); err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} 

	result := u.db.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}