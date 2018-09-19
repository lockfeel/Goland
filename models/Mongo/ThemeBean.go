package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ResourceTabBean数据类型
 ***********************************************************
 */
type ThemeBean struct{
	ObjectId 		bson.ObjectId			`bson:"_id"`
	Id 				int32					`bson:"id"`
	OldId			string					`bson:"old_id"`
	ColumnId		string					`bson:"column_id"`
	UserId  		string					`bson:"user_id"`
	UserRole 		string					`bson:"user_role"`
	Title			string					`bson:"title"`
	SourceUrl		string					`bson:"source_url"`
	CoverPic		string					`bson:"cover_pic"`
	Type            string					`bson:"type"`
	Status			string					`bson:"status"`
	Nature			string					`bson:"nature"`
	PlayCount		string					`bson:"play_count"`
	ThumbCount		string					`bson:"thumb_count"`
	CommentCount    string					`bson:"comment_count"`
	ShareCount		string					`bson:"share_count"`
	VideoDuration   string					`bson:"video_duration"`
	UserPic			string					`bson:"user_pic"`
	UserNickname	string					`bson:"user_nickname"`
	Width			string					`bson:"width"`
	Height			string					`bson:"height"`
	CreateTime		string					`bson:"create_time"`
	GroupUserId		string					`bson:"group_user_id"`
	Images 			string					`bson:"images"`
	ForwardData		ForwardDataBean			`bson:"forward_data"`
	IsArticleVideo			string			`bson:"is_article_video"`
	ArticleOriginalTitle	string			`bson:"article_orginal_title"`
	ArticleOriginalUrl		string			`bson:"article_orginal_url"`
	GroupIsFollowed			string			`bson:"group_is_followed"`
	GroupId					string			`bson:"group_id"`
	GroupPicture			string			`bson:"group_picture"`
	GroupTitle				string			`bson:"group_title"`
	IsOriginal				string			`bson:"is_original"`
	IsEssence				string			`bson:"is_essence"`
	IsSticked				string			`bson:"is_sticked"`
}
