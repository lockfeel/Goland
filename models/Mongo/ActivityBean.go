package Mongo

import "gopkg.in/mgo.v2/bson"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述ActivityBean 数据类型
 ***********************************************************
 */
 type ActivityBean struct {
 	ObjectId 				bson.ObjectId	`bson:"_id"`
 	ActivityId				string			`bson:"activity_id"`
 	Title					string			`bson:"title"`
 	Bgpic					string			`bson:"bg_pic"`
 	Summary					string			`bson:"summary"`
 	Auther					string			`bson:"auther"`
 	TaskId					string			`bson:"task_id"`
 	GroupId					string			`bson:"group_id"`
 	Desc					string			`bson:"desc"`
 	Sort					int32			`bson:"sort"`
 	IsReward				string			`bson:"is_reward"`
 	CoverPic				string			`bson:"cover_pic"`
 	StartTime				string			`bson:"start_time"`
 	EndTime					string			`bson:"end_time"`
 	UserIds					[]*string		`bson:"user_ids"`
	BurnIds                 []*string		`bson:"burn_ids"`
 	VideoIds				[]*string		`bson:"video_ids"`
 	GroupTitle				string			`bson:"group_title"`
 	GroupPic				string			`bson:"group_pic"`
 	Rule					string			`bson:"rule"`
 }