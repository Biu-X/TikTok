package cmd

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/gen"
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
	db.Init()
	gen.Init()
	return nil
}
