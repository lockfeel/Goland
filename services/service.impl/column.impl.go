package service_impl

import (
	"xiaomatv.cn/models/Mongo"
	"xiaomatv.cn/daos"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-16
 **---------------------------------------------------------
 **描述：专栏服务实现类
 ***********************************************************
 */

 type ColumnServiceImpl struct {
	 Dao  daos.ColumnDao `inject:""`
 }
 /**获取列表*/
 func (a ColumnServiceImpl) GetListService(userId,pageNum,pageSize int) (b []*Mongo.ColumnBean){
	return a.Dao.GetLists(userId,pageNum,pageSize)
 }
 /**获取用户专栏*/
 func (a ColumnServiceImpl) GetUserService(userId int) (b []*Mongo.ColumnBean) {
	return a.Dao.GetUsers(userId)
 }
 /**获取专栏详情*/
 func (a ColumnServiceImpl) GetDescService(columnId,userId int) (b *Mongo.ColumnBean) {
	return a.Dao.GetDesc(columnId,userId)
 }