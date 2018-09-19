package controllers

import (
	"github.com/astaxie/beego"
	"xiaomatv.cn/services"
	"strconv"
	"xiaomatv.cn/models/Mongo"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018.08.13
 **---------------------------------------------------------
 **描述：内涵段子处理控制器
 ***********************************************************
 */

type Connotation struct {
	/*继承beego*/
	beego.Controller
	Serv services.ConnotationSerivce `inject:""`
}

/**
 ************************************************************
 **方法::Connotation::TopicShow
 **----------------------------------------------------------
 **描述::获取推荐展示
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : type 		:: 请求类型
 **param2:in-- Int : user_id    :: 用户ID
 **param3:in-- Int : page_num   :: 分页数量
 **param4:in-- Int : page_size  :: 分页大小
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.13  Add by zwx
 ************************************************************
 */
func (a *Connotation) TopicShow() {
	userId,err1    := strconv.Atoi(a.GetString("user_id"))
	pageNum,err2   := strconv.Atoi(a.GetString("page_num"))
	pageSize,err3  := strconv.Atoi(a.GetString("page_size"))
	atype,err4	   := strconv.Atoi(a.GetString("type"))
	result := make(map[string]interface{})
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		result["retCode"] = 1
		result["message"] = "参数缺失"
		result["data"]	  = [0]*Mongo.ThemeBean{}
		a.Data["json"] 	  = result
		a.ServeJSON()
		return
	}
	a.Data["json"] = a.Serv.TopicShowSerivce(userId,atype,pageNum,pageSize)
	a.ServeJSON()
}

/**
 ************************************************************
 **方法::Connotation::CommentList
 **----------------------------------------------------------
 **描述::评论展示
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : request_id :: 评论Id
 **param2:in-- Int : user_id    :: 用户ID
 **param3:in-- Int : page_num   :: 分页数量
 **param4:in-- Int : page_size  :: 分页大小
 **param5:in-- Int : type		:: 主题类型
 **param6:in-- Int : theme_id   :: 主题ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.13  Add by zwx
 ************************************************************
 */
func (a *Connotation) CommentList(){
	requestId,err1 := strconv.Atoi(a.GetString("request_id"))
	userId,err2    := strconv.Atoi(a.GetString("user_id"))
	pageNum,err3   := strconv.Atoi(a.GetString("page_num"))
	themeId,err4   := strconv.Atoi(a.GetString("theme_id"))
	pageSize,err5  := strconv.Atoi(a.GetString("page_size"))
	atype,err6	   := strconv.Atoi(a.GetString("type"))
	result := make(map[string]interface{})
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil{
		result["retCode"] = 1
		result["message"] = "参数缺失"
		result["data"]	  = [0]*Mongo.ThemeBean{}
		a.Data["json"] 	  = result
		a.ServeJSON()
		return
	}
	println(atype)
	a.Data["json"] = a.Serv.CommentListSerivce(themeId,requestId,userId,pageNum,pageSize)
	a.ServeJSON()

}

/**
 ************************************************************
 **方法::Connotation::TopicChoice
 **----------------------------------------------------------
 **描述::话题精选
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : title 		:: 话题标题
 **param2:in-- Int : user_id    :: 用户ID
 **param3:in-- Int : page_num   :: 分页数量
 **param4:in-- Int : page_size  :: 分页大小
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.13  Add by zwx
 ************************************************************
 */
func (a *Connotation) TopicChoice(){
	a.Data["json"] = a.Serv.TopicChoiceSerivce("x",1,1,1)
	a.ServeJSON()
}
