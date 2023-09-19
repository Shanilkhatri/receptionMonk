package utility

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"

	"golang.org/x/crypto/bcrypt"
)

// Template Pool
var View *template.Template

// Session Store
var Store *sessions.FilesystemStore

// DB Connections
var Db *sqlx.DB

type Session struct {
	Key   string
	Value interface{}
}

type Flash struct {
	Type    string
	Message string
}

type AjaxResponce struct {
	Status      string
	Message     string
	Payload     interface{}
	LastId      int64
	TotalRecord int64
}

type Login_auth struct {
	Email    string
	Password string
}
type AjaxRequest struct {
	Token   string
	Userid  string
	Payload map[string]string
	Email   string
}
type Userdata struct {
	Userid         int    `db:"id"`
	Email          string `db:"email"`
	Type           string `db:"type"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	Avatar         string `db:"avatar"`
	Password       string `db:"password"`
	Dob            string `db:"dob"`
	TokenTimestamp int64  `db:"token_timestamp"`
	Token          string `db:"token"`
	PaymentMethod  string `db:"payout_method"`
	PaymentAddress string `db:"payout_address"`
	Phone          string `db:"phone"`
	AjaxRequest    AjaxRequest
}

type JoinLeagueNotification struct {
	LeagueId int64  `json:"leagueid"`
	SportsId int64  `json:"sportsId"`
	Todate   string `json:"todate"`
	Fromdate string `json:"fromdate"`
	Type     string `json:"leaguetype"`
	WeekNo   int    `json:"weekno"`
}
type MatchWinNotification struct {
	MatchId           int64  `json:"matchid"`
	BetId             int64  `json:"betid"`
	WinningTeam       string `json:"winningteam"`
	LossingTeam       string `json:"lossingteam"`
	WinningTeamScores int    `json:"winnigteamscore"`
	LossingTeamScores int    `json:"lossingteamscore"`
}
type ChatNotification struct {
	ChatId   int64  `json:"chatid"`
	ChatType string `json:"chattype"`
	Message  string `json:"chatmessage"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}
type MatchStartEndNotification struct {
	MatchId            int64  `json:"matchid"`
	Team1              string `json:"team1"`
	Team2              string `json:"team2"`
	OddAPICommenceTime int64  `json:"oddAPICommenceTime"`
}
type NotificationStatus struct {
	Subject                   string `json:"subject"`
	Status                    string `json:"status"`
	Image                     string `json:"image"`
	JoinLeagueNotification    JoinLeagueNotification
	MatchWinNotification      MatchWinNotification
	ChatNotification          ChatNotification
	MatchStartEndNotification MatchStartEndNotification
}
type AppNotification struct {
	Id            int64  `db:"id"`
	UserId        int64  `db:"user_id"`
	Message       string `db:"message"`
	Timestamp     int64  `db:"timestamp"`
	Status        string `db:"status"`
	FromTimestamp int64
	ToTimestamp   int64
	Limit         string
	IsRead        bool `db:"isread"`
	ReadAll       bool
	//IsGet         string `db:"isget"`
	NotificationStatus NotificationStatus
}

func RedirectTo(w http.ResponseWriter, r *http.Request, path string) {
	http.Redirect(w, r, os.Getenv("APP_URL")+"/"+path, http.StatusFound)
}

func SessionSet(w http.ResponseWriter, r *http.Request, data Session) {
	session, _ := Store.Get(r, os.Getenv("SESSION_NAME"))
	// Set some session values.
	session.Values[data.Key] = data.Value
	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		Logger(err)
	}
}

func SessionGet(r *http.Request, key string) interface{} {
	session, _ := Store.Get(r, os.Getenv("SESSION_NAME"))
	// Set some session values.
	return session.Values[key]
}

func fetchSession(r *http.Request) map[interface{}]interface{} {
	session, _ := Store.Get(r, os.Getenv("SESSION_NAME"))
	return session.Values
}
func APICheckTokenVerfity(userid string, token string) bool {
	userid = userid + "reak_api"
	match := bcrypt.CompareHashAndPassword([]byte(userid), []byte(token))
	if match != nil {
		return false
	} else {
		return true
	}
}

// function for compare listof allaccessdata to usertype
func stringInSlice(list []string, value string) bool {
	for _, b := range list {
		if b == value {
			return true
		}
	}
	return false
}

// function for getuserprfiledata for verification
func GetProfilData(Userid string) (Userdata, error) {
	var requestedProfileData Userdata
	err := Db.Get(&requestedProfileData, "SELECT `id`,`email`,`type`,`first_name`,IFNULL(`last_name`,'') as last_name, IFNULL(`avatar`,'') as avatar,`password`,`dob`,`token_timestamp`,`token`,IFNULL(`payout_method`,'') as payout_method,IFNULL(`payout_address`,'') as payout_address,IFNULL(`phone`,'') as phone FROM authentication WHERE `id`=?", Userid)
	return requestedProfileData, err
}

