package router

import (
	"net/http"
	"reakgo/controllers"
	"strings"
)

func Routes(w http.ResponseWriter, r *http.Request) {
	// Trailing slash is a pain in the ass so we just drop it
	route := strings.Trim(r.URL.Path, "/")
	switch route {
	case "", "index":
		check := controllers.CheckACL(w, r, []string{"guest", "admin", "user"})
		if check {
			controllers.BaseIndex(w, r)
		}
	case "login":
		controllers.CheckACL(w, r, []string{"admin", "user"})
		controllers.Login(w, r)
	case "dashboard":
		check := controllers.CheckACL(w, r, []string{"admin", "user"})
		if check {
			controllers.Dashboard(w, r)
		}
	case "addSimpleForm":
		controllers.CheckACL(w, r, []string{"admin", "user"})
		controllers.AddForm(w, r)
	case "viewSimpleForm":
		controllers.CheckACL(w, r, []string{"admin", "user"})
		controllers.ViewForm(w, r)
	case "register-2fa":
		controllers.CheckACL(w, r, []string{"admin", "user"})
		controllers.RegisterTwoFa(w, r)
	case "verify-2fa":
		controllers.CheckACL(w, r, []string{"admin", "user"})
		controllers.VerifyTwoFa(w, r)
	case "orders":
		if controllers.CheckACL(w, r, []string{"owner", "user", "super-admin"}) {
			if r.Method == "GET" {
				controllers.GetOrders(w, r)
			}

		}
	case "users":
		if controllers.CheckACL(w, r, []string{"owner", "user", "super-admin"}) {

			if r.Method == "POST" {
				controllers.PostUser(w, r)
			}
		}
		if r.Method == "PUT" {

			controllers.PutUser(w, r)
		}
	case "calllogs":
		if controllers.CheckACL(w, r, []string{"owner", "user", "super-admin"}) {
			if r.Method == "PUT" {
				controllers.GetOrders(w, r)
			}
		}
	}

}
