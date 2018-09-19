package Mongo
/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述ColumnVideoBean 数据类型
 ***********************************************************
 */
 type ColumnVideoBean struct {
 	ColumnId	string		`bson:"column_id" json:"column_id"`
 	VideoId     string		`bson:"video_id"  json:"video_id"`
 	Title		string		`bson:"title" json:"title"`
 	VideoUrl    string		`bson:"video_url" json:"video_url"`
 	CoverPic    string		`bson:"cover_pic" json:"cover_pic"`
 	OwnnerId    string		`bson:"ownner_id" json:"ownner_id"`
 	VideoStyle  string		`bson:"video_style" json:"video_style"`
 	Width       string		`bson:"width" json:"width"`
 	Height      string		`bson:"height" json:"height"`
 	UserName    string		`bson:"user_name" json:"user_name"`
 	UserPic     string		`bson:"user_pic" json:"user_pic"`
 	PlayCount   string		`bson:"play_count" json:"play_count"`
 	PlayNum		string		`bson:"play_num" json:"play_num"`
 	CommentNum  string 		`bson:"comment_num" json:"comment_num"`
 }