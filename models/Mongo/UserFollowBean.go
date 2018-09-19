package Mongo

import "gopkg.in/mgo.v2/bson"
/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：UserFollowedBean数据类型
 ***********************************************************
 */
type UserFollowBean struct {
	ObejctId  		bson.ObjectId		`bson:"_id"`
	Id 				string				`bson:"id"`
	UserId			string				`bson:"user_id"`
	FollowedId		string				`bson:"followed_id"`
	Action			string				`bson:"action"`
	Status			string				`bson:"status"`
	ModifyDate		string				`bson:"modify_time"`
	CreateTime		string 				`bson:"create_time"`
}