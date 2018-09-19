package tools

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"xiaomatv.cn/conf"
	"sync"
)

/**
 ***********************************************************
 **作者：zwx<lockfeel1215@gmail.com>
 **---------------------------------------------------------
 **时间：2018-08-14
 **---------------------------------------------------------
 **描述：MongoDB工具类
 ***********************************************************
 */
type MongoDB struct {}
/**载体*/
var  instance *MongoDB
/**线程*/
var  once sync.Once
/**单例*/
func Mongo() *MongoDB {
	once.Do(func(){
		instance = &MongoDB{}
	})
	return instance
}
/**查询*/
func (m MongoDB) SELECT(dbName,tbName string,where,filter bson.M,sort string,start,size int,result interface{}){
	sos,err := mgo.Dial(conf.MONGOURL)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if er := db.Find(where).Select(filter).Sort(sort).Skip(start).Limit(size).All(result);er !=nil {
		panic(er)
	}
}
/**单条*/
func (m MongoDB) FindOne(dbName,tbName string,where,filter bson.M,result interface{}){
	sos,err := mgo.Dial(conf.MONGOURL)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if er := db.Find(where).Select(filter).One(result);er !=nil {
		panic(er)
	}
}
/**增加*/
func (m MongoDB) INSERT(dbName,tbName string,data interface{}) bool{
	sos,err := mgo.Dial(conf.MONGOPRO)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if er := db.Insert(&data);er != nil {
		return  false
	}else{
		return  true
	}
}
/**修改*/
func (m MongoDB) UPDATE(dbName,tbName string,where,reset bson.M) bool {
	sos,err := mgo.Dial(conf.MONGOPRO)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if er := db.Update(where,reset);er != nil {
		return  false
	}else{
		return  true
	}
}
/**批量*/
func (m MongoDB) UPDATEX(dbName,tbName string,where,reset bson.M) (affected int){
	sos,err := mgo.Dial(conf.MONGOPRO)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if x,er := db.UpdateAll(where,reset);er != nil {
		return  0
	}else{
		return  x.Updated
	}
}
/**删除*/
func (m MongoDB) DELETE(dbName,tbName string,where bson.M) (affected int){
	sos,err := mgo.Dial(conf.MONGOPRO)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if er := db.Remove(where);er != nil {
		return  0
	}else{
		return  1
	}
}
/**清除*/
func (m MongoDB) REMOVE(dbName,tbName string,where bson.M) (affected int){
	sos,err := mgo.Dial(conf.MONGOPRO)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if x,er := db.RemoveAll(where);er != nil {
		return  0
	}else{
		return  x.Removed
	}
}
/**求和*/
func (m MongoDB) COUNT(dbName,tbName string,where bson.M) (num int){
	sos,err := mgo.Dial(conf.MONGOURL)
	if err !=nil {
		panic(err)
	}
	db := sos.DB(dbName).C(tbName)
	if x,er := db.Find(where).Count();er != nil {
		return  0
	}else{
		return  x
	}
}

