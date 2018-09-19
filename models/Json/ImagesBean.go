package Json

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：images数据类型
 ***********************************************************
 */
type ImagesBean struct {
	ImgUrl 			string `json:"img_url"`
	Width			int64  `json:"width"`
	Height			int64  `json:"height"`
	Hash			string `json:"hash"`
	ImgUrlSmall 	string `json:"small_url_small"`
	SmallWidth  	int64  `json:"small_width"`
	SmallHeight 	int64  `json:"small_height"`
}

