package gen

import (
	"biu-x.org/TikTok/module/db"
	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func Init() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/query",
		Mode:         gen.WithDefaultQuery,
		ModelPkgPath: "./model",

		WithUnitTest: true,

		FieldNullable:     false,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db.DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}
