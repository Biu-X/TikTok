package dao

import (
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"gorm.io/gorm"
)

// 创建User记录
func CreateUser(user *model.User) (err error) {
	u := query.User
	err = u.Create(user)
	return err
}

// 通过ID获得对应用户信息
func GetUserByID(id int64) (user *model.User, err error) {
	u := query.User
	if count, _ := u.Where(u.ID.Eq(id)).Count(); count == 0 {
		return &model.User{}, gorm.ErrRecordNotFound
	}
	user, err = u.Where(u.ID.Eq(id)).First()
	return user, err
}

// 通过用户名获得User
func GetUserByName(name string) (user *model.User, err error) {
	u := query.User
	user, err = u.Where(u.Name.Eq(name)).First()
	return user, err
}

// 通过用户名获得密码
func GetPasswordByName(name string) (password string, err error) {
	u := query.User
	user, err := u.Where(u.Name.Eq(name)).First()
	return user.Password, err
}

// 通过用户ID设置头像
func SetAvatarByID(id int64, avatarURL string) (err error) {
	u := query.User
	_, err = u.Where(u.ID.Eq(id)).Update(u.Avatar, avatarURL)
	return err
}

// 通过用户ID设置个人简介
func SetSignatureByID(id int64, signature string) (err error) {
	u := query.User
	_, err = u.Where(u.ID.Eq(id)).Update(u.Signature, signature)
	return err
}

// 通过用户ID设置密码
func SetPasswordByID(id int64, password string) (err error) {
	u := query.User
	_, err = u.Where(u.ID.Eq(id)).Update(u.Password, password)
	return err
}

// 通过用户ID设置个人页顶部大图
func SetBackgroundImageByID(id int64, backgroundImageURL string) (err error) {
	u := query.User
	_, err = u.Where(u.ID.Eq(id)).Update(u.BackgroundImage, backgroundImageURL)
	return err
}

// 通过用户ID修改用户名
func SetNameByID(id int64, name string) (err error) {
	u := query.User
	_, err = u.Where(u.ID.Eq(id)).Update(u.Name, name)
	return err
}

// 通过用户ID删除用户
func DeleteUserByID(id int64) (err error) {
	u := query.User
	_, err = u.Where(u.ID.Eq(id)).Delete()
	return err
}
