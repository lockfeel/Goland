package services

import "xiaomatv.cn/models/Json"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：ConnotationService层
 ***********************************************************
 */
type ConnotationSerivce interface {
	/**推荐展示*/
	TopicShowSerivce (userId,aType,pageNum,pageSize int) (list []*Json.ConnotationBean)
	/**精选展示*/
	TopicChoiceSerivce(title string,userId,pageNum,pageSize int) (list []*Json.ConnotationBean)
	/**评论列表*/
	CommentListSerivce(themeId,commentId,userId,pageNum,pageSize int ) (list []*Json.CommentBean)
}