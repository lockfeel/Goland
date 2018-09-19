package Json

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-07-30
 **---------------------------------------------------------
 **描述：TokenBean数据类型
 ***********************************************************
 */
 type TokenBean struct {
 	Token 	  	string `json:"token"`
 	VideoToken  string `json:"video_token"`
 	PicToken    string `json:"pic_token"`
 	RetCode     int    `json:"retCode"`
 	Key		    string `json:"key"`
}
