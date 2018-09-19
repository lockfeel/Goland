package services

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述：活动服务层
 ***********************************************************
 */

 type ActivityService interface {
 	 /*圈友秀秀*/
	 MemberShowService(userId,atype,pageNum,pageSize int) ()
	 /*活动列表*/
	 GetListsService(userId,atype,pageNum,pageSize int) ()
	 /*活动视频*/
	 GetVideosService(activityId,userId,atype,pageNum,pageSize int) ()
	 /*周排行榜*/
	 VideoRankService(userId,atype,pageNum,pageSize int) ()
	 /*检测状态*/
	 CheckPartakeInService(activityId,userId int) ()

 }
