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
			Helper.RenderTemplate(w, r, "login", "demo")
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
					Helper.AddFlash("error", "Credentials didn't match, Please try again.", w, r)
					Helper.RenderTemplate(w, r, "login", "demo")
				} else {
					match := bcrypt.CompareHashAndPassword([]byte(row.PasswordHash), []byte(r.FormValue("password")))
					if match != nil {
						Helper.AddFlash("error", "Credentials didn't match, Please try again.", w, r)
						Helper.RenderTemplate(w, r, "login", "demo")
					} else {
						// Password match has been a success
						Helper.SessionSet(w, r, utility.Session{Key: "id", Value: row.ID})
						Helper.SessionSet(w, r, utility.Session{Key: "email", Value: row.Email})
						Helper.SessionSet(w, r, utility.Session{Key: "type", Value: row.AccountType})
						Helper.AddFlash("success", "Success !, Logged in.", w, r)
						if r.FormValue("rememberMe") == "yes" {
							utility.Store.Options = &sessions.Options{
								MaxAge: 60 * 1,
							}
						}
						token := models.Authentication{}.CheckTwoFactorRegistration(int32(row.ID))
						Helper.SessionSet(w, r, utility.Session{Key: "2faSecret", Value: token})

						if token != "" {
							Helper.RedirectTo(w, r, "verify-2fa")
						} else {
							Helper.RedirectTo(w, r, "dashboard")
						}
						//utility.RenderTemplate(w, r, "login", "demo")
					}
				}
			}
		}
	} else {
		Helper.RenderTemplate(w, r, "login", "demo")
	}
}

func VerifyTwoFa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Helper.RenderTemplate(w, r, "verifyTwoFa", "demo")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			twoFaVerify := r.FormValue("twoFaVerify")
			secret := fmt.Sprintf("%v", Helper.SessionGet(r, "2faSecret"))
			if totp.Validate(twoFaVerify, secret) {
				Helper.RedirectTo(w, r, "dashboard")
			} else {
				Helper.RedirectTo(w, r, "login")
			}
		}
	}
}

func RegisterTwoFa(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		email := Helper.SessionGet(r, "email")
		key, _ := totp.Generate(totp.GenerateOpts{
			Issuer:      os.Getenv("TOTP_ISSUER"),
			AccountName: fmt.Sprintf("%v", email),
		})
		Helper.SessionSet(w, r, utility.Session{Key: "totpSecret", Value: key.Secret()})

		img, _ := key.Image(400, 400)

		var buf bytes.Buffer
		png.Encode(&buf, img)

		data := b64.StdEncoding.EncodeToString([]byte(buf.String()))
		data = "data:image/png;base64," + data

		Helper.RenderTemplate(w, r, "twoFactorRegister", data)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {

		} else {
			verifyToken := r.FormValue("challengeCode")
			validationResult := totp.Validate(verifyToken, fmt.Sprintf("%v", Helper.SessionGet(r, "totpSecret")))
			if validationResult {
				secret := fmt.Sprintf("%v", Helper.SessionGet(r, "totpSecret"))
				userId := fmt.Sprintf("%v", Helper.SessionGet(r, "id"))
				intUserId, _ := strconv.Atoi(userId)
				models.Authentication{}.TwoFactorAuthAdd(secret, intUserId)
				Helper.RenderTemplate(w, r, "successTwoFactor", nil)
			} else {
				// Show Error Page
				Helper.RenderTemplate(w, r, "failureTwoFactor", nil)
			}
		}
	}
}

// Generate Otp .
func GenerateOTP() (string, int64, int64) {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10) // Generate a random digit (0-9).
		otp += fmt.Sprint(digit)
	}

	currentTime := time.Now().Unix()
	expirationTime := time.Now().Add(10 * time.Minute).Unix()

	return otp, expirationTime, currentTime
}

// sendEmail.
func EmailSend(otp string, signupData models.SignupDetails) bool {
	userEmailId := []string{signupData.Email} // set email address.
	data := make(map[string]interface{})
	data["subject"] = "OTP verification"
	data["email"] = signupData.Email
	data["otp"] = otp
	data["currentTime"] = signupData.EpochCurrent

	//if email success return true.
	count, isok, err := Helper.SendEmail(userEmailId, "emailforotp", data)
	// count, isok, _ := Helper.SendEmail(userEmailId, "emailforotp", data)
	if isok {
		return true
	} else {
		if count >= Helper.StrToInt(os.Getenv("RETRIES_BEFORE_CRITICAL_EMAIL")) {
			// trigger crictical mail
			Helper.Logger(err, true)
			// resetting value
			utility.Count = 0
			return false
		}
		return false
	}
}

