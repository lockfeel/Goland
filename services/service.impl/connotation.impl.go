package service_impl

import (
	"xiaomatv.cn/models/Json"
	"xiaomatv.cn/daos"
	"xiaomatv.cn/models/Mongo"
	"strconv"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：ConnotationServiceImpl层
 ***********************************************************
 */

type ConnotationServiceImpl struct {
	Dao  daos.ConnotationDao `inject:""`
	Sys  daos.BaseDao		 `inject:""`
}

/**推荐展示*/
func (this ConnotationServiceImpl) TopicShowSerivce (userId,aType,pageNum,pageSize int) (list []*Json.ConnotationBean){
	data := this.Dao.TopicShow(userId,aType,pageNum,pageSize)
	result := make([]*Json.ConnotationBean,pageSize)
	for i:= 0;i<len(data);i++{
		muse := this.Decode(data[i],strconv.Itoa(userId))
		result[i] = &muse
	}
	return result
}
/**精选展示*/
func (this ConnotationServiceImpl) TopicChoiceSerivce(title string,userId,pageNum,pageSize int) (list []*Json.ConnotationBean){
	data := this.Dao.TopicChoice(title,userId,pageNum,pageSize)
	println(len(data))
	return nil
}
/**评论列表*/
func (this ConnotationServiceImpl) CommentListSerivce(themeId,commentId,userId,pageNum,pageSize int ) (list []*Json.CommentBean){
	data := this.Dao.CommentList(themeId,commentId,userId,pageNum,pageSize)
	println(data)
	return nil
}

/**对象转换*/
func (this ConnotationServiceImpl) Decode(a *Mongo.ThemeBean,userId string) (b Json.ConnotationBean){
	var bean Json.ConnotationBean
	bean.Id 	 		= a.Id
	bean.Type   		= a.Type
	bean.Title  		= a.Title
	bean.UserId 		= a.UserId
	bean.ColumnId 		= a.ColumnId
	bean.Status			= a.Status
	bean.SourceUrl		= a.SourceUrl
	bean.CoverPic		= a.CoverPic
	bean.Nature			= a.Nature
	bean.PlayCount		= a.PlayCount
	bean.ThumbCount		= a.ThumbCount
	bean.ShareCount		= a.ShareCount
	bean.CommentCount   = a.CommentCount
	bean.VideoDuration	= a.VideoDuration
	bean.UserPic		= a.UserPic
	bean.UserRole		= a.UserRole
	bean.UserNickname   = a.UserNickname
	bean.Width			= a.Width
	bean.Height			= a.Height
	bean.CreateTime		= a.CreateTime
	if len(a.GroupId) > 1 {
		bean.ResourceType   = "1"
	}else{
		bean.ResourceType   = "2"
	}
	if a.GroupUserId == a.UserId {
		bean.IsShowFounder = "1"
	}else{
		bean.IsShowFounder = "0"
	}
	if len(a.UserId) > 0 {
		bean.IsAuther = "1"
	}else{
		bean.IsAuther = "0"
	}
	if len(a.GroupUserId) > 0 {
		bean.IsFounder = "0"
	} else {
		bean.IsFounder = "1"
	}
	if this.Sys.IsFollowed(userId,a.UserId) {
		bean.IsFollowed = "1"
	} else {
		bean.IsFollowed = "0"
	}
	if this.Sys.IsThumbedT(userId,strconv.Itoa(int(a.Id)),a.Type) {
		bean.IsThumbed = "1"
	} else {
		bean.IsThumbed = "0"
	}
	if a.UserRole == "5" {
		bean.IsVip = "1"
	} else {
		bean.IsVip = "0"
	}
	bean.TimeDescribe  = this.Sys.EncodeTime(a.CreateTime)
	return bean
}


