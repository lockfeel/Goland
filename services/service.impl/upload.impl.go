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
 **描述：UploadServiceimpl层
 ***********************************************************
 */

type UploadServiceimpl struct {
	Dao  daos.UploadDao `inject:""`
}

/**获取文章详情*/
func (impl UploadServiceimpl) GetUploadTokenService(atype int) (bean Json.TokenBean) {
	return impl.Dao.GetUploadToken(atype)
}