// check value exist in array or not
func CheckAclToken(listofallaccess []string, w http.ResponseWriter, r *http.Request) (bool, Userdata) {
	if IsCurlApiRequest(r) {
		authToken := r.Header.Get("authtoken")
		authUserId := r.Header.Get("authuserid")
		if authToken != "" && authUserId != "" {
			getUserProfileData, err := GetProfilData(authUserId)
			if err != nil {
				Logger(err)
			} else {
				userid := strconv.Itoa(int(getUserProfileData.Userid))
				tokenToCompare := userid + getUserProfileData.Email
				if bcrypt.CompareHashAndPassword([]byte(authToken), []byte(tokenToCompare)) == nil && stringInSlice(listofallaccess, getUserProfileData.Type) {
					return true, getUserProfileData
				}
			}
		}
	} else {
		if SessionGet(r, "type") == "admin" {
			userId := fmt.Sprintf("%v", SessionGet(r, "id"))
			return true, Userdata{Type: "admin", Userid: StrToInt(userId)}
		}
	}
	return false, Userdata{}
}

// checkACL function to match token
// func CheckACL(w http.ResponseWriter, r *http.Request, checkUserAcl []string, isTemplate bool) (bool, Userdata) {
// 	userDataIsValid, userdata := CheckAclToken(checkUserAcl, w, r)
// 	if userDataIsValid {
// 		return userDataIsValid, userdata
// 	} else {
// 		response := AjaxResponce{Status: "failure", Message: "403 forbidden", Payload: ""}
// 		var template string
// 		if isTemplate {
// 			template = "forbidden"
// 		}
// 		RenderTemplate(w, r, template, response)
// 	}
// 	return userDataIsValid, userdata
// }

func CheckACL(w http.ResponseWriter, r *http.Request, checkUserAcl []string, isTemplate bool) (bool, Userdata) {
	userDataIsValid, userdata := CheckAclToken(checkUserAcl, w, r)
	if userDataIsValid {
		return userDataIsValid, userdata
	} else {
		response := AjaxResponce{Status: "failure", Message: "403 forbidden", Payload: ""}
		var template string

		if isTemplate {
			template = "forbidden"
		}
		if r.URL.Query().Get("isAjax") == "true" {
			template = ""
		}
		RenderTemplate(w, r, template, response)
	}
	return userDataIsValid, userdata
}

func AddFlash(flavour string, message string, w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, os.Getenv("SESSION_NAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//flash := make(map[string]string)
	//flash["Flavour"] = flavour
	//flash["Message"] = message
	flash := Flash{
		Type:    flavour,
		Message: message,
	}
	err = session.Save(r, w)
	session.AddFlash(flash, "message")
	if err != nil {
		log.Println(err)
	}
}

func viewFlash(w http.ResponseWriter, r *http.Request) interface{} {
	session, err := Store.Get(r, os.Getenv("SESSION_NAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fm := session.Flashes("message")
	if fm == nil {
		return nil
	}
	session.Save(r, w)
	return fm
}

/* if is http@header["reak-api"] return true otherwise false*/
func IsCurlApiRequest(r *http.Request) bool {
	return r.Header.Get("reak-api") == "true"
}

/* if isCurlAPI w.Write json otherwise ExcuteTemplate() */
func RenderTemplate(w http.ResponseWriter, r *http.Request, template string, data interface{}) {
	tmplData := make(map[string]interface{})
	// tmplData = data
	if IsCurlApiRequest(r) || template == "" {
		jsonresponce, err := json.Marshal(data)
		if err != nil {
			Logger(err)
		}
		//debug payload log.Println(string(jsonresponce))
		w.Write([]byte(jsonresponce))
	} else {
		ajaxReq, ok := data.(AjaxResponce)
		if ok {
			AddFlash(ajaxReq.Status, ajaxReq.Message, w, r)
		}
		tmplData["data"] = data
		tmplData["flash"] = viewFlash(w, r)
		tmplData["session"] = fetchSession(r)
		tmplData["appUrl"] = os.Getenv("APPURL")
		tmplData["supportEmail"] = os.Getenv("SUPPORTEMAIL")
		View.ExecuteTemplate(w, template, tmplData)
	}
}

// check user type of given id
func IsCurrectAdminOrUser(requstedvalue AjaxRequest, listACL []string) bool {
	var UserTypeInDb string
	err := Db.Get(&UserTypeInDb, "SELECT type FROM authentication WHERE id = ?", requstedvalue.Userid)
	if err != nil {
		Logger(err)
	}
	return stringInSlice(listACL, UserTypeInDb)
}

// convert string to int
func StrToInt(num string) int {
	if num != "" {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		return intNum
	}
	return 0
}

// convert string to integer with default err response
func StrToIntConversion(w http.ResponseWriter, num string) (int, error) {
	intNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
		response := AjaxResponce{Status: "failure", Message: "enter integer values only", Payload: []interface{}{}}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			Logger(err)
		}
		w.Write([]byte(jsonResponse))
	}
	return intNum, err
}

// convert string to int64
func StrToInt64(str string) (int64, error) {
	strint64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return int64(0), err
	}
	return strint64, err
}
func StrToBool(str string) bool {
	strToConvert, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return strToConvert

}

