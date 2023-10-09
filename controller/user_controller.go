package controller

import "gorm.io/gorm"

type UserController struct {
	sql *gorm.DB
}

func NewUserController(sql *gorm.DB) *UserController {
	return &UserController{sql: sql}
}