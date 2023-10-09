package main

import (
	"fmt"
	"log"
	"os"
	"siam-phishing/controller"
	"siam-phishing/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("load gagal")
		log.Fatalln(err)
	}
	
	sql, err := db.InitSQL()
	if err != nil {
		panic(err)
	}
	
	if err := db.AutoMigrate(sql); err != nil {
		panic(err)
	}

	// app:
	app := gin.Default()

	// load templates & statics
	app.LoadHTMLGlob("./views/*.go.tmpl")
	app.StaticFile("favicon.ico", "./public/favicon.ico")
	app.Static("./css", "./public/css/")
	app.Static("./img", "./public/img/")
	
	controller.LoadRoutes(app, sql)

	app.Run(":" + os.Getenv("PORT"))
}