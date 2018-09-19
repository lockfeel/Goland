package dao_impl

import (
		"xiaomatv.cn/models/Json"
)

/**图片列表*/
var (
	Images []*Json.ImagesBean
)

/**段落列表*/
var (
	Pends []*Json.ParagraphBean
)

/**实现类*/
type ArticleDaoImpl struct {

}

/**获取文章详情*/
func (impl ArticleDaoImpl) FindDesc(articleId,userId int) (bean Json.ArticleBean) {
	Images = make([]*Json.ImagesBean,2)
	Images[0] = &Json.ImagesBean{"x",100,100,"x","x",100,100}
	Images[1] = &Json.ImagesBean{"y",100,100,"u","y",100,100}
	return Json.ArticleBean{"id","X","1",100,300,Images}
}
/**获取文章段落*/
func (impl ArticleDaoImpl) FindPend(articleId,userId,pageNum,pageSize int) (pends []*Json.ParagraphBean) {
	Pends = make([]*Json.ParagraphBean,3)
	Pends[0] = &Json.ParagraphBean{"1","x","x","x","x"}
	Pends[1] = &Json.ParagraphBean{"2","y","y","y","y"}
	Pends[3] = &Json.ParagraphBean{"3","z","z","z","z"}
	return Pends;
}
/**检测文章存在*/
func (impl ArticleDaoImpl) CheckArticle(title string) (check bool) {
	return false
}
/**检测视频存在*/
func (impl ArticleDaoImpl) CheckVideo(title string) (check bool) {
	return true;
}