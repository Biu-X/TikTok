package dao

import (
	"time"

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

// 返回在某个时间前两个用户间的聊天记录 Order By CreatedAt ASC
func GetMessageByBoth(userA int64, userB int64, preMsgTime time.Time) (messages []*model.Message, err error) {
	f := query.Message
	messages, err = f.Where(f.CreatedAt.Gt(preMsgTime)).Where(f.FromUserID.Eq(userA), f.ToUserID.Eq(userB)).Or(f.FromUserID.Eq(userB), f.ToUserID.Eq(userA)).Find()
	//messages, err = f.Where(f.FromUserID.Eq(userA), f.ToUserID.Eq(userB)).Or(f.FromUserID.Eq(userB), f.ToUserID.Eq(userA)).Where(f.CreatedAt.Gt(preMsgTime)).Order(f.CreatedAt).First()
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

func GetEarliestTimeMessageByBoth(ownerID, targetID int64) (time.Time, error) {
	f := query.Message
	t, err := f.Select(f.CreatedAt).Where(f.FromUserID.Eq(ownerID), f.ToUserID.Eq(targetID)).Or(f.FromUserID.Eq(targetID), f.ToUserID.Eq(ownerID)).First()
	return t.CreatedAt, err
}
