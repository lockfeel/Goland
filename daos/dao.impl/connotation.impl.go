package dao_impl

import (
	"xiaomatv.cn/models/Mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"xiaomatv.cn/conf"
	"strconv"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：ConnotationDaoImpl层
 ***********************************************************
 */

 type ConnotationDaoImpl struct {

 }

/**推荐展示*/
func (this ConnotationDaoImpl) TopicShow (userId,aType,pageNum,pageSize int) (result []*Mongo.ThemeBean){
	session,err := mgo.Dial(conf.MONGOURL)
	if err !=nil {
		return
	}
	any := session.DB("xiaomatv").C("t_theme")
	any.Find(bson.M{"type": strconv.Itoa(aType)}).Sort("-id").Skip((pageNum-1)*pageSize).Limit(pageSize).All(&result)
	if len(result) <=0 {
		return
	} else {
		return
	}
}
/**精选展示*/
func (this ConnotationDaoImpl) TopicChoice(title string,userId,pageNum,pageSize int) (list []*Mongo.ThemeBean){
	return nil
}
/**评论列表*/
func (this ConnotationDaoImpl) CommentList(videoId,commentId,userId,pageNum,pageSize int ) (list []*Mongo.ThemeCommentBean){
	return nil
}

