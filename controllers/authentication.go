package controllers

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"reakgo/models"
	"reakgo/utility"

	"github.com/gorilla/sessions"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Check for any form parsing error
		err := r.ParseForm()
		if err != nil {
			log.Println("form parsing failed")
			utility.RenderTemplate(w, r, "login", "demo")
		} else {
			// Parsing form went fine, Now we can access all the values
			email := r.FormValue("email")
			password := r.FormValue("password")
			confirmPassword := r.FormValue("confirmPassword")
			//rememberMe := r.FormValue("rememberMe")

			// No need to check for empty values as DB Authentication will take care of it

			// Backend validation for password and confirmPassword
			if confirmPassword == password {
				row, err := models.Authentication{}.GetUserByEmail(email)
				if err != nil {
					// In case of MYSQL issues or no results are returned
					log.Println(err)
					utility.AddFlash("error", "Credentials didn't match, Please try again.", w, r)
					utility.RenderTemplate(w, r, "login", "demo")
				} else {
					match := bcrypt.CompareHashAndPassword([]byte(row.PasswordHash), []byte(r.FormValue("password")))
					if match != nil {
						utility.AddFlash("error", "Credentials didn't match, Please try again.", w, r)
						utility.RenderTemplate(w, r, "login", "demo")
					} else {
						// Password match has been a success
						utility.SessionSet(w, r, utility.Session{Key: "id", Value: row.ID})
						utility.SessionSet(w, r, utility.Session{Key: "email", Value: row.Email})
						utility.SessionSet(w, r, utility.Session{Key: "type", Value: row.AccountType})
						utility.AddFlash("success", "Success !, Logged in.", w, r)
						if r.FormValue("rememberMe") == "yes" {
							utility.Store.Options = &sessions.Options{
								MaxAge: 60 * 1,
							}
						}
						token := models.Authentication{}.CheckTwoFactorRegistration(int32(row.ID))
						utility.SessionSet(w, r, utility.Session{Key: "2faSecret", Value: token})

						if token != "" {
							utility.RedirectTo(w, r, "verify-2fa")
						} else {
							utility.RedirectTo(w, r, "dashboard")
						}
						//utility.RenderTemplate(w, r, "login", "demo")
					}
				}
			}
		}
	} else {
		utility.RenderTemplate(w, r, "login", "demo")
	}
}

func VerifyTwoFa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utility.RenderTemplate(w, r, "verifyTwoFa", "demo")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			twoFaVerify := r.FormValue("twoFaVerify")
			secret := fmt.Sprintf("%v", utility.SessionGet(r, "2faSecret"))
			if totp.Validate(twoFaVerify, secret) {
				utility.RedirectTo(w, r, "dashboard")
			} else {
				utility.RedirectTo(w, r, "login")
			}
		}
	}
}

func RegisterTwoFa(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		email := utility.SessionGet(r, "email")
		key, _ := totp.Generate(totp.GenerateOpts{
			Issuer:      os.Getenv("TOTP_ISSUER"),
			AccountName: fmt.Sprintf("%v", email),
		})
		utility.SessionSet(w, r, utility.Session{Key: "totpSecret", Value: key.Secret()})

		img, _ := key.Image(400, 400)

		var buf bytes.Buffer
		png.Encode(&buf, img)

		data := b64.StdEncoding.EncodeToString([]byte(buf.String()))
		data = "data:image/png;base64," + data

		utility.RenderTemplate(w, r, "twoFactorRegister", data)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {

		} else {
			verifyToken := r.FormValue("challengeCode")
			validationResult := totp.Validate(verifyToken, fmt.Sprintf("%v", utility.SessionGet(r, "totpSecret")))
			if validationResult {
				secret := fmt.Sprintf("%v", utility.SessionGet(r, "totpSecret"))
				userId := fmt.Sprintf("%v", utility.SessionGet(r, "id"))
				intUserId, _ := strconv.Atoi(userId)
				models.Authentication{}.TwoFactorAuthAdd(secret, intUserId)
				utility.RenderTemplate(w, r, "successTwoFactor", nil)
			} else {
				// Show Error Page
				utility.RenderTemplate(w, r, "failureTwoFactor", nil)
			}
		}
	}
}

type SignupDetails struct {
	Email    string `json:"authEmailId"`
	Password string `json:"password"`
	Otp      string `json:"authSignInOTP"`
}

// Generate Otp .
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10) // Generate a random digit (0-9).
		otp += fmt.Sprint(digit)
	}
	return otp
}

// sendEmail.
func EmailSend(otp string, email string) bool {
	userEmailId := []string{email} // set email address.
	data := make(map[string]interface{})
	data["subject"] = "OTP"
	data["email"] = email
	data["opt"] = otp

	//if email success return true.
	if utility.SendEmail(userEmailId, "EmailForOtp", data) {
		return true
	}
	return false
}

// login by email.
func LoginByEmail(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	var signupDetails SignupDetails
	err := utility.StrictParseDataFromJson(r, &signupDetails)
	log.Println("signupDetails: ", signupDetails)
	if err != nil {
		utility.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if signupDetails.Email == "" {
		response.Status = "400"
		response.Message = "Please Enter valid Email Address."
		utility.RenderJsonResponse(w, r, response, 400)
		return true
	}
	//otp generate.
	otp := GenerateOTP()
	//otp set in session.
	utility.SessionSet(w, r, utility.Session{Key: "otp", Value: otp})
	log.Println("otp", otp)
	utility.SessionSet(w, r, utility.Session{Key: "email", Value: signupDetails.Email})

	boolType := EmailSend(otp, signupDetails.Email)
	if boolType {
		response.Status = "200"
		response.Message = "New OTP has been sent, Please check your inbox"
	} else {
		response.Message = "OTP email couldn't be sent at the moment, Please try again."
	}
	utility.RenderJsonResponse(w, r, response, 200)
	return false

}

// call this function only for otp submit.
func MatchOtp(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	otpSession := fmt.Sprintf("%v", utility.SessionGet(r, "otp"))

	var signupDetails SignupDetails
	err := utility.StrictParseDataFromJson(r, &signupDetails)
	log.Println("signupDetails: ", signupDetails)
	if err != nil {
		utility.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return true
	}

	//checking otp if correct or not.
	if otpSession == signupDetails.Otp {
		utility.DeleteSessionValues(w, r, "otp")

		//fetch user details or insert email.
		data, err := models.Authentication{}.GetUserByEmailIds(signupDetails.Email)
		if err != nil {
			response.Status = "500"
			response.Message = "Internal server error, Any serious issues which cannot be recovered from."
			utility.RenderJsonResponse(w, r, response, 400)
			return true
		}

		response.Status = "200"
		response.Message = "Login success."
		response.Payload = data
		utility.RenderJsonResponse(w, r, response, 200)
		return true
	}

	response.Status = "403"
	response.Message = "Please Insert Correct Opt."
	utility.RenderJsonResponse(w, r, response, 400)
	return false
}
