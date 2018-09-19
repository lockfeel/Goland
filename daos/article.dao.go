package daos

import "xiaomatv.cn/models/Json"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2017-09-09
 **---------------------------------------------------------
 **描述：ArticleDao层
 ***********************************************************
 */
type ArticleDao interface{
	/**获取文章详情*/
	FindDesc(articleId,userId int) (bean Json.ArticleBean)
	/**获取文章段落*/
	FindPend(articleId,userId,pageNum,pageSize int) (pends []*Json.ParagraphBean)
	/**检测文章存在*/
	CheckArticle(title string) (check bool)
	/**检测视频存在*/
	CheckVideo(title string) (check bool)
}