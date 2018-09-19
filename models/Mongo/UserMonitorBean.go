package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：UserMonitorBean数据类型
 ***********************************************************
 */
 type  UserMonitorBean struct {
 	ObjectId	bson.ObjectId	`bson:"_id"`
 	UserId		string 			`bson:"user_id"`
 	NUM			string 			`bson:"NUM"`
 	Type 		string 			`bson:"Type"`
 }
