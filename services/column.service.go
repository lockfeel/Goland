package services

import "xiaomatv.cn/models/Mongo"

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-16
 **---------------------------------------------------------
 **描述：专栏服务类
 ***********************************************************
 */

 type ColumnService interface {
 	/**获取列表*/
 	GetListService(userId,pageNum,pageSize int) (b []*Mongo.ColumnBean)
 	/**获取用户专栏*/
 	GetUserService(userId int) (b []*Mongo.ColumnBean)
 	/**获取专栏详情*/
 	GetDescService(columnId,userId int) (b *Mongo.ColumnBean)
 }