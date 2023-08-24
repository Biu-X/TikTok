package dao

import (
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
)

// 创建评论
func CreateComment(comment *model.Comment) (err error) {
	c := query.Comment
	err = c.Create(comment)
	return err
}

// 通过视频ID获得评论数组 按CreatedAt的Desc顺序排序
func GetCommentByVideoID(videoID int64) (comments []*model.Comment, err error) {
	c := query.Comment
	comments, err = c.Where(c.VideoID.Eq(videoID)).Order(c.CreatedAt.Desc()).Find()
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
