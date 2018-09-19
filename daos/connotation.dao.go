package daos

import "xiaomatv.cn/models/Mongo"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：ConnotationDao层
 ***********************************************************
 */
 type ConnotationDao interface {
 	/**推荐展示*/
 	TopicShow (userId,aType,pageNum,pageSize int) (list []*Mongo.ThemeBean)
 	/**精选展示*/
 	TopicChoice(title string,userId,pageNum,pageSize int) (list []*Mongo.ThemeBean)
 	/**评论列表*/
 	CommentList(videoId,commentId,userId,pageNum,pageSize int ) (list []*Mongo.ThemeCommentBean)
 }