package dao

import (
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/dal/query"
)

// 创建评论
func CreateComment(comment *model.Comment) (err error) {
	c := query.Comment
	err = c.Create(comment)
	return err
}

// 通过视频ID获得评论数组
func GetCommentByVideoID(videoID int64) (comments []*model.Comment, err error) {
	c := query.Comment
	comments, err = c.Where(c.VideoID.Eq(videoID)).Find()
	return comments, err
}

// 通过评论ID删除评论
func DeleteCommentByID(id int64) (err error) {
	c := query.Comment
	_, err = c.Where(c.ID.Eq(id)).Delete()
	return err
}

// 通过视频 ID 获取评论总量
func GetCommentCountByVideoID(id int64) (int64, error) {
	c := query.Comment
	count, err := c.Where(c.VideoID.Eq(id)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
