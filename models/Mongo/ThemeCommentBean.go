package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ThemeCommentBean数据类型
 ***********************************************************
 */
 type ThemeCommentBean struct {
 	ObjectId			bson.ObjectId	`bson:"_id"`
 	Id					int32			`bson:"id"`
 	ThemeId				string			`bson:"theme_id"`
 	UserId				string			`bson:"user_id"`
 	Type        		string			`bson:"type"`
 	CommentComtent		string			`bson:"comment_content"`
 	ThumbCount			string			`bson:"thumb_count"`
 	RequestNum			string			`bson:"request_num"`
 	RequestId			string			`bson:"request_id"`
 	CreateTime			string			`bson:"create_time"`
 	UserNickname		string			`bson:"user_nickname"`
 	UserPic				string			`bson:"user_pic"`
 	RequestUserNickname string			`bson:"request_user_nickname"`
 	RequestUserId		string			`bson:"request_id"`
 	RequestUserPic		string			`bson:"request_user_pic"`
 	RequestContent		string			`bson:"request_content"`
 	IsThumbed			string			`bson:"is_thumbed"`
 	Depth				string			`bson:"depth"`
 	IsVip				string			`bson:"is_vip"`
 	TimeDescribe		string 			`bson:"time_describe"`
 }
