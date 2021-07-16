package service

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/page"
)

type ICommentService interface {
	GetComments(articleId int, page *page.IPage) (*pb.CommentInfo, error)
	GetReplies(commentId int, page *page.IPage) ([]*pb.Reply, error)
	AddComment(comment *pb.CsComment) error
	LikeComment(commentId int64, userId int64) error
	GetAdminComments(c *pb.CsCondition) (*pb.ScAdminComments, error)
	UpdateCommentStatus(status *pb.CsUpdateCommentStatus) error
	DeleteComments(ids *pb.CsDeleteComments) error
}
