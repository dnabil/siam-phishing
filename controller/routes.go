package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(app *gin.Engine, db *gorm.DB) {
	// controllers
	home := NewHomeController(db)
	user := NewUserController(db)


	// routes
	app.NoRoute(func(c *gin.Context) { c.Redirect(http.StatusFound, "/") }) // niru behaviour siam

	app.GET("/", home.IndexController)
	app.GET("/index.php", home.IndexController)
	app.POST("/index.php", user.Login)
}