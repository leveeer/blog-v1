package impl

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

type CommentServiceImpl struct {
	wg sync.WaitGroup
}

func (c *CommentServiceImpl) DeleteComments(ids *pb.CsDeleteComments) error {
	panic("implement me")
}

func (c *CommentServiceImpl) UpdateCommentStatus(status *pb.CsUpdateCommentStatus) error {
	panic("implement me")
}

func (c *CommentServiceImpl) GetAdminComments(csCondition *pb.CsCondition) (*pb.ScAdminComments, error) {
	condition := "c.is_delete = ? "
	if csCondition.GetKeywords() != "" {
		condition = fmt.Sprintf(condition+"%s", " AND u.nickname LIKE ?")
	}
	comments, err := model.GetCommentsByConditionWithPage(condition,
		&page.IPage{Current: int(csCondition.GetCurrent()), Size: int(csCondition.GetSize())},
		csCondition.GetIsDelete(),
		"%"+csCondition.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	var commentSlice []*pb.ScComment
	for _, comment := range comments {
		countString, err := common.GetRedisUtil().HashGet(common.CommentLikeCount, strconv.Itoa(comment.ID))
		count, _ := strconv.Atoi(countString)
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
		if errors.Is(err, redis.Nil) {
			count = 0
		}
		commentSlice = append(commentSlice, &pb.ScComment{
			Id:             int64(comment.ID),
			Avatar:         comment.Avatar,
			Nickname:       comment.Nickname,
			ReplyNickname:  comment.ReplyNickname,
			ArticleTitle:   comment.ArticleTitle,
			CommentContent: comment.CommentContent,
			CreateTime:     comment.CreateTime,
			IsDelete:       int32(comment.IsDelete),
			LikeCount:      int64(count),
		})
	}

	commentsCount, err := model.GetCommentsCountByCondition(condition, csCondition.GetIsDelete(), "%"+csCondition.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}

	return &pb.ScAdminComments{
		CommentList: commentSlice,
		Count:       commentsCount,
	}, nil
}

func (c *CommentServiceImpl) LikeComment(commentId int64, userId int64) error {
	// 查询当前用户点赞过的评论id集合
	ids, err := common.GetRedisUtil().HashGet(common.CommentUserLike, strconv.Itoa(int(userId)))
	if err != nil && err != redis.Nil {
		return err
	}
	if ids == "" {
		//不存在
		var slice []int32
		slice = append(slice, int32(commentId))
		err = common.GetRedisUtil().HashSet(common.CommentUserLike, strconv.Itoa(int(userId)), common.Serialization(slice))
		if err != nil {
			return err
		}
		err = common.GetRedisUtil().HashIncrBy(common.CommentLikeCount, strconv.Itoa(int(commentId)), 1)
		if err != nil {
			return err
		}
	} else {
		//存在，说明已点赞，再次点赞则取消
		var commentLikeSet []int32
		err := json.Unmarshal([]byte(ids), &commentLikeSet)
		if err != nil {
			return err
		}
		isExist, index := common.SliceFind(commentLikeSet, int32(commentId))
		logging.Logger.Debug(isExist)
		if isExist {
			commentLikeSet = append(commentLikeSet[:index], commentLikeSet[index+1:]...)
			err = common.GetRedisUtil().HashSet(common.CommentUserLike, strconv.Itoa(int(userId)), common.Serialization(commentLikeSet))
			if err != nil {
				return err
			}
			err = common.GetRedisUtil().HashIncrBy(common.CommentLikeCount, strconv.Itoa(int(commentId)), -1)
			if err != nil {
				return err
			}
		} else {
			commentLikeSet = append(commentLikeSet, int32(commentId))
			err = common.GetRedisUtil().HashSet(common.CommentUserLike, strconv.Itoa(int(userId)), common.Serialization(commentLikeSet))
			if err != nil {
				return err
			}
			err = common.GetRedisUtil().HashIncrBy(common.CommentLikeCount, strconv.Itoa(int(commentId)), 1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewCommentServiceImpl() *CommentServiceImpl {
	return &CommentServiceImpl{}
}

func (c *CommentServiceImpl) GetReplies(commentId int, ipage *page.IPage) ([]*pb.Reply, error) {
	replies, err := model.GetRepliesByCommentId(ipage, []int64{int64(commentId)})
	if err != nil {
		return nil, err
	}
	var repliesSlice []*pb.Reply
	for _, reply := range replies {
		repliesSlice = append(repliesSlice, &pb.Reply{
			Id:             int32(reply.ID),
			ParentId:       int32(reply.ParentID),
			UserId:         int32(reply.UserID),
			Nickname:       reply.Nickname,
			Avatar:         reply.Avatar,
			WebSite:        reply.WebSite,
			ReplyId:        int32(reply.ReplyID),
			ReplyNickname:  reply.ReplyNickname,
			ReplyWebSite:   reply.ReplyWebSite,
			CommentContent: reply.CommentContent,
			CreateTime:     reply.CreateTime,
		})
	}
	return repliesSlice, err
}

func (c *CommentServiceImpl) AddComment(comment *pb.CsComment) error {
	c1 := &model.Comment{
		ArticleID:      int(comment.GetArticleId()),
		UserID:         int(comment.GetUserId()),
		CommentContent: comment.GetCommentContent(),
		CreateTime:     time.Now().Unix(),
		ReplyID:        int(comment.ReplyId),
		ParentID:       int(comment.ParentId),
		IsDelete:       0,
	}
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		err := model.AddComment(tx, c1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentServiceImpl) GetComments(articleId int, ipage *page.IPage) (*pb.CommentInfo, error) {
	condition := "article_id = ? and parent_id = ? and is_delete = ?"
	//查询文章评论量
	commentCount, err := model.GetCommentsCountByCondition(condition, articleId, 0, false)
	if err != nil {
		return nil, err
	}
	if commentCount == 0 {
		return nil, nil
	}
	//分页查询评论集合
	comments, err := model.GetCommentsAndUserInfo(ipage, condition, articleId, 0, false)
	if err != nil {
		return nil, err
	}
	var commentsSlice []*pb.Comment
	for _, comment := range comments {
		logging.Logger.Debug(comment)
		commentsSlice = append(commentsSlice, &pb.Comment{
			Id:             int32(comment.ID),
			UserId:         int32(comment.UserID),
			Nickname:       comment.Nickname,
			Avatar:         comment.Avatar,
			WebSite:        comment.WebSite,
			CommentContent: comment.CommentContent,
			CreateTime:     comment.CreateTime,
		})
	}
	//查询评论点赞数据
	likeCountMap, err := common.GetRedisUtil().HashGetAll(common.CommentLikeCount)
	if err != nil {
		return nil, err
	}
	//封装评论点赞量
	var commentIdSlice []int64
	for _, comment := range commentsSlice {
		commentIdSlice = append(commentIdSlice, int64(comment.Id))
		likeCount, _ := strconv.Atoi(likeCountMap[string(comment.Id)])
		comment.LikeCount = uint32(likeCount)
	}
	//根据评论id集合查询回复数据
	replies, err := model.GetReplies(commentIdSlice)
	if err != nil {
		return nil, err
	}
	var repliesSlice []*pb.Reply
	for _, reply := range replies {
		repliesSlice = append(repliesSlice, &pb.Reply{
			Id:             int32(reply.ID),
			ParentId:       int32(reply.ParentID),
			UserId:         int32(reply.UserID),
			Nickname:       reply.Nickname,
			Avatar:         reply.Avatar,
			WebSite:        reply.WebSite,
			ReplyId:        int32(reply.ReplyID),
			ReplyNickname:  reply.ReplyNickname,
			ReplyWebSite:   reply.ReplyWebSite,
			CommentContent: reply.CommentContent,
			CreateTime:     reply.CreateTime,
		})
	}
	//封装回复点赞量
	for _, reply := range repliesSlice {
		likeCount, _ := strconv.Atoi(likeCountMap[string(reply.Id)])
		reply.LikeCount = uint32(likeCount)
	}
	//根据评论id分组回复数据
	replyMap := make(map[int32][]*pb.Reply)
	for _, reply := range repliesSlice {
		repliesByGroup := append(replyMap[reply.ParentId], reply)
		replyMap[reply.ParentId] = repliesByGroup
	}
	//根据评论id查询回复量
	replyCountList, err := model.GetReplyCountByCommentId(commentIdSlice)
	replyCountMap := make(map[int32]int32)
	for _, reply := range replyCountList {
		replyCountMap[int32(reply.ParentID)] = int32(reply.ReplyCount)
	}
	//将分页回复数据和回复量封装进对应的评论
	for _, comment := range commentsSlice {
		comment.ReplyList = replyMap[comment.Id]
		comment.ReplyCount = uint32(replyCountMap[comment.Id])
	}

	return &pb.CommentInfo{
		CommentList: commentsSlice,
		Count:       uint32(commentCount),
	}, nil
}
