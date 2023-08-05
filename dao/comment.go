package dao

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
)

var c = query.Comment

func CreateComment(comment *model.Comment) (err error) {
	err = c.Create(comment)
	return err
}

func GetCommentByVideoID(videoID int64) (comments []*model.Comment, err error) {
	comments, err = c.Where(c.VideoID.Eq(videoID)).Find()
	return comments, err
}

func DeleteCommentByID(id int64) (err error) {
	_, err = c.Where(c.ID.Eq(id)).Delete()
	return err
}

