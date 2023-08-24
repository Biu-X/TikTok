package cmd

import (
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/gen"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/urfave/cli/v2"
)

// CmdGen 子命令
var CmdGen = &cli.Command{ //nolint:typecheck
	Name:        "gen",
	Usage:       "gen gorm code",
	Description: `GEN: Friendly & Safer GORM powered by Code Generation.`,
	Action:      runGen,
} //nolint:typecheck

func runGen(ctx *cli.Context) error { //nolint:typecheck
	config.Init()
	log.Init()
	db.Init()
	gen.Init()
	return nil
}
