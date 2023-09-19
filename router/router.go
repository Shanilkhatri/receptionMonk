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
	// case "", "index", "login":
	// 	//utility.CheckACL(w, r, 0)
	// 	controllers.Login(w, r)
	// case "forgotpassword":
	// 	//utility.CheckACL(w, r, 0)
	// 	controllers.ForgotPassword(w, r)
	// case "changepassword":
	// 	//utility.CheckACL(w, r, 0)
	// 	controllers.ChangePassword(w, r)
	// case "dashboard":
	// 	// utility.CheckACL(w, r, 1)
	// 	controllers.Dashboard(w, r)
	// case "ajaxData":
	// 	// utility.CheckACL(w, r, 1)
	// 	controllers.AjaxData(w, r)
	case "products":
		// utility.CheckACL(w, r, 1)
		controllers.PutProducts(w, r)
	}

}