/* smtp send email*/
func SendEmailSMTP(to []string, subject string, body string) bool {
	//Sender data.
	from := os.Getenv("FROM_EMAIL")
	// Set up email information.
	header := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg := []byte("From: " + from + "\n" + "To: " + strings.Join(to, ",") + "\n" + "Subject: " + subject + "\r\n" + header + body)
	// Sending email.
	// fmt.Println("From: " + from + "\n" + "To: " + strings.Join(to, ",") + "\n" + "Subject: " + subject + "\r\n" + header + "\r\n" + body)
	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), smtp.PlainAuth("", os.Getenv("FROM_APIKEY"), os.Getenv("EMAILSECRATE"), os.Getenv("SMTP_HOST")), from, to, msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*
* main function to call for send mail
#input

	to : string array
	template : tempatePath
	data : associative array of data which is set on template, by-default app_url and app_name is set
*/
func SendEmail(to []string, template string, data map[string]interface{}) bool {
	buf := new(bytes.Buffer)
	//extra information on email
	data["app_url"] = os.Getenv("APPURL")
	data["app_name"] = os.Getenv("APPNAME")
	// Set up email information.
	err := View.ExecuteTemplate(buf, template, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return SendEmailSMTP(to, fmt.Sprintf("%v", data["subject"]), buf.String())
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// upload image file
// upload image file
func UploadFile(r *http.Request, fileName string, controlName string, avatar string) (string, error, string) {
	var message string
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println(err)
	}
	file, _, err := r.FormFile(controlName)
	if err != nil {
		fmt.Println(err)
	} else {
		// files
		defer file.Close()
		tempFile, err := ioutil.TempFile("assets/images/"+avatar, fileName+"*.png")
		if err != nil {
			fmt.Println(err)
		}

		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		fileExtention := http.DetectContentType(fileBytes)

		fileInfo, err := os.Stat(tempFile.Name())
		if err != nil {
			fmt.Println(err)
		}
		//check file size
		fileSize := fileInfo.Size()

		allowedFileSize, err := StrToInt64(os.Getenv("UPLOAD_FILE_SIZE_IN_KB"))

		if err != nil {
			fmt.Println(err)
			allowedFileSize = 50 //IN KB
		}

		if allowedFileSize == 0 {
			err := errors.New("either upload_file_size in kb env variable not found or found null")
			Logger(err)
			allowedFileSize = 50 //IN KB
		}

		//upload file only till 5 mb and fileExtension is jpeg/jpg/png
		if fileSize <= (allowedFileSize*1000) && fileExtention == "image/jpeg" || fileExtention == "image/jpg" || fileExtention == "image/png" {
			return tempFile.Name(), err, message
		} else {
			// to-do calcualate image size and apply remove file
			message = "Image size should not be greater than  " + fmt.Sprint(allowedFileSize) + "kb and accepted formats are jpeg,jpg,png"
		}
	}
	return "", err, message
}

/*
ODDS API DOCS
https://api.the-odds-api.com/v4/sports/SPORTS_KEY/scores/?daysFrom=1&apiKey=YOUR_API_KEY
*/
// Post match Api data  fatch and convert in bytes & this function return bytes[] & error

func PostMatchApiByteData(key string, oddsApiKey string) ([]byte, error) {
	var apiDataInByte []byte
	requestForOddsData := os.Getenv("ODDSAPI_MATCH_API_URL") + "sports/" + key + "/scores/?daysFrom=" + os.Getenv("MATCHSCORE_DAYS_BEFORE") + "&apiKey=" + oddsApiKey
	//requestForOddsData := "http://localhost:8000/winMatchNewApi.json"
	apiData, err := http.Get(requestForOddsData)
	if err != nil {
		Logger(err)
	} else {
		defer apiData.Body.Close()
		//Read all request data
		apiDataInByte, err = ioutil.ReadAll(apiData.Body)
		if err != nil {
			fmt.Println(err)
			Logger(err)
		}
	}
	// file name and date,time
	fileName := "winmatchdata/winmatchAPIdata_"
	// add data in json file convert byte to string
	GetJsonDataAddFile(string(apiDataInByte), fileName)
	return apiDataInByte, err
}

// Generate NewPasswordHash
func NewPasswordHash(NewPassword string) (string, error) {
	//NewPassword Change bcrypt code
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(NewPassword), 10)
	//modify NewPassword
	NewPassword = string(newPasswordHash)
	if err != nil || NewPassword == "" {
		Logger(err)
	} else {
		return NewPassword, err
	}
	return "", err
}

func ParseDate(date string) int64 {
	if date != "" {
		toTime, err := time.Parse("02-01-2006", date)
		if err != nil {
			fmt.Println(err)
		}
		return toTime.Unix()
	} else {
		return 0
	}
}

