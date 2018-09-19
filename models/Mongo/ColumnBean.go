package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述ColumnBean 数据类型
 ***********************************************************
 */

 type ColumnBean struct {
 	ObjectId    bson.ObjectId		`bson:"_id"`
 	Id			int32				`bson:"id" json:"id"`
 	UserId		string				`bson:"user_id" json:"user_id"`
 	ColumnId	string				`bson:"column_id" json:"column_id"`
 	Title		string				`bson:"title" json:"title"`
 	Picture		string				`bson:"picture" json:"picture"`
 	Description string				`bson:"description" json:"description"`
 	CreateTime  int32				`bson:"create_time" json:"create_time"`
 	PlayNum		string				`bson:"play_num" json:"play_num"`
 	CommentNum  string				`bson:"comment_num" json:"comment_num"`
 	VideoNum    string				`bson:"video_num" json:"video_num"`
 	UserName    string				`bson:"user_name" json:"user_name"`
 	UserPic     string 				`bson:"user_pic" json:"user_pic"`
 	FollowNum	string 				`json:"follow_num"`
 	VideoList	[]*ColumnVideoBean  `bson:"videoList" json:"videoList"`

 }


