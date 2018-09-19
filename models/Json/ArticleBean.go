package Json

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：ArticleBean数据类型
 ***********************************************************
 */
type ArticleBean struct {
	Id 		string `json:"id"`
	Title 	string `json:"title"`
	Type	string `json:"type"`
	Width 	int64  `json:"width"`
	Height 	int64  `json:"height"`
	Images 	[]*ImagesBean `json:"images"`
}