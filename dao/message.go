package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

func CreateMessage(message *model.Message) (err error) {
	m := query.Message
	err = m.Create(message)
	return err
}

