package Mongo

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ForwardDataBean数据类型
 ***********************************************************
 */

 type ForwardDataBean struct {
 	Id 				string		`bson:"id"`
 	Type			string		`bson:"type"`
 	Title			string		`bson:"title"`
 	CoverPic		string		`bson:"cover_pic"`
 	Images 			string		`bson:"images"`
 	VideoDuraction	string		`bson:"video_duration"`
 	PlayCount		string		`bson:"play_count"`
 	SourceUrl		string		`bson:"source_url"`
 	IsArticleVideo	string		`bson:"is_article_video"`
 	GroupId			string		`bson:"group_id"`
 	GroupTitle		string		`bson:"group_title"`
 	UserId			string		`bson:"user_id"`
 	UserNickname	string		`bson:"user_nickname"`
 }
