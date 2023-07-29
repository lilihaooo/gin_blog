package models

import "gorm.io/gorm"

// CommentModel 评论表
type CommentModel struct {
	gorm.Model
	SubComments        []CommentModel `gorm:"foreignkey:ParentCommentID" json:"sub_comments"`     // 子评论列表
	ParentCommentModel *CommentModel  `gorm:"foreignkey:ParentCommentID" json:"comment_model"`    // 父级评论
	ParentCommentID    *uint          `gorm:"comment:父评论id" json:"parent_comment_id"`             // 父评论id
	Content            string         `gorm:"size:256;comment:评论内容" json:"content"`               // 评论内容
	DiggCount          int            `gorm:"size:8;default:0;comment:点赞数" json:"digg_count"`     // 点赞数
	CommentCount       int            `gorm:"size:8;default:0;comment:子评论数" json:"comment_count"` // 子评论数
	ArticleID          string         `gorm:"size:32;comment:文章id" json:"article_id"`             // 文章id
	User               UserModel      `json:"user"`                                               // 关联的用户
	UserID             uint           `gorm:"comment:评论的用户Id" json:"user_id"`                     // 评论的用户
}

func (m *CommentModel) TableName() string {
	// 自定义表名的逻辑
	return "comment"
}