func SessionDestroy(w http.ResponseWriter, r *http.Request) bool {
	session, err := Store.Get(r, os.Getenv("SESSION_NAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func FloatAmountToInt(amount float64) int64 {
	// intAmount := amount * 100
	integerAmount := int64(math.Round(amount))
	return integerAmount
}

func Float64ToInt64(amount float64) int64 {
	integerAmount := int64(amount * 100)
	return integerAmount
}

func ForGetDataInFloat(amount float64) float64 {
	floatAmount := (float64(amount) / 100)
	return floatAmount

}

/* Go: email sent of Critical Error Message*/
func GetErrorMessage(currentFilePath string, lineNumbers int, errorMessage error) bool {
	emailForErrorMessageSend := os.Getenv("EMAIL_FOR_CRITICAL_ERROR")
	email := []string{emailForErrorMessageSend} // set email address
	data := make(map[string]interface{})
	data["subject"] = "Error message " + os.Getenv("APPNAME")
	data["errorMessage"] = errorMessage
	data["currentFilePath"] = currentFilePath
	data["lineNumbers"] = lineNumbers
	if !SendEmail(email, "errorMessage", data) {
		fmt.Println("Error log Email couldn't be sent at the moment")
	}
	return false
}

var LogFile *os.File
var ErrorLog *log.Logger

/*Go:Open file for Log Critical Error Message */
func OpenLogFile() *os.File {
	loggedErrorPath := os.Getenv("PATH_OF_LOG_API")
	if loggedErrorPath == "" {
		loggedErrorPath = "APIdata/"
		fmt.Println("Error Log File: env variable not found")
	}
	logFileURI := loggedErrorPath + "loggederror/Errorlogged.txt"
	LogFile, err := os.OpenFile(logFileURI, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error Log File: ", err)
	}
	ErrorLog = log.New(LogFile, "[DANGER] ", log.Llongfile|log.LstdFlags) // set new logger ErrorLog because of Logger is Reserved word

	fmt.Println("Log File Open: ", logFileURI)
	return LogFile
}

/* Go: Log Critical Error Message on file if CI_ENVIRONMENT is production in env file then send email to EMAIL_FOR_CRITICAL_ERROR */
func Logger(errObject error) {
	if errObject != nil { // null checking because of stuck server when error is null
		//using 1 indicate actually error
		_, currentFilePath, lineNumbers, ok := runtime.Caller(1)
		if !ok {
			err := errors.New("failed to get filename")
			fmt.Println("Error Log File: ", err)
		}
		ErrorLog.Output(2, errObject.Error())

		if os.Getenv("CI_ENVIRONMENT") == "production" { // when development environment is set email not to be sent to developer because of this rise a error
			go GetErrorMessage(currentFilePath, lineNumbers, errObject)
		}
	}
}

/*
	remove files from the specified directory

ex- RemoveFile("assets/images/market/testing3375083430.png")
*/
func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}
}

// get sql error string from sql error
func GetSqlErrorString(err error) string {
	mes := strings.SplitN(err.Error(), ":", -1)
	return mes[1]
}

// match sql error with particular sql error string
func CheckSqlError(err error, errString string) (bool, string) {
	sqlerrorString := GetSqlErrorString(err)
	exists := strings.HasPrefix(sqlerrorString, errString)
	return exists, sqlerrorString
}

// func PutNotification(notify AppNotification, noti NotificationStatus) (bool, error) {
// 	var err error
// 	notify.Status, err = StatusUnmarshalToString(noti)
// 	if err != nil {
// 		Logger(err)
// 		return false, err
// 	}
// 	go TriggerPushNotification(notify)

// 	if notify.UserId >= 0 && notify.Message != "" && notify.Timestamp > int64(0) {
// 		rows, err := Db.NamedExec("INSERT INTO `notification` (message,timestamp,status) VALUES (:Message,:TimeStamp,:Status)", map[string]interface{}{"Message": notify.Message, "TimeStamp": notify.Timestamp, "Status": notify.Status})
// 		if err != nil {
// 			Logger(err)
// 			return false, err
// 		} else {
// 			//Fetch lastInsertId
// 			rowAffected, err := rows.LastInsertId()
// 			if err != nil {
// 				Logger(err)
// 			} else {
// 				notify.Id = rowAffected
// 				boolType, err := PutMultiUserNotification(notify)
// 				return boolType, err
// 			}
// 			return true, err
// 		}
// 	}
// 	return false, err
// }

type NotificationCondition struct {
	WhereCondition string
	AppNotification
}

// func GetNotification(w http.ResponseWriter, r *http.Request) bool {
// 	response := AjaxResponce{Status: "failure", Message: "Operation couldn't be performed, Please try again after some time", Payload: []interface{}{}}
// 	var param NotificationCondition

// 	var err error
// 	var notification []AppNotification

// 	param.UserId, err = StrToInt64(r.URL.Query().Get("userid"))
// 	if err != nil {
// 		Logger(err)
// 		RenderTemplate(w, r, "", response)
// 		return false
// 	}
// 	fromTime := r.URL.Query().Get("fromtimestamp")
// 	if fromTime != "" {
// 		param.FromTimestamp, err = StrToInt64(fromTime)
// 		if err != nil {
// 			Logger(err)
// 			RenderTemplate(w, r, "", response)
// 			return false
// 		}
// 	}
// 	toTime := r.URL.Query().Get("totimestamp")
// 	if toTime != "" {
// 		param.ToTimestamp, err = StrToInt64(toTime)
// 		if err != nil {
// 			Logger(err)
// 			RenderTemplate(w, r, "", response)
// 			return false
// 		}
// 	}

// 	id := r.URL.Query().Get("id")

// 	if id != "" {
// 		param.Id, err = StrToInt64(id)
// 		if err != nil {
// 			Logger(err)
// 			RenderTemplate(w, r, "", response)
// 			return false
// 		}
// 	}

