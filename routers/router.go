package routers

import (
	"github.com/astaxie/beego"
	"xiaomatv.cn/controllers"
	"encoding/json"
	"xiaomatv.cn/conf"
	"github.com/facebookgo/inject"
	"xiaomatv.cn/daos/dao.impl"
	"xiaomatv.cn/services/service.impl"
	"os"
)
/**路由配置类*/
type RouterTree struct {
	Title 	 string  `json:"title"`
	Path  	 string  `json:"path"`
	Action   string  `json:"action"`
	Function string  `json:"function"`
}
/**初始化方法*/
func init() {
	/**注入对象*/
	var ioc inject.Graph
	/**控制器对象*/
	action  := make(map[string]beego.ControllerInterface)
	action["a"] = &controllers.Article{}
	action["b"] = &controllers.Upload{}
	action["c"] = &controllers.Connotation{}
	action["d"] = &controllers.Classify{}
	action["e"] = &controllers.Column{}
	action["f"] = &controllers.Interactive{}
	action["g"] = &controllers.Republic{}
	action["h"] = &controllers.Special{}
	action["i"] = &controllers.Activity{}
	for key,_ := range action {
		if err := ioc.Provide(&inject.Object{Value: action[key]});err != nil{
			panic(err)
		}
	}
	/**服务层对象*/
	server := make(map[string]interface{})
	server["as"] = &service_impl.ArticleServiceImpl{}
	server["cs"] = &service_impl.ConnotationServiceImpl{}
	server["es"] = &service_impl.ColumnServiceImpl{}
	server["bs"] = &service_impl.UploadServiceimpl{}
	for key,_ := range server {
		if err := ioc.Provide(&inject.Object{Value: server[key]});err != nil{
			panic(err)
		}
	}
	/**数据层对象*/
	dao := make(map[string]interface{})
	dao["_di"]=  &dao_impl.BaseDaoImpl{}
	dao["ad"] =  &dao_impl.ArticleDaoImpl{}
	dao["bd"] =  &dao_impl.UploadDaoImpl{}
	dao["ed"] =  &dao_impl.ColumnDaoImpl{}
	dao["cd"] =  &dao_impl.ConnotationDaoImpl{}
	for key,_ := range dao {
		if err := ioc.Provide(&inject.Object{Value: dao[key]});err != nil{
			panic(err)
		}
	}
	/**多对象拼装*/
	if err := ioc.Populate(); err != nil {
		panic(err)
		os.Exit(1)
	}
	/**拉取路由配置文件*/
	var routers []*RouterTree
	file,_ := os.Open(conf.ROUTERJS)
	code   := json.NewDecoder(file)
	if err := code.Decode(&routers); err != nil {
		panic(err)
	}
	for i,_ := range routers {
		beego.Router(routers[i].Path,action[routers[i].Action], routers[i].Function)
	}
	defer file.Close()
}

