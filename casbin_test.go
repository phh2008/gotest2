package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

var e *casbin.Enforcer

func init() {
	var err error
	e, err = casbin.NewEnforcer("D:\\codespace\\golang_test\\go-test2\\model.conf", "D:\\codespace\\golang_test\\go-test2\\policy.csv")
	if err != nil {
		panic(err)
	}
}

func Test1(t *testing.T) {
	m := e.GetModel()
	fmt.Println(m)
}

// 增加策略
func Test2(t *testing.T) {
	//if ok, _ := e.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
	if ok, _ := e.AddPolicy("admin", "/api/v1/hello/:id", "GET"); !ok {
		fmt.Println("Policy已经存在")
	} else {
		e.SavePolicy()
		fmt.Println("增加成功")
	}
}

// 删除策略
func Test3(t *testing.T) {
	fmt.Println("删除Policy")
	if ok, _ := e.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
		fmt.Println("Policy不存在")
	} else {
		fmt.Println("删除成功")
		e.SavePolicy()
	}
}

// 获取策略
func Test4(t *testing.T) {
	fmt.Println("查看policy")
	list := e.GetPolicy()
	for _, vlist := range list {
		for _, v := range vlist {
			fmt.Printf("value: %s, ", v)
		}
		fmt.Printf("\n")
	}
}

// 更新
func Test5(t *testing.T) {
	e.UpdatePolicy([]string{"alice", "data1", "read"}, []string{"alice", "data1", "write"})
	e.SavePolicy()
}

// 查询，条件过滤
func Test6(t *testing.T) {
	// 更新策略，根据条件(fieldIndex：表示参数开始字段索引)
	// 生成 sql如下，
	// DELETE FROM `casbin_rule` WHERE ptype = 'p' and v0 = 'bob' and v1 = 'data2'
	// INSERT INTO `casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`) VALUES ('p','admin','/api/v3/hello','POST','','','','','')
	e.UpdateFilteredPolicies([][]string{{"admin", "/api/v3/hello", "POST"}}, 0, "bob", "data2")
	e.SavePolicy()
}

// 添加用户角色
func Test7(t *testing.T) {
	ret, _ := e.AddGroupingPolicy("tom", "admin")
	e.SavePolicy()
	fmt.Println(ret)
}

// 验证是否具有权限
func Test8(t *testing.T) {
	//ret, err := e.Enforce("tom", "/api/v1/hello", "GET")
	ret, err := e.Enforce("tom", "/api/v1/hello/11", "GET")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}

// 查询用户角色
func Test9(t *testing.T) {
	ret, _ := e.GetRolesForUser("tom")
	fmt.Println(ret)
}

// 查询角色权限
func Test10(t *testing.T) {
	ret := e.GetPermissionsForUser("admin")
	fmt.Println(ret)
}

func Test11(t *testing.T) {
	//admin
	e.AddPermissionForUser("role:admin", "ad_campaign", "GET")
	e.AddPermissionForUser("role:admin", "ad_campaign", "LIST")
	//area-admin
	e.AddPermissionForUser("role:area_ad_admin", "campaign", "WRITE")
	e.AddPermissionForUser("role:area_ad_admin", "adgroup", "WRITE")
	e.AddPermissionForUser("role:area_ad_admin", "adcreative", "WRITE")
	//assigns role of area_ad_admin to admin.
	e.AddRoleForUser("role:admin", "role:area_ad_admin")

	//end users.
	e.AddRoleForUser("kevin", "role:admin")
	e.SavePolicy()

	//print.
	ret, _ := e.GetImplicitPermissionsForUser("kevin")
	fmt.Printf("%#v\n", ret)

	fmt.Println(e.Enforce("kevin", "adgroup", "WRITE"))
}