// 	param.Status = r.URL.Query().Get("status")
// 	limit := strings.ToLower(r.URL.Query().Get("limit"))
// 	//if limit empty or zero set by default limit 5
// 	if limit == "" || limit == "0" {
// 		limit = os.Getenv("DEFAULT_LIMIT_PAGINATION")
// 		//if env variable no found or limit value is empty string
// 		if limit == "" {
// 			limit = "5"
// 			log.Println("env variable (DEFAULT_LIMIT_PAGINATION) not found")
// 		}
// 	}

// 	params := GetParamsForFilterNotification(param)
// 	query := "SELECT user_notification.id,user_notification.user_id,user_notification.isread ,notification.message, notification.timestamp, notification.status FROM `notification` as notification INNER JOIN multi_user_notification as user_notification ON user_notification.notification_id = notification.id Where 1=1 " + params.WhereCondition + " Order by notification.timestamp DESC,user_notification.`id` DESC LIMIT " + limit
// 	condtion := map[string]interface{}{
// 		"userId":        params.AppNotification.UserId,
// 		"toTimestamp":   params.AppNotification.ToTimestamp,
// 		"fromTimestamp": params.AppNotification.FromTimestamp,
// 		"id":            params.AppNotification.Id,
// 		"status":        params.AppNotification.Status,
// 	}
// 	rows, err := Db.NamedQuery(query, condtion)
// 	if err != nil {
// 		Logger(err)
// 		return false
// 	} else {
// 		defer rows.Close()
// 		for rows.Next() {
// 			var selectSingleRow AppNotification
// 			err := rows.StructScan(&selectSingleRow)
// 			if err != nil {
// 				Logger(err)
// 				return false
// 			}
// 			notification = append(notification, selectSingleRow)

// 		}
// 		//notification data is zero status success but recored not found
// 		if len(notification) != 0 {
// 			response.Status = "success"
// 			response.Message = "Record found successfully"
// 			response.Payload = notification
// 			//last show id
// 			response.LastId = int64(notification[len(notification)-1].Id)

// 		} else {
// 			response.Status = "success"
// 			response.Message = "No notification were found for this search"
// 		}
// 		RenderTemplate(w, r, "", response)
// 		return true
// 	}
// }

// // FUNCTION TO CHECK THE PARAMETER PASSED
// func GetParamsForFilterNotification(params NotificationCondition) NotificationCondition {
// 	if params.UserId != 0 {
// 		params.WhereCondition += " AND user_notification.user_id=:userId "
// 	} else {
// 		//if user id zero
// 		//#TO_DO this is currently not used but i need this statement remove comment.
// 		//params.WhereCondition += " AND user_notification.user_id=:userId AND notification.status=:status"

// 	}
// 	if params.ToTimestamp != 0 && params.FromTimestamp != 0 {
// 		params.WhereCondition += " AND notificaion.timestamp BETWEEN :fromTimestamp AND :toTimestamp "
// 	}
// 	if params.ToTimestamp != 0 {
// 		params.WhereCondition += " AND notificaion.timestamp<=:toTimestamp"
// 	}
// 	if params.FromTimestamp != 0 {
// 		params.WhereCondition += " AND notificaion.timestamp>=:fromTimestamp"
// 	}

// 	if params.Id != 0 {
// 		params.WhereCondition += " AND user_notification.id <:id "
// 	}
// 	//#
// 	if params.Status == "chat" {
// 		params.WhereCondition += " AND JSON_EXTRACT(status, '$.subject')=:status "
// 	}

// 	return params
// }

func StrToFloat64(str string) (float64, error) {
	strfloat64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		return float64(0), err
	}
	return strfloat64, err
}

// match ,matchwin and betcalculation function use AuthKeyChecking
func AuthKeyCheking(w http.ResponseWriter, r *http.Request) bool {
	response := AjaxResponce{Status: "failure", Message: "Access denied", Payload: []interface{}{}} // by default response
	authKey := r.Header.Get("authkey")
	authKeybyenv := os.Getenv("GET_AUTHKEY_BYENV")
	//chech env variable not empty
	if authKeybyenv == "" {
		log.Println("missing env variable")
		return true
	}
	//if authkey and authkeybyenv not match show error message access denied
	if authKey == authKeybyenv {
		return true
	}
	response.Message = "Access denied"
	RenderTemplate(w, r, "", response)
	return false
}

// update notification
// func UpdateNotification(w http.ResponseWriter, r *http.Request) bool {

// 	response := AjaxResponce{Status: "failure", Message: "Operation couldn't be performed, Please try again after some time", Payload: []interface{}{}}
// 	notificationData, err := JsonDecoder(r)
// 	if err != nil {
// 		Logger(err)
// 		response.Message = "Record couldn't be updated"
// 		RenderTemplate(w, r, "", response)
// 		return true
// 	}
// 	var notificationupdate sql.Result

// 	//nofication validation
// 	if notificationData.Id != 0 || notificationData.UserId != 0 {
// 		//if update all data based on userId
// 		if notificationData.ReadAll {
// 			notificationupdate, err = Db.NamedExec("UPDATE `multi_user_notification` SET isread=:IsRead WHERE user_id=:UserId ", map[string]interface{}{"IsRead": notificationData.IsRead, "UserId": notificationData.UserId})
// 		} else {
// 			notificationupdate, err = Db.NamedExec("UPDATE `multi_user_notification` SET isread=:IsRead WHERE id=:Id AND user_id=:UserId ", map[string]interface{}{"IsRead": notificationData.IsRead, "Id": notificationData.Id, "UserId": notificationData.UserId})
// 		}

