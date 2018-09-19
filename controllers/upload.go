package controllers

import (
"github.com/astaxie/beego"
"xiaomatv.cn/services"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-12
 **---------------------------------------------------------
 **描述：资源上传控制器
 ***********************************************************
 */

type Upload struct {
	/*继承beego*/
	beego.Controller
	Serv services.UploadService `inject:""`
}

/**
 ************************************************************
 **方法::Upload::GetVideoToken
 **----------------------------------------------------------
 **描述::获取视频上传TOKEN
 **----------------------------------------------------------
 **参数:
 **param1:in-- 无
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018-08-12  Add by zwx
 ************************************************************
 */
func (a Upload) GetVideoToken() {
	a.Data["json"] = a.Serv.GetUploadTokenService(1)
	a.ServeJSON()
}
/**
 ************************************************************
 **方法::Upload::GetImageToken
 **----------------------------------------------------------
 **描述::获取主题|图片上传TOKEN
 **----------------------------------------------------------
 **参数:
 **param1:in-- 无
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018-08-12  Add by zwx
 ************************************************************
 */
func (a Upload) GetImageToken() {
	a.Data["json"] = a.Serv.GetUploadTokenService(2)
	a.ServeJSON()
}
/**
 ************************************************************
 **方法::Upload::GetCoverToken
 **----------------------------------------------------------
 **描述::获取用户头像上传TOKEN
 **----------------------------------------------------------
 **参数:
 **param1:in-- 无
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018-08-12  Add by zwx
 ************************************************************
 */
func (a Upload) GetCoverToken() {
	a.Data["json"] = a.Serv.GetUploadTokenService(3)
	a.ServeJSON()
}

/**
 ************************************************************
 **方法::Upload::GetCoverToken
 **----------------------------------------------------------
 **描述::视频转码成功回调
 **----------------------------------------------------------
 **参数:
 **param1:in-- 无
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018-08-12  Add by zwx
 ************************************************************
 */
func (a Upload) TurnsVideoCallback(){
	a.Serv.GetUploadTokenService(3)
}

/**
 ************************************************************
 **方法::Upload::TurnsImageCallbacK
 **----------------------------------------------------------
 **描述::图片转码成功回调
 **----------------------------------------------------------
 **参数:
 **param1:in-- 无
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018-08-12  Add by zwx
 ************************************************************
 */
func (a Upload) TurnsImageCallbacK(){
	a.Serv.GetUploadTokenService(3)
}