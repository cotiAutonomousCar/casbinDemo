package ACS

import (
	"casbinDemo/utils/DB"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

func init() {
	// mysql 适配器
	adapter, _ := gormadapter.NewAdapterByDB(DB.Mysql)
	// 通过mysql适配器新建一个enforcer
	Enforcer, _ = casbin.NewEnforcer("config/keymatch2_model.conf", adapter)
	// 日志记录
	Enforcer.EnableLog(true)
}
