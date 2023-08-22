package dao

import (
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/dal/query"
	"gorm.io/gorm"
)

// 创建聊天消息
func CreateMessage(message *model.Message) (err error) {
	m := query.Message
	err = m.Create(message)
	return err
}

// 通过消息ID获取对应消息
func GetMessageByID(id int64) (message *model.Message, err error) {
	f := query.Message
	message, err = f.Where(f.ID.Eq(id)).First()
	return message, err
}

// Order By CreatedAt ASC
func GetMessageByBoth(userA int64, userB int64) (messages []*model.Message, err error) {
	f := query.Message
	messages, err = f.Where(f.FromUserID.Eq(userA), f.ToUserID.Eq(userB)).Or(f.FromUserID.Eq(userB), f.ToUserID.Eq(userA)).Order(f.CreatedAt).Find()
	return messages, err
}

// 返回两个用户之间的最新消息
func GetLatestBidirectionalMessage(userA int64, userB int64) (message *model.Message, err error) {
	f := query.Message

	if count, _ := f.Where(f.FromUserID.Eq(userA), f.ToUserID.Eq(userB)).Count(); count == 0 {
		return &model.Message{}, gorm.ErrRecordNotFound
	}

	message, err =
		f.Where(f.FromUserID.Eq(userA), f.ToUserID.Eq(userB)).
			Or(f.FromUserID.Eq(userB), f.ToUserID.Eq(userA)).
			Last()
	return message, err
}
