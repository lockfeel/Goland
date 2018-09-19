package services

import "xiaomatv.cn/models/Json"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：文章服务类
 ***********************************************************
 */

type ArticleService interface {
	/**获取文章详情*/
	FindDescService(articleId,userId int) (bean Json.ArticleBean)
	/**获取文章段落*/
	FindPendService(articleId,userId,pageNum,pageSize int) (pends []*Json.ParagraphBean)
	/**检测文章存在*/
	CheckArticleService(title string) (check bool)
	/**检测视频存在*/
	CheckVideoService(title string) (check bool)
}
