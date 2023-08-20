package main

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/router"
)

const version = "v0.1"

func main() {
	config.Init()
	log.Init()
	db.Init()
	router.Init()
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
