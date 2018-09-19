package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述FollowLabelBean 数据类型
 ***********************************************************
 */
type FollowLabelBean struct {
	ObjectId		bson.ObjectId		`bson:"_id"`
	Id 				int32				`bson:"id"`
	UserId          string				`bson:"user_id"`
	LabelId         string				`bson:"label_id"`
	UnRead			string				`bson:"unread"`
	IsLook			string				`bson:"is_look"`
	IsPush			string				`bson:"is_push"`
	Action			string				`bson:"action"`
	Status          string				`bson:"status"`
	CreateTime		string				`bson:"create_time"`
	Modifydate		string				`bson:"modifydate"`
	Remark			string				`bson:"remark"`
}