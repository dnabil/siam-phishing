package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HomeController struct {
	sql *gorm.DB
}

func NewHomeController(sql *gorm.DB) *HomeController {
	return &HomeController{sql: sql}
}

func (home *HomeController) IndexController(c *gin.Context){
	c.HTML(http.StatusOK, "index.go.tmpl", nil)
}