// 		// Check error
// 		if err != nil {
// 			Logger(err)
// 			response.Message = "Record couldn't be updated"
// 			RenderTemplate(w, r, "", response)
// 			return true
// 		}
// 		rowEffect, _ := notificationupdate.RowsAffected()
// 		//if row update message success
// 		if rowEffect > 0 {
// 			response.Status = "success"
// 			response.Message = "Record successfully updated"
// 			//response.Payload = []interface{}{}
// 		} else {
// 			response.Status = "success"
// 			response.Message = ""
// 		}
// 		//GetNotification(w, r)
// 		RenderTemplate(w, r, "", response)
// 	}
// 	return false
// }

// json decode
func JsonDecoder(r *http.Request) (AppNotification, error) {
	var notiofication AppNotification
	err := json.NewDecoder(r.Body).Decode(&notiofication)
	return notiofication, err
}

// func NotificationCount(w http.ResponseWriter, r *http.Request) bool {
// 	response := AjaxResponce{Status: "failure", Message: "Operation couldn't be performed, Please try again after some time", Payload: []interface{}{}}
// 	notificationMap := make(map[string]int64)
// 	var totalNotification int64 = 0
// 	var param NotificationCondition
// 	var err error

// 	param.UserId, err = StrToInt64(r.URL.Query().Get("userid"))
// 	if err != nil {
// 		Logger(err)
// 		RenderTemplate(w, r, "", response)
// 		return false
// 	}
// 	messageReadUnreadBool := r.URL.Query().Get("isread")
// 	if messageReadUnreadBool != "" {
// 		if strings.ToLower(messageReadUnreadBool) == "true" {
// 			param.IsRead = true
// 		} else {
// 			param.IsRead = false
// 		}
// 	}

// 	err = Db.QueryRow("SELECT COUNT(isread) as totalNotification FROM `multi_user_notification` WHERE isread = ? AND user_id = ?", messageReadUnreadBool, param.UserId).Scan(&totalNotification)
// 	if err != nil {
// 		Logger(err)
// 		RenderTemplate(w, r, "", response)
// 		return false
// 	} else {
// 		response.Status = "success"
// 		response.Message = ""
// 		if param.IsRead {
// 			notificationMap["ReadedNotification"] = totalNotification
// 		}
// 		if !param.IsRead {
// 			notificationMap["UnreadedNotification"] = totalNotification
// 		}
// 		response.Payload = notificationMap
// 	}

// 	RenderTemplate(w, r, "", response)
// 	return true

// }

// Read json data and add file
func GetJsonDataAddFile(jsonData string, fileName string) bool {
	apiDataPath := os.Getenv("PATH_OF_LOG_API")
	if apiDataPath == "" {
		apiDataPath = "APIdata/"
		log.Println("env variable not found")
	}

	LOG_FILE := apiDataPath + fileName + time.Now().Format("2006-01-02 15:04:05") + ".txt"

	err := ioutil.WriteFile(LOG_FILE, []byte(jsonData), 0644)
	if err != nil {
		Logger(err)
		return true
	}
	// // open log file
	// logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	// //error
	// if err != nil {
	// 	log.Println(err, "err")
	// 	// Logger(err)
	// 	return true
	// }
	// defer logFile.Close()
	// //write data in file
	// _, err2 := logFile.WriteString(jsonData)
	// if err2 != nil {
	// 	log.Println(err2)
	// }
	return false
}

// type ReadNotification struct {
// 	Notification struct {
// 		Signup []struct {
// 			Message string `json:"message"`
// 		}
// 		Parlaywinmatch []struct {
// 			Message string `json:"messagewin"`
// 		}

// 		Parlaylossmatch []struct {
// 			Message string `json:"messageloss"`
// 		}
// 		MatchWinNormal []struct {
// 			Message string `json:"normalmessagewin"`
// 		}

// 		MatchLossNormal []struct {
// 			Message string `json:"normalmessageloss"`
// 		}

// 		Matchstart []struct {
// 			Message string `json:"startmessage"`
// 		}

// 		LeagueJoin []struct {
// 			Message string `json:"messagejoin"`
// 		}
// 		LeagueStart []struct {
// 			Message string `json:"activemsg"`
// 		}

// 		Trending []struct {
// 			Message string `json:"messagetrending"`
// 		}
// 	}
// }

// // read notification message in json file.
// func ReadNotificationByJsonFile(msgStr string) string {
// 	requestForNotifications := os.Getenv("APPURL") + "/notification/notification.json"
// 	resp, err := http.Get(requestForNotifications)
// 	//check error
// 	if err != nil {
// 		Logger(err)
// 	}
// 	defer resp.Body.Close()
// 	//Read all request data
// 	responseData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		Logger(err)
// 	}
// 	var jsonUnmarshal ReadNotification
// 	//Decode Json data
// 	err = json.Unmarshal(responseData, &jsonUnmarshal)
// 	if err != nil {
// 		Logger(err)
// 	}

// 	var message struct{ Message string }
// 	//check validation only call specific function.
// 	if msgStr == "signUp" {
// 		message = SignupMessageRandomly(jsonUnmarshal)
// 	}

