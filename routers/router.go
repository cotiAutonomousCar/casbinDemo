package routers

import (
	"casbinDemo/middleware"
	"casbinDemo/utils/ACS"
	"casbinDemo/utils/APIResponse"
	"casbinDemo/utils/Cache"
	"context"
	"github.com/gin-gonic/gin"
)

var (
	R *gin.Engine
)

func init() {
	R = gin.Default()
	R.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "message": "Bad Request"})
	})
	api()
}
func api() {
	auth := R.Group("/api")
	{
		// 模拟添加一条Policy策略
		auth.POST("acs", func(c *gin.Context) {
			APIResponse.C = c
			subject := "tom"
			domain := "supTech"
			object := "/api/routers"
			action := "POST"
			cacheName := subject + domain + object + action
			result, _ := ACS.Enforcer.AddPolicy(subject, domain, object, action)
			if result {
				// 清除缓存
				_ = Cache.RedisClient.Expire(context.Background(), cacheName, 0)
				APIResponse.Success("add success")
			} else {
				APIResponse.Error("add fail")
			}
		})
		// 模拟删除一条Policy策略
		auth.DELETE("acs/:id", func(context *gin.Context) {
			APIResponse.C = context
			result, _ := ACS.Enforcer.RemovePolicy("tom", "supTech", "/api/routers", "POST")
			if result {
				// 清除缓存 代码省略
				APIResponse.Success("delete Policy success")
			} else {
				APIResponse.Error("delete Policy fail")
			}
		})
		// 获取路由列表
		auth.POST("/routers", middleware.Privilege(), func(c *gin.Context) {
			type data struct {
				Method string `json:"method"`
				Path   string `json:"path"`
			}
			var datas []data
			routers := R.Routes()
			for _, v := range routers {
				var temp data
				temp.Method = v.Method
				temp.Path = v.Path
				datas = append(datas, temp)
			}
			APIResponse.C = c
			APIResponse.Success(datas)
			return
		})
	}
	// 定义路由组
	user := R.Group("/api/v1")
	// 使用访问控制中间件
	user.Use(middleware.Privilege())
	{
		user.POST("user", func(c *gin.Context) {
			c.JSON(200, gin.H{"code": 200, "message": "user add success"})
		})
		user.DELETE("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
		})
		user.PUT("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
		})
		user.GET("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
		})
	}
}
