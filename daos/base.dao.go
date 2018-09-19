package daos
/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-01
 **---------------------------------------------------------
 **描述：BaseDao层
 ***********************************************************
 */
 type BaseDao interface {
 	/**是否收藏*/
 	IsCollected(userId,themeId,atype string) bool
 	/**是否点赞*/
 	IsThumbedT(userId,themeId,atype string) bool
 	/**是否点赞评论*/
 	IsThumbedC(userId,themeId,atype,commentId string) bool
 	/**是否关注*/
 	IsFollowed(userId,followedId string ) bool
 	/**是否关注分类*/
	IsFollowCircle(userId,circleId,atype string) bool
	/**是否不感兴趣*/
 	IsElide(userId,themeId string,atype int) bool
 	/**时间戳转换*/
	EncodeTime(s string) (str string)
 }
