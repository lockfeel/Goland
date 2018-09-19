package daos

import "xiaomatv.cn/models/Mongo"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-16
 **---------------------------------------------------------
 **描述：专栏DAO层
 ***********************************************************
 */

type ColumnDao interface {
	/**获取列表*/
	GetLists(userId,pageNum,pageSize int) []*Mongo.ColumnBean
	/**获取用户专栏*/
	GetUsers(userId int) []*Mongo.ColumnBean
	/**获取专栏详情*/
	GetDesc(columnId,userId int) *Mongo.ColumnBean
}