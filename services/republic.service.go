package services

import "xiaomatv.cn/models/Json"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述：主页服务层
 ***********************************************************
 */
type RepublicService interface {
	/*发现*/
	DiscoverService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	/*推荐*/
	RecommendService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	/*快讯*/
	NewsFalshService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	/*小圈*/
	XMCircleService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	/*关注*/
	FollowedService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	/*沉浸*/
	ImmerseService (userId,atype,pageNum,pageSize int) (b []*Json.ConnotationBean)
	/*搜索*/
	SearchService (userId,atype,pageNum,pageSize int,code string) (b []*Json.ConnotationBean)
	/*活动*/
	XMCActivtyService (userId,atype,pageNum,pageSize int,code string) (b []*Json.ConnotationBean)
}

