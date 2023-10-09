package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(app *gin.Engine, db *gorm.DB) {
	// controllers
	home := NewHomeController(db)
	_ = NewUserController(db)

	// routes
	app.GET("/", home.IndexController)
	
}