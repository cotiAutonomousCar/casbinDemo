package middleware

import (
	"casbinDemo/utils/ACS"
	"casbinDemo/utils/APIResponse"
	"casbinDemo/utils/Cache"
	"context"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Privilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		APIResponse.C = c
		var userName = c.GetHeader("userName")
		if userName == "" {
			APIResponse.Error("header miss userName")
			c.Abort()
			return
		}
		var domain = c.GetHeader("domain")
		if userName == "" {
			APIResponse.Error("header miss domain")
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method
		cacheName := userName + domain + path + method
		// 从缓存中读取&判断
		entry := Cache.RedisClient.Get(context.Background(), cacheName)
		if entry == nil && entry != nil {
			if string(entry.Val()) == "true" {
				c.Next()
			} else {
				APIResponse.Error("access denied")
				c.Abort()
				return
			}
		} else {
			// 从数据库中读取&判断
			//记录日志
			ACS.Enforcer.EnableLog(true)
			ACS.Enforcer.AddNamedDomainMatchingFunc("p", "KeyMatch2", util.KeyMatch2)
			ACS.Enforcer.BuildRoleLinks()
			// 加载策略规则
			err := ACS.Enforcer.LoadPolicy()
			if err != nil {
				log.Println("loadPolicy error")
				panic(err)
			}
			// 验证策略规则
			result, err := ACS.Enforcer.Enforce(userName, domain, path, method)
			if err != nil {
				APIResponse.Error("No permission found")
				c.Abort()
				return
			}
			if !result {
				// 添加到缓存中
				Cache.RedisClient.Set(context.Background(), cacheName, []byte("false"), 3600*time.Second)
				APIResponse.Error("access denied")
				c.Abort()
				return
			} else {
				Cache.RedisClient.Set(context.Background(), cacheName, []byte("true"), 3600*time.Second)
			}
			c.Next()
		}
	}
}
