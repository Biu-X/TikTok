package gen

import (
	"github.com/Biu-X/TikTok/module/db"
	"gorm.io/gen"
)

func Init() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/query",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		ModelPkgPath: "dal/model",

		WithUnitTest: true,

		FieldNullable:     false,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db.DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)

	g.GenerateModel("comment", gen.FieldJSONTag("created_at", "create_date"),
		gen.FieldJSONTag("deleted_at", "delete_date"))

	g.GenerateModel("message", gen.FieldJSONTag("created_at", "create_date"))

	g.GenerateModel("video", gen.FieldJSONTag("created_at", "publish_time"))

	// Generate the code
	g.Execute()
}
