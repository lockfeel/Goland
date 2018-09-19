package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ThemeCollectionBean数据类型
 ***********************************************************
 */

 type ThemeCollectionBean struct {
	ObjectId 		bson.ObjectId	`bson:"_id"`
	ThemeId			string			`bson:"theme_id"`
	UserId			string			`bson:"user_id"`
	Type 			string			`bson:"type"`
	Action			string			`bson:"action"`
	CreateTime		int32			`bson:"create_time"`
 }
