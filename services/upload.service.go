package services

import "xiaomatv.cn/models/Json"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-12
 **---------------------------------------------------------
 **描述：上传服务类
 ***********************************************************
 */

type UploadService interface {
	/**获取文章详情*/
	GetUploadTokenService(atype int) (bean Json.TokenBean)
}
