package main

import (
	"log"
	"os"

	"github.com/Biu-X/TikTok/cmd"
)

const version = "v0.1"

func main() {
	app := cmd.NewApp()
	app.Name = "TikTok"
	app.Usage = "TikTok Server"
	app.Description = "A TikTok Server Written in Go"
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
