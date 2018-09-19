package service_impl

import (
	"xiaomatv.cn/daos"
	"xiaomatv.cn/models/Json"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-01
 **---------------------------------------------------------
 **描述：ArticleServiceImpl层
 ***********************************************************
 */

type ArticleServiceImpl struct {
	Dao  daos.ArticleDao `inject:""`
}

/**获取文章详情*/
func (impl ArticleServiceImpl) FindDescService(articleId,userId int) (bean Json.ArticleBean) {
	return impl.Dao.FindDesc(articleId,userId)
}
/**获取文章段落*/
func (impl ArticleServiceImpl) FindPendService(articleId,userId,pageNum,pageSize int) (pends []*Json.ParagraphBean) {
	return impl.Dao.FindPend(articleId,userId,pageNum,pageSize)
}
/**检测文章存在*/
func (impl ArticleServiceImpl) CheckArticleService(title string) (check bool) {
	return impl.Dao.CheckArticle(title)
}
/**检测视频存在*/
func (impl ArticleServiceImpl) CheckVideoService(title string) (check bool) {
	return impl.Dao.CheckVideo(title)
}