package services

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018.08.13
 **---------------------------------------------------------
 **描述：用户互动服务层
 ***********************************************************
 */

 type InteractiveService interface {
 	/*评论列表*/
 	GetCommentService() ()
 	/*点赞列表*/
 	GetThumbService() ()
 	/*关注用户*/
 	GetFollowUserService() ()
 	/*分享用户*/
 	GetShareUserService() ()
 	/*收藏用户*/
 	GetCollectUserService() ()
 	/*点赞用户*/
 	GetThumbCommentUserService() ()
 }
