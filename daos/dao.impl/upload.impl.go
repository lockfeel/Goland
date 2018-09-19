package dao_impl

import (
		"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"xiaomatv.cn/models/Json"
)

/**实现类*/
type UploadDaoImpl struct {

}
/**七牛云*/
var (
	accessKey    = "h6KUC7vP0w-Tl3tgNRh4r6cug0ZuavIQCqzVW6W2"
	secretKey    = "DBCciIrGd-00yMGUN74pJsEbu0915yRZ3zdSia5z"
	bucketPic    = "xiaomauserpic"
	bucketImg    = "xiaomapostpic"
	bucketVideo  = "xiaomaoriginal"
	bucketCover  = "xiaomaavatar"
)

/**获取文章详情*/
func (impl UploadDaoImpl) GetUploadToken(atype int) (bean Json.TokenBean) {
	mac := qbox.NewMac(accessKey,secretKey)
	var token = ""
	var cover = ""
	if atype == 1{
		putPolicy := storage.PutPolicy{
			Scope: bucketVideo,
		}
		putPolicy.Expires = 7200
		token = putPolicy.UploadToken(mac)
		putPerry := storage.PutPolicy{
			Scope: bucketCover,
		}
		putPerry.Expires = 7200
		cover = putPolicy.UploadToken(mac)
	}else if atype == 2 {
		putPolicy := storage.PutPolicy{
			Scope: bucketImg,
		}
		putPolicy.Expires = 7200
		token = putPolicy.UploadToken(mac)
	}else if atype == 3 {
		putPolicy := storage.PutPolicy{
			Scope: bucketPic,
		}
		putPolicy.Expires = 7200
		token = putPolicy.UploadToken(mac)
	}
	return Json.TokenBean{token,token,cover,1,accessKey}
}