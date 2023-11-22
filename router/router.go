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
		if r.Method == "POST" {
			if controllers.CheckACL(w, r, []string{"owner", "user", "super-admin"}) {
				controllers.PostUser(w, r)
			}
		}
		if r.Method == "PUT" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin", "guest"}) {
				controllers.PutUser(w, r)
			}
		}
	case "calllogs":
		if r.Method == "PUT" {
			if controllers.CheckACL(w, r, []string{"owner"}) {
				controllers.PutCallLogs(w, r)
			}
		}

		if r.Method == "GET" {
			if controllers.CheckACL(w, r, []string{"owner"}) {
				controllers.GetCallLogsDetails(w, r)
			}
		}

	case "response":
		if r.Method == "GET" {
			if controllers.CheckACL(w, r, []string{"owner", "user"}) {
				controllers.GetResponse(w, r)
			}
		}
	//call this endpoint when user enter email.
	case "loginbyemail":
		if r.Method == "POST" {
			controllers.LoginByEmail(w, r)
		}
		//call this endpoint when user enter otp and submit.
	case "matchotp":
		if r.Method == "POST" {
			controllers.MatchOtp(w, r)
		}
		// controllers.Add(w, r)
	case "tokencheckfrontend":
		if r.Method == "GET" {
			controllers.CheckACLFrontend(w, r)
		}
	case "company":
		if r.Method == "POST" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin"}) {
				controllers.PostCompany(w, r)
			}
		}
		if r.Method == "PUT" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin"}) {
				controllers.PutCompany(w, r)
			}
		}
		if r.Method == "GET" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin"}) {
				controllers.GetCompany(w, r)
			}
		}
	case "kyc":
		if r.Method == "POST" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin", "user"}) {
				controllers.PostKycDetails(w, r)
			}
		}
		if r.Method == "PUT" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin", "user"}) {
				controllers.PutKycDetails(w, r)
			}
		}
		if r.Method == "GET" {
			if controllers.CheckACL(w, r, []string{"owner", "super-admin", "user"}) {
				controllers.GetKycDetails(w, r)
			}
		}
	}

}
