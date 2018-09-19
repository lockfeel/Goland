package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述AccusationBean 数据类型
 ***********************************************************
 */

 type AccusationBean struct {
 	ObjectId 		 bson.ObjectId  `bson:"_id"`
 	UserId			 string			`bson:"user_id"`
 	ThemeId			 int32			`bson:"theme_id"`
 	Type			 int32			`bson:"type"`
 	Nature			 int32			`bson:"nature"`
 	OperationType    int32			`bson:"operation"`
 	CreateTime		 int32			`bson:"create_time"`
 }