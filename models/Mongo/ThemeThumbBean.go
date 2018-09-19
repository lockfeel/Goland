package Mongo

import "gopkg.in/mgo.v2/bson"
/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ThemeThumbBean数据类型
 ***********************************************************
 */
 type ThemeThumbBean struct {
 	ObjectId			bson.ObjectId	`bson:"_id"`
 	UserId				string			`bson:"user_id"`
 	ThemeId				string			`bson:"theme_id"`
 	Type				string			`bson:"type"`
 	UserName			string			`bson:"user_name"`
 	UserNickname		string			`bson:"user_nickname"`
 	UserPic				string			`bson:"user_pic"`
 	UserSign			string			`bson:"user_sign"`
 	CreateTime			int32			`bson:"create_time"`
 	IsVip				string			`bson:"is_vip"`
 	Sts					string			`bson:"sts"`

 }
