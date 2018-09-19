package Mongo

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
	ImgUrl 			string `bson:"img_url"`
	Width			int32  `bson:"width"`
	Height			int32  `bson:"height"`
	Hash			string `bson:"hash"`
	ImgUrlSmall 	string `bson:"small_url_small"`
	SmallWidth  	int32  `bson:"small_width"`
	SmallHeight 	int32  `bson:"small_height"`
}