// login by email.
func LoginByEmail(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	var signupDetails models.SignupDetails
	err := Helper.StrictParseDataFromJson(r, &signupDetails)
	if err != nil {
		log.Println("error: ", err)
		Helper.Logger(err, false)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if signupDetails.Email == "" {
		response.Status = "400"
		response.Message = "Please Enter valid Email Address."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	//otp set in session.
	// utility.SessionSet(w, r, utility.Session{Key: "otp", Value: otp})
	// utility.SessionSet(w, r, utility.Session{Key: "email", Value: signupDetails.Email})
	//otp generate.
	boolValues, err, signupDetails := HelpingPostUser(signupDetails)
	if err != nil {
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	if boolValues {
		go EmailSend(signupDetails.Otp, signupDetails)
		// log.Println("Global faulty mail count: ", utility.Count)
		if utility.Count >= Helper.StrToInt(os.Getenv("RETRIES_BEFORE_CRITICAL_EMAIL"))-1 {
			response.Status = "400"
			response.Message = "Can't send OTP at the moment!"
			response.Payload = signupDetails.EmailToken
			Helper.RenderJsonResponse(w, r, response, 400)
			return false
		}
		response.Status = "200"
		response.Message = "New OTP has been sent, Please check your inbox"
		response.Payload = signupDetails.EmailToken
		Helper.RenderJsonResponse(w, r, response, 200)
		return false

	}
	return false
}

// call this function only for otp submit.
func MatchOtp(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	// otpSession := fmt.Sprintf("%v", utility.SessionGet(r, "otp"))
	emailToken := r.Header.Get("emailVerfToken")
	log.Println("emailVerfToekn: ", emailToken)
	var signupDetails models.SignupDetails
	err := Helper.StrictParseDataFromJson(r, &signupDetails)
	if err != nil {
		Helper.Logger(err, false)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	data, err := models.Authentication{}.GetUserDetailsByEmail(signupDetails.Email)
	if err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	if emailToken != data.EmailToken {
		response.Status = "400"
		response.Message = "Email token not valid."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	//checking otp if correct or not.
	if data.Otp == signupDetails.Otp {
		currentTime := time.Now().Unix() // data.EpochCurrent
		// Check if the current time is within 10 minutes from the expiration time.

		if currentTime <= data.EpochExpired {
			fmt.Println("OTP is still valid")
			response.Status = "200"
			response.Message = "Login success."
			// generate a token for header auth and send in payload
			// -> using the same function to generate main token
			// -> that we used for email verf token
			// -> mainAuthToken = bcryptOf(data.email + data.otp)
			data.Token = GenerateEmailToken(data.Email + data.Otp)
			// saving the token in db
			_, err := models.Authentication{}.PostOrPutUserByEmailIds(data)
			if err != nil {
				response.Status = "500"
				response.Message = "Internal server error, Any serious issues which cannot be recovered from."
				Helper.RenderJsonResponse(w, r, response, 500)
				return true
			}
			response.Payload = data.Token
			Helper.RenderJsonResponse(w, r, response, 200)
			return true
		} else {
			fmt.Println("OTP has expired")
			response.Status = "403"
			response.Message = "Please Insert Correct Opt."
			Helper.RenderJsonResponse(w, r, response, 403)
			return false
		}
	} else {
		response.Status = "403"
		response.Message = "Please Insert Correct Opt."
		Helper.RenderJsonResponse(w, r, response, 403)
		return false
	}

}

// generate Email Token.
func GenerateEmailToken(userid string) string {
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(userid), 10)
	if err != nil {
		log.Println(err)
	}
	return string(newPasswordHash)
}

// only for test emailforotp template run.
func Add(w http.ResponseWriter, r *http.Request) {
	Helper.RenderTemplate(w, r, "emailforotp", nil)
}
func HelpingPostUser(signupDetails models.SignupDetails) (bool, error, models.SignupDetails) {
	otp, expirationTime, currentTime := GenerateOTP()

	signupDetails.EpochCurrent = currentTime
	signupDetails.EpochExpired = expirationTime
	emailToken := signupDetails.Email + otp + strconv.FormatInt(currentTime, 10)
	signupDetails.EmailToken = GenerateEmailToken(emailToken)
	signupDetails.Otp = otp
	signupDetails.Token = GenerateEmailToken(signupDetails.EmailToken + signupDetails.Otp)
	boolValues, err := models.Authentication{}.PostOrPutUserByEmailIds(signupDetails)
	return boolValues, err, signupDetails
}
