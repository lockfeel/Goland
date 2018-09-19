package dao_impl

import (
	"time"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"xiaomatv.cn/tools"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：BaseDaoImpl层
 ***********************************************************
 */

type BaseDaoImpl struct {}
/**是否被收藏了*/
func (this BaseDaoImpl) IsCollected(userId,themeId,atype string) bool {
	num := tools.Mongo().COUNT("xiaomatv","t_theme",bson.M{
		"type"	  : atype,
		"theme_id":themeId,
		"user_id" :userId,
		"action"  :"1",
	})
	return num > 0
}
/**是否被赞主题*/
func (this BaseDaoImpl) IsThumbedT(userId,themeId,atype string) bool {
	num := tools.Mongo().COUNT("xiaomatv","t_theme_thumb",bson.M{
		"type"	  	: atype,
		"theme_id"	: themeId,
		"user_id" 	: userId,
		"action"  	: "1",
	})
	return num > 0
}
/**是否被赞评论*/
func (this BaseDaoImpl) IsThumbedC(userId,themeId,atype,commentId string) bool{
	num := tools.Mongo().COUNT("xiaomatv","t_theme_comment_thumb",bson.M{
		"type"	  		: atype,
		"theme_id"		: themeId,
		"comment_id" 	: commentId,
		"user_id" 		: userId,
	})
	return num > 0
}
/**是否关注了谁*/
func (this BaseDaoImpl) IsFollowed(userId,followedId string ) bool{
	num := tools.Mongo().COUNT("xiaomatv","t_user_follow",bson.M{
		"followed_id": followedId,
		"user_id"    : userId,
		"action"     : "1",
	})
	return num > 0
}
/**是否加入了圈*/
func (this BaseDaoImpl) IsFollowCircle(userId,circleId,atype string) bool {
	num := tools.Mongo().COUNT("xiaomatv","t_follow_label",bson.M{
		"type"	  : atype,
		"label_id": circleId,
		"user_id" : userId,
		"action"  : "1",
	})
	return num > 0
}
/**是否被忽略了*/
func (this BaseDaoImpl) IsElide(userId,themeId string,atype int) bool{
	num := tools.Mongo().COUNT("xiaomatv","t_accusation",bson.M{
		"type"	  : atype,
		"theme_id":themeId,
		"user_id" :userId,
	})
	return num > 0
}
/**时间戳转换*/
func(this BaseDaoImpl) EncodeTime(s string) (str string){
	t := time.Now()
	var timeStr string = ""
	var num int64
	num = t.Unix()
	if x,error := strconv.Atoi(s);error != nil{
		return timeStr
	} else {
		num = num - int64(x)
		if num < 60 {
			timeStr = strconv.Itoa(int(num / 1)) + "秒前"
		} else if num >= 60 && num < 3600 {
			timeStr = strconv.Itoa(int(num / 60)) + "分钟前"
		} else if num >= 3600 && num < 86400 {
			timeStr = strconv.Itoa(int(num / 3600)) + "小时前"
		} else if num > 86400 && num < 86400*7 {
			timeStr = strconv.Itoa(int(num / 86400)) + "天前"
		}else if num >= 86400*7 {
			timeStr = time.Unix(int64(x), 0).Format("2006-01-02 03:04:05")
		}
	}
	return timeStr
}
