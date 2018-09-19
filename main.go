package main
import (
	"github.com/astaxie/beego"
	_"xiaomatv.cn/routers"
	"net/http"
	"xiaomatv.cn/tools"
)

/**404自定义*/
func Page404(rw http.ResponseWriter, r *http.Request){
	panic("404")
}
/**
200 　OK ，客户端请求成功
201 　已创建
202　 已接收
203　 非认证信息 （非权威性信息）
204　 无内容
205 　重置内容
206　 部分内容
300 　多路选择 (表示被请求的文档可以在多个地方找到，并将在返回的文档中列出来。如果服务器有首选设置，首选项将会被列于定位响应头信息中。 )
301　 永久转移 (状态是指所请求的文档在别的地方；文档新的URL会在定位响应头信息中给出。浏览器会自动连接到新的URL。 )
302　 暂时转移 (与301有些类似，只是定位头信息中所给的URL应被理解为临时交换地址而不是永久的。)
303　 参见其他信息
304 　未修改（Not Modified）
305　 使用代理
400　 错误请求（Bad Request） 指出客户端请求中的语法错误。
401 　未授权（未授权客户机访问数据。）
402 　需要付费 (表示计费系统已有效。)
403　 禁止（Forbidden） （即使有授权也不需要访问。 ）
404　 未找到（Not Found） （服务器找不到给定的资源；文档不存在。 ）
405　 方法不允许 （用来访问本页面的 HTTP 谓词不被允许。）
406　 不接受 （客户端浏览器不接受所请求页面的 MIME 类型。）
407　 需要代理认证 （客户机首先必须使用代理认证自身。）
408　 请求超时
409　 冲突
410 　失败
411 　需要长度
412　 条件失败 （前提条件失败）
413 　请求实体太大
414 　请求URI太长
415 　不支持媒体类型
416   所请求的范围无法满足。
417   执行失败。
500　 服务器内部错误 501　 未实现（Not Implemented） （页眉值指定了未实现的配置。）
502　 网关失败 （Web 服务器用作网关或代理服务器时收到了无效响应。）
503 　服务不可用。这个错误代码为 IIS 6.0 所专用。
504 　网关超时
505   HTTP版本不支持

*/
func main() {
	tools.Momom()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	}
	beego.ErrorHandler("404",Page404)
	beego.Run()
}
