package helper

import (
	"blog-go-gin/dao"
	"blog-go-gin/logging"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var CasbinEnforcer *casbin.SyncedEnforcer

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup() {
	Apter, err := gormAdapter.NewAdapterByDBUseTableName(dao.Db, "tb", "casbin")
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
	CasbinEnforcer = e
}

func Casbin() *casbin.SyncedEnforcer {
	return CasbinEnforcer
}

func LoadPolicy() (*casbin.SyncedEnforcer, error) {
	if err := CasbinEnforcer.LoadPolicy(); err != nil {
		logging.Logger.Error("casbin rbac_model or policy init error, message: %v \r\n", err.Error())
		return nil, err
	}
	return CasbinEnforcer, nil
}
