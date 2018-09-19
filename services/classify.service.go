package services

import "xiaomatv.cn/models/Json"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述：分类服务层
 ***********************************************************
 */

 type ClassifyService interface {
 	 /*列表*/
	 ClassifyListsService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	 /*视频*/
	 ClassifyVideoService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	 /*用户*/
	 GetUsersClassifyService (userId int) (b []*Json.ConnotationBean)
	 /*视频*/
	 GetVideoClassifyService (userId int) (b []*Json.ConnotationBean)
	 /*兴趣*/
	 GetInterestService (userId int) (b []*Json.ConnotationBean)
	 /*精简*/
	 GetNewsListService (userId int) (b []*Json.ConnotationBean)
 }