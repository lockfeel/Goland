package dao_impl

import (
	"xiaomatv.cn/models/Mongo"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"xiaomatv.cn/tools"
	)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-16
 **---------------------------------------------------------
 **描述：DAO层实现类
 ***********************************************************
 */
var(
	dbName = "xiaomatv"
	tbName = "t_column"
)
type ColumnDaoImpl struct {}
/**获取列表*/
func (a ColumnDaoImpl) GetLists(userId,pageNum,pageSize int) []*Mongo.ColumnBean{
	var result []*Mongo.ColumnBean
	tools.Mongo().SELECT(dbName,tbName,bson.M{},bson.M{"videoList":0},"id",(pageNum-1)*pageSize,pageSize,&result)
	if len(result) <=0 {
		return result
	} else {
		return result
	}
}
/**获取用户专栏*/
func (a ColumnDaoImpl) GetUsers(userId int) []*Mongo.ColumnBean{
	var result []*Mongo.ColumnBean
	tools.Mongo().SELECT(dbName,tbName,bson.M{"user_id":strconv.Itoa(userId)},bson.M{"videoList":0},"id",0,1000,&result)
	if len(result) <=0 {
		return result
	} else {
		return result
	}
}
/**获取专栏详情*/
func (a ColumnDaoImpl)  GetDesc(columnId,userId int) *Mongo.ColumnBean{
	var result Mongo.ColumnBean
	tools.Mongo().FindOne(dbName,tbName,bson.M{"column_id":strconv.Itoa(columnId)},bson.M{"_id":0},&result)
	return &result
}