// 	if msgStr == "parlayMatchWin" {
// 		message = ParlayMatchWinMessageRandomly(jsonUnmarshal)

// 	}
// 	if msgStr == "parlayMatchLoss" {
// 		message = ParlayMatchLossMessageRandomly(jsonUnmarshal)

// 	}
// 	if msgStr == "normalMatchWin" {
// 		message = NormalMatchWinMessageRandomly(jsonUnmarshal)

// 	}
// 	if msgStr == "normalMatchLoss" {
// 		message = NormalMatchLossMessageRandomly(jsonUnmarshal)

// 	}
// 	if msgStr == "leagueJoin" {

// 		message = LeagueJoinMessageRandomly(jsonUnmarshal)
// 	}

// 	if msgStr == "leaguestart" {

// 		message = LeagueMessageRandomly(jsonUnmarshal)
// 	}

// 	if msgStr == "matchStart" {

// 		message = MatchStartMessageRandomly(jsonUnmarshal)
// 	}

// 	return message.Message
// }

// // Randomly Get signUp message
// func SignupMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {

// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.Signup))
// }

// // Randomly Get Parlay match Win message
// func ParlayMatchWinMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.Parlaywinmatch))
// }

// // Randomly Get Parlay match loss message
// func ParlayMatchLossMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.Parlaylossmatch))
// }

// // Randomly Get Normal match win message
// func NormalMatchWinMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.MatchWinNormal))
// }

// // Randomly Get normal match loss message
// func NormalMatchLossMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.MatchLossNormal))
// }

// // Randomly Get League join message
// func LeagueJoinMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.LeagueJoin))
// }

// func LeagueMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.LeagueStart))
// }

// // Randomly Get Match start message
// func MatchStartMessageRandomly(notificationstruct ReadNotification) struct{ Message string } {
// 	return Randomly([]struct{ Message string }(notificationstruct.Notification.Matchstart))
// }

// // Randomly Get signUp message
// func Randomly(notificationstruct []struct{ Message string }) struct{ Message string } {
// 	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(notificationstruct))))
// 	if err != nil {
// 		Logger(err)
// 	}
// 	return notificationstruct[num.Int64()]
// }

//#TO_DO upadate notification if user id is zero .
// func UpdateNotificationForStartMatch(notificationData AppNotification) {
// 	_, err := Db.NamedExec("UPDATE notification SET isget='deactive' WHERE user_id=:UserId AND status=:Status", map[string]interface{}{"IsGet": notificationData.IsGet, "UserId": notificationData.UserId, "Status": notificationData.Status})
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// return time diffrence in minutes.
func ReturnTimeInMinutes(currentTime time.Time, matchTime time.Time) int {

	difference := matchTime.Sub(currentTime)

	minutes := int((difference.Minutes()))

	return minutes
}

// // add notification details in multi user tables.
// func PutMultiUserNotification(multiUser AppNotification) (bool, error) {
// 	userIdArr, _ := GetUser()

// 	if multiUser.UserId == 0 {
// 		for _, singleId := range userIdArr {
// 			multiUser.UserId = int64(singleId)
// 			InsertMultiUser(multiUser)
// 		}
// 		return true, nil
// 	} else {
// 		return InsertMultiUser(multiUser)
// 	}
// }

// func InsertMultiUser(multiUser AppNotification) (bool, error) {
// 	_, err := Db.NamedExec("INSERT INTO `multi_user_notification` (`user_id`,`notification_id`) VALUES (:UserId,:Id)", map[string]interface{}{"UserId": multiUser.UserId, "Id": multiUser.Id})
// 	if err != nil {
// 		Logger(err)
// 		return false, err
// 	}
// 	return true, nil
// }

// func GetUser() ([]int64, error) {
// 	var userDataArray []int64
// 	rows, err := Db.Query("SELECT `authentication`.`id` FROM `authentication`")
// 	if err != nil {
// 		Logger(err)
// 		return userDataArray, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var singleRow int64
// 		err := rows.Scan(&singleRow)
// 		if err != nil {
// 			Logger(err)
// 			return userDataArray, err
// 		}
// 		userDataArray = append(userDataArray, singleRow)
// 	}
// 	return userDataArray, err
// }

type ReplaceKeys struct {
	ImageSRCNew      string
	GoldCoins        string
	SingleDataAmount string
	SubbetArr        string
	Team1Normal      string
	Team2Normal      string
	Contest          string
	Timedifference   string
	LeagueName       string
	Team1            string
	Team2            string
	WeekNumber       string
	SilverCoinCredit string
}

// all json file keys for replace values.
func SetKeys(replaceValue ReplaceKeys) map[string]string {
	data := map[string]string{}
	data["imageSRCnew"] = replaceValue.ImageSRCNew
	data["goldCoins"] = replaceValue.GoldCoins
	data["singleDataAmount"] = replaceValue.SingleDataAmount
	data["subbetArr"] = replaceValue.SubbetArr
	data["team1Normal"] = replaceValue.Team1Normal
	data["team2Normal"] = replaceValue.Team2Normal
	data["contest"] = replaceValue.Contest
	data["timedifference"] = replaceValue.Timedifference
	data["leagueName"] = replaceValue.LeagueName
	data["team1"] = replaceValue.Team1
	data["team2"] = replaceValue.Team2
	data["weekNumber"] = replaceValue.WeekNumber
	data["silverCoinCredit"] = replaceValue.SilverCoinCredit
	return data
}

