package cmd

import (
	"biu-x.org/TikTok/models/user"
	"biu-x.org/TikTok/modules/config"
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

// CmdGen 子命令
var CmdGen = &cli.Command{ //nolint:typecheck
	Name:        "gen",
	Usage:       "gen gorm code",
	Description: `GEN: Friendly & Safer GORM powered by Code Generation.`,
	Action:      runGen,
}

func runGen(ctx *cli.Context) error { //nolint:typecheck
	g := gen.NewGenerator(gen.Config{
		OutPath: "./models/gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	url := fmt.Sprintf("%v:%v@(%v:%v)/%v", config.Get("mysql.username"), config.Get("mysql.password"), config.Get("mysql.host"), config.Get("mysql.port"), config.Get("mysql.database"))

	db, err := gorm.Open(mysql.Open(url + "?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		return err
	}
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(user.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, user.User{})

	// Generate the code
	g.Execute()
	return nil
}
