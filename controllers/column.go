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
 **时间：2018-08-15
 **---------------------------------------------------------
 **描述：专栏类型
 ***********************************************************
 */

 type Column struct {
	/*继承beego*/
	beego.Controller
	Serv services.ColumnService `inject:""`
 }

/**
 ************************************************************
 **方法::Column::GetLists
 **----------------------------------------------------------
 **描述::获取专栏列表
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id    :: 用户ID
 **param2:in-- Int : page_num   :: 分页数量
 **param3:in-- Int : page_size  :: 分页大小
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Column) GetLists() {
	 userId,err1    := strconv.Atoi(a.GetString("user_id"))
	 pageNum,err2   := strconv.Atoi(a.GetString("page_num"))
	 pageSize,err3  := strconv.Atoi(a.GetString("page_size"))
	 result := make(map[string]interface{})
	 if err1 != nil || err2 != nil || err3 != nil{
		 result["retCode"] = 1
		 result["message"] = "参数缺失"
		 result["data"]	   = [0]*Mongo.ColumnBean{}
	 }else{
		 result["retCode"] = 0
		 result["message"] = "SUCCESS"
		 result["data"]	   = a.Serv.GetListService(userId,pageNum,pageSize)
	 }
	 a.Data["json"] = result
	 a.ServeJSON()
 }

/**
 ************************************************************
 **方法::Column::GetUsers
 **----------------------------------------------------------
 **描述::获取用户专栏
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : user_id  :: 用户ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data  :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Column) GetUsers(){
	 userId,err1    := strconv.Atoi(a.GetString("user_id"))
	 result := make(map[string]interface{})
	 if err1 != nil {
		 result["retCode"] = 1
		 result["message"] = "参数缺失"
		 result["data"]	   = [0]*Mongo.ColumnBean{}
	 }else{
		 result["retCode"] = 0
		 result["message"] = "SUCCESS"
		 result["data"]	   = a.Serv.GetUserService(userId)
	 }
	 a.Data["json"] = result
	 a.ServeJSON()
 }
/**
 ************************************************************
 **方法::Column::GetDesc
 **----------------------------------------------------------
 **描述::获取专栏详情
 **----------------------------------------------------------
 **参数:
 **param1:in-- Int : column_id  :: 专栏ID
 **param2:in-- Int : user_id    :: 用户ID
 **----------------------------------------------------------
 **返回：
 **return:out--  JSON ：data    :: JSON对象
 **----------------------------------------------------------
 **日期:2018.08.15  Add by zwx
 ************************************************************
 */
 func (a *Column) GetDesc(){
	 userId,err1     := strconv.Atoi(a.GetString("user_id"))
	 columnId,err2   := strconv.Atoi(a.GetString("column_id"))
	 result := make(map[string]interface{})
	 if err1 != nil || err2 != nil{
		 result["retCode"] = 1
		 result["message"] = "参数缺失"
		 result["columnDesc"] = [0]*Mongo.ColumnBean{}
	 }else{
		 result["retCode"] = 0
		 result["message"] = "SUCCESS"
		 result["columnDesc"] = a.Serv.GetDescService(columnId,userId)
	 }
	 a.Data["json"] = result
	 a.ServeJSON()

 }
