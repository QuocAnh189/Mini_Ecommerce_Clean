package casbin

import (
	"ecommerce_clean/db"

	"github.com/casbin/casbin/v2"
	casbinadapter "github.com/casbin/gorm-adapter/v3"
)

func InitCasbinEnforcer(db db.IDatabase) (*casbin.Enforcer, error) {
	adapter, err := casbinadapter.NewAdapterByDB(db.GetDB())
	if err != nil {
		return nil, err
	}

	// current working dir
	//for local development
	// wd, err := os.Getwd()
	// if err != nil {
	// 	return nil, err
	// }

	// modelPath := filepath.Join(wd, "policy/rbac_model.conf")

	//for docker container
	// modelPath := filepath.Join("/app", "policy/rbac_model.conf")
	modelPath := "/policy/rbac_model.conf"

	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, err
	}

	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	SetupPolicy(enforcer)

	return enforcer, nil
}
