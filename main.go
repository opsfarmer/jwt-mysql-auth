package main

import (
	"github.com/gin-gonic/gin"
	"github.com/parikshitg/jwt-mysql-auth/handlers"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	// Routes
	r.GET("/", handlers.HomeHandler)
	r.GET("/login", handlers.LoginHandler)
	r.GET("/logout", handlers.LogoutHandler)
	r.GET("/register", handlers.RegisterHandler)
	r.GET("/welcome", handlers.WelcomeHandler)

	// Static Files
	r.Static("/css", "static/css")
	r.Static("/js", "static/js")

	r.Run()
}
