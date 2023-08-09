package user

import (
	"biu-x.org/TikTok/dal/query"
	"fmt"
)

func SaveUser() {
	//q := query.Use(db.DB)
	//err := query.Use(db.DB).User.Create(&model.User{Name: "hiifong2", Password: "123456", Signature: "lazy", Avatar: "https://hiif.ong/logo.png", BackgroundImage: "https://hiif.ong/logo.png"})
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//}
	//query.SetDefault(db.DB)
	userDo, err := query.User.Where(query.User.Name.Eq("hiifong")).First()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(userDo)
}
