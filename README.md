# 访问接口
# 参数缺失
curl -X POST http://127.0.0.1:8080/api/routers
{"code":-1,"message":"header miss userName","data":null}


# 无访问权限
curl -X POST -H "userName:tom,domain:supTech" http://127.0.0.1:8080/api/routers
{"code":-1,"message":"access denied","data":null}


# 添加一条规则(代码中是模拟数据)
curl -X POST http://127.0.0.1:8080/api/acs
{"code":200,"message":"success","data":"add success"}

# 再次访问(有访问权限,可以访问)
curl -X POST -H "userName:tom,domain:supTech" http://127.0.0.1:8080/api/routers

{
"code":200,
"message":"success",
"data":[
{
"method":"POST",
"path":"/api/acs"
},
{
"method":"POST",
"path":"/api/routers"
},
{
"method":"POST",
"path":"/api/v1/user"
},
{
"method":"DELETE",
"path":"/api/acs/:id"
},
{
"method":"DELETE",
"path":"/api/v1/user/:id"
},
{
"method":"PUT",
"path":"/api/v1/user/:id"
},
{
"method":"GET",
"path":"/api/v1/user/:id"
}
]
}

# 直接向数据库添加几条Policy策略
INSERT INTO `temp`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'bianjie', 'supTech', '/api/v1/user', 'POST', NULL, NULL);
INSERT INTO `temp`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'bianjie', 'supTech', '/api/v1/user/:id', 'GET', NULL, NULL);
INSERT INTO `temp`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'bianjie', 'ccv', '/api/v1/user/:id', 'PUT', NULL, NULL);
INSERT INTO `temp`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', 'admin', 'bianjie', 'supTech', NULL, NULL, NULL);
INSERT INTO `temp`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', 'techCentor', 'bianjie', 'supTech', NULL, NULL, NULL);
INSERT INTO `temp`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', 'appCentor', 'bianjie', 'supTech', NULL, NULL, NULL);

#再测试
## 添加接口
curl -X POST -H "userName:admin,domain:supTech" http://127.0.0.1:8080/api/v1/user
{"code":200,"message":"user add success"}
## 查询接口
curl -X GET -H "userName:admin,domain:supTech" http://127.0.0.1:8080/api/v1/user/99
{"code":200,"message":"user Get success 99"}
## 更新接口
curl -X PUT -H "userName:admin,domain:supTech" http://127.0.0.1:8080/api/v1/user/199
{"code":200,"message":"user update success 199"}
## 删除接口(没有分配访问权限)
curl -X DELETE -H "userName:admin,domain:supTech" http://127.0.0.1:8080/api/v1/user/299
{"code":-1,"message":"access denied","data":null}
