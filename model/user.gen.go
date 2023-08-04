// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255);not null;uniqueIndex:uq_user_name,priority:1;comment:用户名" json:"name"` // 用户名
	Password  string `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`                              // 密码
	Signature string `gorm:"column:signature;type:varchar(255);not null;comment:个人简介" json:"signature"`                          // 个人简介
	Avatar    string `gorm:"column:avatar;type:varchar(255);not null;comment:头像" json:"avatar"`                                  // 头像
	/*
		用户个人页顶部大图

	*/
	BackgroundImage string `gorm:"column:background_image;type:varchar(255);comment:用户个人页顶部大图\n" json:"background_image"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
