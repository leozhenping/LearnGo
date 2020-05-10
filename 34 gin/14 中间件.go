package main


/*
	c.Next() : 执行下一个中间件或功能函数，当下一个中件间或功能函数执行完成后，会返回c.next()的执行流程。并继续执行
	c.Abort() : 拒绝执行下一个函数，返回调用方

	c.use(middleware_func, ... ) : 加载全局中间件。 执行顺序从左到右执行.

	c.set(key, value) : 在中间件传递参数时使用的方法。在中间件函数执行的下一个函数，可以获取到指定的key
	c.get(key) :  获取中间件传递的参数，返回value及status两个值
	c.MustGet(key) : 必须获取到指定key的value， 否则panic


	// 全局添加中间件函数
	r.USE(middleware_func1,middleware_func2,middleware_func3)

	// 单独的URL添加中间件函数
	r.GET("/status", middleware_func1,middleware_func, statusHandler)

	// 给每个路由组添加添加中间件函数
	authV1 := r.group("/api/v1", middleware_func1,middleware_func2...)
	或者
	authV1.USE(middleware_func1, middleware_func2,...)

	中间件注意事项
	1. gin默认中间件
		gin.Default()默认使用了Logger和Recovery中间件，其中：
			Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
			Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
		如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

	2. gin中间件中使用goroutine
		当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。


*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthHandler() gin.HandlerFunc { // 返回类型必须是gin.HandlerFunc类型。 只要函数接受了*gin.context参数，就为gin.HandlerFunc类型
	return func(c *gin.Context) {
		c.Set("username","zhenping")
		fmt.Println("in AuthHandler")
		c.Next()
		//c.Abort()
		fmt.Println("out authHandler")

	}
}

func indexHandler(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		username = nil
	}
	fmt.Println("index func")
	c.JSON(http.StatusOK, gin.H{
		"msg": username,
	})
}

func statusHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"msg" : "status",
	})

}

func m1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In m1")
		start := time.Now()
		c.Next()
		costTime := time.Since(start)
		fmt.Println(costTime)
		fmt.Println("Out m1")
	}

}

func main() {

	r := gin.Default()

	// 1. 全局添加中间件函数
	r.Use(AuthHandler())

	r.GET("/index", indexHandler)

	// 2. 单独的URL添加中间件函数
	r.GET("/status", m1(), statusHandler)

	// 3. 为路径组添加中间件函数

	authV1 := r.Group("/api/v1", m1())
	{
		authV1.GET("/hostname", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg" : "auth v1 hostname",
			})
		})
	}

	// 或者使用以下方法为路由组添加中间件函数

	{
		authV1.Use(m1())
	}

	r.Run()
}