// replace all message key by dynamically.
func ReplaceMessageKeys(message string, value map[string]string) string {

	re := regexp.MustCompile(`\{{.*?}}`)
	allStr := re.FindAllString(message, -1)

	for _, element := range allStr {
		data := map[string]string{}
		data["key"] = element

		element = strings.Trim(element, "{{")
		element = strings.Trim(element, "}}")

		for i, singleValue := range value {

			if element == i {

				message = strings.Replace(message, data["key"], singleValue, -1)

			}

		}

	}
	return message
}

// string convert to error object.
func StrToErr(errString string) error {
	err := errors.New(errString)
	return err
}

// func PushNotification(NotificationString string, mid string) {

// 	authenticationKeyFileAddress := os.Getenv("PUSH_NOTIFICATION_AUTHKEY") // Load from ENV
// 	authKey, err := token.AuthKeyFromFile(authenticationKeyFileAddress)
// 	if err != nil {
// 		Logger(err)
// 		log.Fatal("token error:", err)
// 	}
// 	keyid := os.Getenv("KEY_ID")
// 	teamId := os.Getenv("TEAM_ID")
// 	notificationTopic := os.Getenv("NOTIFICATION_TOPIC")

// 	customtoken := &token.Token{
// 		AuthKey: authKey, // get by file
// 		KeyID:   keyid,   // get by env
// 		TeamID:  teamId,  // get by env
// 	}

// 	notification := &apns2.Notification{}
// 	notification.DeviceToken = mid

// 	notification.Topic = notificationTopic // get by env

// 	//notification.Payload = []byte(`{"aps":{"alert":"Hello PushTest!", "sound": "default"}}`) // See Payload section below

// 	// example string  //"{\"aps\":{\"alert\":{\"title\":\"Hello PushTest!\",\"subtitle\":\"Subtitle\",\"body\":\"Message body\"},\"sound\":\"default\"},\"Status\":{\"subject\":\"match\",\"status\":\"matchstart\",\"matchid\":11,\"leagueid\":15}}"

// 	notification.Payload = []byte(NotificationString)

// 	// If you want to test push notifications for builds running directly from XCode (Development), use
// 	// client := apns2.NewClient(cert).Development()
// 	// For apps published to the app store or installed as an ad-hoc distribution use Production()

// 	client := apns2.NewTokenClient(customtoken)
// 	_, err = client.Push(notification)

// 	if err != nil {
// 		Logger(err)
// 		log.Fatal("Error:", err)
// 	}

// 	//fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)

// }

// func TriggerPushNotification(notify AppNotification) {
// 	var mid string
// 	var appLastUpdatedTimestap int64

// 	MessageString := `{"aps":{"alert":{"title":"Hello PushTest!","subtitle":"Subtitle","body":"Message body"},"sound":"default"},"Status":` + notify.Status + `}`

// 	finalString := strconv.Quote(MessageString)

// 	err := Db.QueryRow("SELECT authentication.device_token,authentication.last_update_timestamp FROM `authentication` WHERE authentication.id = ?", notify.UserId).Scan(&mid, &appLastUpdatedTimestap)
// 	if err != nil {
// 		Logger(err)
// 	}
// 	sendNotification := IsWithin30SecondsNotification(appLastUpdatedTimestap) // check user app last updated time is greter than 30 sec or not

// 	if sendNotification {
// 		log.Println("sendNotification", sendNotification)
// 		PushNotification(finalString, mid)

// 	}

// }

// let currentTime = 100 appupdatedTime = 60 waitingTime = 30 { 100 - 60 > 30 } return true so send Notification  }
func IsWithin30SecondsNotification(timestamp int64) bool {
	var waitingSecString int64 = 0
	var err error

	waitingSecString, err = StrToInt64(os.Getenv("PUSHNOTIFICATION_WAITING_SEC"))
	if err != nil {
		Logger(err)
	}

	if waitingSecString == 0 {
		waitingSecString = 30
		err := errors.New("env variable missing PUSHNOTIFICATION_WAITING_SEC")
		fmt.Println("Error Log File: ", err)
	}
	currentTimestamp := time.Now().Unix()
	return currentTimestamp-timestamp > waitingSecString

}
func StatusUnmarshalToString(notify NotificationStatus) (string, error) {
	jsonStr, err := json.Marshal(notify)
	if err != nil {
		Logger(err)
		return "", err
	}
	return string(jsonStr), err
}

func TimestampToDateString(timestamp int64) string {
	// Convert the timestamp to a time.Time object
	tm := time.Unix(timestamp, 0)

	// Format the time as a string in a desired layout
	dateString := tm.Format("2006-01-02")
	//dateString := tm.Format("02-01-2006")
	return dateString
}
func GetNotificationRandomMessage() ([]byte, error) {
	var content []byte
	var err error
	filePath := "./notification/notification.json"
	// Read the file content
	content, err = ioutil.ReadFile(filePath)
	if err != nil {
		Logger(err)
		return content, err
	}
	return content, err
}
