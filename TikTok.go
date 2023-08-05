package main

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/service/auth"
	"biu-x.org/TikTok/service/user"
	"github.com/gin-gonic/gin"
)

const version = "v0.1"

func main() {
	config.Init()
	db.Init()
	r := gin.Default()
	r.POST("/login", user.Login)
	r.POST("/signup", user.Signup)
	r.GET("/userinfo", auth.RequireAuth(), user.UserInfo)
	r.Run()
	// app := cmd.NewApp()
	// app.Name = "TikTok"
	// app.Usage = "TikTok Server"
	// app.Description = "A TikTok Server Written in Go"
	// app.Version = version

	// err := app.Run(os.Args)
	// if err != nil {
	// 	log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	// }
}
