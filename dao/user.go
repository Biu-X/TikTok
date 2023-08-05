package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

var u = query.User

func initQueryUser() {
	u = query.User
}

func CreateUser(user *model.User) (err error) {
	err = u.Create(user)
	return err
}

func GetUserByID(id int64) (user *model.User, err error) {
	user, err = u.Where(u.ID.Eq(id)).First()
	return user, err
}

func GetUserByName(name string) (user *model.User, err error) {
	user, err = u.Where(u.Name.Eq(name)).First()
	return user, err
}

func GetPasswordByName(name string) (password string, err error) {
	user, err := u.Where(u.Name.Eq(name)).First()
	return user.Password, err
}

func SetAvatarByID(id int64, avatarURL string) (err error) {
	_, err = u.Where(u.ID.Eq(id)).Update(u.Avatar, avatarURL)
	return err
}

func SetSignatureByID(id int64, signature string) (err error) {
	_, err = u.Where(u.ID.Eq(id)).Update(u.Signature, signature)
	return err
}

func SetPasswordByID(id int64, password string) (err error) {
	_, err = u.Where(u.ID.Eq(id)).Update(u.Password, password)
	return err
}

func SetBackgroundImageByID(id int64, backgroundImageURL string) (err error) {
	_, err = u.Where(u.ID.Eq(id)).Update(u.BackgroundImage, backgroundImageURL)
	return err
}

func SetNameByID(id int64, name string) (err error) {
	_, err = u.Where(u.ID.Eq(id)).Update(u.Name, name)
	return err
}
