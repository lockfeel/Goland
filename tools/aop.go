package tools

import (
	"fmt"
	"github.com/gogap/aop"
)

type Auth struct {
}

func (p *Auth) Login(userName, password string) bool {
	fmt.Println("@Login", userName, password)
	if userName == "zeal" && password == "gogap" {
		return true
	}
	return false
}

type Moli struct {}

// @AfterReturning, the method could have args of aop.Result,
// it will get the result from real func return values
func (x *Moli) AfterReturning(result aop.Result){
	fmt.Println("Return value:", result)
}

func (x *Moli) Before(jp aop.JoinPointer){
	username := ""
	password := ""
	jp.Args().MapTo(func(u, p string) {
		username = u
		password = p
	})
	fmt.Printf("Before Login: %s %s\n", username,password)
}
// the args is same as Login
func (x *Moli) After(username, password string) bool {
	fmt.Printf("After Login: %s %s\n", username, password)
	return true
}

// use join point to around the real func of login
func (x *Moli) Around(pjp aop.ProceedingJoinPointer) {
	fmt.Println("@Begin Around")
	ret := pjp.Proceed("zeal","gogap")
	ret.MapTo(func(loginResult bool) {
		fmt.Println("@Proceed Result is", loginResult)
	})
	fmt.Println("@End Around")
}
func Momom() {

	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("auth", new(Auth))
	beanFactory.RegisterBean("foo", new(Moli))

	aspect := aop.NewAspect("aspect_1", "auth")
	aspect.SetBeanFactory(beanFactory)

	aspectFoo := aop.NewAspect("aspect_2", "foo")
	aspectFoo.SetBeanFactory(beanFactory)

	pointcut := aop.NewPointcut("pointcut_1").Execution(`Login()`)

	aspect.AddPointcut(pointcut)
	aspectFoo.AddPointcut(pointcut)

	aspectFoo.AddAdvice(&aop.Advice{Ordering: aop.Before, 		  	Method: "Before",  			  PointcutRefID: "pointcut_1"})
	aspectFoo.AddAdvice(&aop.Advice{Ordering: aop.After, 			Method: "After",  			  PointcutRefID: "pointcut_1"})
	aspectFoo.AddAdvice(&aop.Advice{Ordering: aop.Around, 			Method: "Around", 			  PointcutRefID: "pointcut_1"})
	aspectFoo.AddAdvice(&aop.Advice{Ordering: aop.AfterReturning, 	Method: "AfterReturning", 	  PointcutRefID: "pointcut_1"})

	gogapAop := aop.NewAOP()
	gogapAop.SetBeanFactory(beanFactory)
	gogapAop.AddAspect(aspect)
	gogapAop.AddAspect(aspectFoo)

	var err error
	var proxy *aop.Proxy

	// Get proxy
	if proxy, err = gogapAop.GetProxy("auth"); err != nil {
		fmt.Println("get proxy failed", err)
		return
	}
	// start trace for debug
	aop.StartTrace()
	fmt.Println("==========Func Type Assertion==========")
	login := proxy.Method(new(Auth).Login).(func(string, string) bool)("zeal", "gogap")
	fmt.Println("Login result:", login)
	t, _ := aop.StopTrace()
	// print trace result
	for _, item := range t.Items() {
		fmt.Println(item.ID, item.InvokeID, item.BeanRefID, item.Pointcut, item.Method)
	}
}