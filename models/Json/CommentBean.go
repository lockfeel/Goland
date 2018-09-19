package Json
/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：CommentBean数据类型
 ***********************************************************
 */
type CommentBean struct {
	Id 				string			`json:"id"`
	ThemeId			string			`json:"theme_id"`
	UserId			string			`json:"user_id"`
	Type           	string			`json:"type"`
	CommentContent 	string			`json:"commentContent"`
	ThumbCount		string			`json:"thumb_count"`
	RequestId		string			`json:"request_id"`
	RequestNum		string			`json:"request_num"`
	CreateTime		string			`json:"create_time"`
	UserNickname	string			`json:"user_nickname"`
	UserPic			string			`json:"user_pic"`
	IsThumbed		string			`json:"is_thumbed"`
	Depth			string			`json:"depth"`
	IsVip			string			`json:"is_vip"`
	TimeDescribe	string			`json:"time_describe"`
	IsFollowed		string			`json:"is_followed"`
	Images			[]*ImagesBean	`json:"images"`
	CommentType		string			`json:"comment_type"`
}