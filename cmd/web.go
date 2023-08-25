package cmd

import (
	"fmt"

	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/middleware/cache"
	"github.com/Biu-X/TikTok/module/oss"
	"github.com/Biu-X/TikTok/module/sensitive"
	"github.com/Biu-X/TikTok/router"
	"github.com/urfave/cli/v2"
)

// CmdWeb api 子命令
var CmdWeb = &cli.Command{ //nolint:typecheck
	Name:        "server",
	Usage:       "Start TikTok api server",
	Description: `Star TikTok api server`,
	Action:      runWeb, //nolint:typecheck
	Flags: []cli.Flag{
		&cli.StringFlag{ //nolint:typecheck
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runWeb(ctx *cli.Context) error { //nolint:typecheck
	defer func() {
		for k := range cache.Clients {
			err := cache.Clients[k].C.Close()
			if err != nil {
				fmt.Printf("close redis: %v", err)
				return
			}
		}
	}()
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	cache.Init()
	sensitive.Init()
	router.Init()
	return nil
}
