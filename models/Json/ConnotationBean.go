package Json
/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ConnotationBean数据类型
 ***********************************************************
 */
type ConnotationBean struct{
	Id 				int32		`json:"id"`
	UserId  		string		`json:"user_id"`
	ColumnId 		string		`json:"column_id"`
	Title			string		`json:"title"`
	SourceUrl		string		`json:"source_url"`
	CoverPic		string		`json:"cover_pic"`
	Type            string		`json:"type"`
	Status          string		`json:"status"`
	Nature			string		`json:"nature"`
	PlayCount		string		`json:"play_count"`
	ThumbCount		string		`json:"thumb_count"`
	CommentCount    string		`json:"comment_count"`
	ShareCount		string		`json:"show_count"`
	VideoDuration   string		`json:"video_duration"`
	UserRole		string		`json:"user_role"`
	UserPic			string		`json:"user_pic"`
	UserNickname	string		`json:"user_nickname"`
	Width			string		`json:"width"`
	Height			string		`json:"height"`
	CreateTime		string		`json:"create_time"`
	Grade			string		`json:"grade"`
	ResourceType	string		`json:"resource_type"`
	IsShowFounder	string		`json:"is_show_Founder"`
	IsFounder		string		`json:"is_founder"`
	IsAuther		string		`json:"is_auther"`
	IsFollowed		string		`json:"is_followed"`
	IsVip			string 		`json:"is_vip"`
	IsThumbed		string		`json:"is_thumbed"`
	TimeDescribe	string		`json:"time_describe"`
	CommentList		[]*CommentBean		`json:"comment_list"`

}
