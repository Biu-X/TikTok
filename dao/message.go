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

func GetMessageByID(id int64) (message *model.Message, err error) {
	f := query.Message
	message, err = f.Where(f.ID.Eq(id)).First()
	return message, err
}

// Order By CreatedAt ASC
func GetMessageByBoth(fromUserID int64, toUserID int64) (messages []*model.Message, err error) {
	f := query.Message
	messages, err = f.Where(f.FromUserID.Eq(fromUserID), f.ToUserID.Eq(toUserID)).Order(f.CreatedAt).Find()
	return messages, err
}
