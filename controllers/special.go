package controllers

import (
	"github.com/astaxie/beego"
	"xiaomatv.cn/services"
)

type Special struct {
	/*继承beego*/
	beego.Controller
	Serv services.ArticleService `inject:""`
}