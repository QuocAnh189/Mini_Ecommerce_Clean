package casbin

import "github.com/casbin/casbin/v2"

func SetupPolicy(enforcer *casbin.Enforcer) error {
	enforcer.AddPolicy("admin", "users", "read")
	enforcer.AddPolicy("admin", "users", "write")
	enforcer.AddPolicy("admin", "users", "delete")

	enforcer.AddPolicy("admin", "products", "read")
	enforcer.AddPolicy("admin", "products", "write")
	enforcer.AddPolicy("admin", "products", "delete")
	enforcer.AddPolicy("customer", "products", "read")

	return nil
}
