package main

import (
	"github.com/azkanurhuda/google-oauth2-golang/controllers"
	"github.com/azkanurhuda/google-oauth2-golang/initializers"
	"github.com/azkanurhuda/google-oauth2-golang/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var server *gin.Engine

func init() {
	initializers.ConnectDB()
	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Implement Google Oauth2 in Golang"})
	})

	authRouter := router.Group("/auth")
	authRouter.POST("/register", controllers.SignUpUser)
	authRouter.POST("/login", controllers.SignInUser)
	authRouter.GET("/logout", middleware.DeserializeUser(), controllers.LogoutUser)

	router.GET("/sessions/oauth/google", controllers.GoogleOAuth)
	router.GET("/users/me", middleware.DeserializeUser(), controllers.GetMe)

	router.StaticFS("/images", http.Dir("public"))
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Route Not Found"})
	})

	log.Print("server is run in 8000")
	log.Fatal(server.Run(":" + "8000"))
}
