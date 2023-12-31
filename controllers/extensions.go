package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"

	"github.com/go-sql-driver/mysql"
)

func PostExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var extensionStruct models.Extensions
	// var userPayload models.Users

	isOk, userPayload := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if userPayload.ID == 0 || userPayload.CompanyID == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	err := Helper.StrictParseDataFromJson(r, &extensionStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	//check validation.
	boolType := ValidationCheck(extensionStruct)
	if boolType {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	tx := utility.Db.MustBegin()

	//Update data in table.
	boolValue, err := models.Extensions{}.PostExtension(extensionStruct, tx)
	if !boolValue || err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	txError := tx.Commit()
	if txError != nil {
		tx.Rollback()
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	} else {
		response.Status = "200"
		response.Message = "Extension updated successfully."
	}

	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func ValidationCheck(extensionStruct models.Extensions) bool {
	switch {
	case extensionStruct.Extension == "":
		return true
	case extensionStruct.UserId <= 0:
		return true
	case extensionStruct.Department <= 0:
		return true
	case extensionStruct.SipServer == "":
		return true
	case extensionStruct.SipUserName == "":
		return true
	case extensionStruct.SipPassword == "":
		return true
	case extensionStruct.SipPort == "":
		return true
	default:
		return false
	}
}

type Payload struct {
	CompanyID    string `json:"companyid"`
	Type         string `json:"type"`
	Secret       string `json:"secret"`
	Name         string `json:"name"`
	Extension    string `json:"extension"`
	NewExtension string `json:"new_extension,omitempty"`
	Host         string `json:"host,omitempty"`
	Number       string `json:"number,omitempty"`
}

type Message struct {
	Work    string  `json:"work"`
	Method  string  `json:"method"`
	Payload Payload `json:"payload"`
}

func PutExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	var extensionStruct models.Extensions
	// var userPayload models.Users

	isOk, userPayload := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if userPayload.ID == 0 || userPayload.CompanyID == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	err := Helper.StrictParseDataFromJson(r, &extensionStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	log.Println("extension struct: ", extensionStruct)
	//check validation.
	boolType := ValidationCheck(extensionStruct)
	if boolType {

		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	tx := utility.Db.MustBegin()

	//put data in table.
	boolValue, err := models.Extensions{}.PutExtension(extensionStruct, tx)
	if boolValue || err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	txError := tx.Commit()
	if txError != nil {
		tx.Rollback()
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	} else {
		response.Status = "200"
		response.Message = "Extension added successfully."
	}

	Helper.RenderJsonResponse(w, r, response, 200)
	return false
	// // Connect to RabbitMQ server
	// conn, err := amqp.Dial("amqp://reak:reak@localhost:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	// // Create a channel
	// ch, err := conn.Channel()
	// failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	// // Declare the queue
	// queueName := "bestsellers_product"
	// _, err = ch.QueueDeclare(
	// 	queueName, // name
	// 	true,      // durable
	// 	false,     // delete when unused
	// 	false,     // exclusive
	// 	false,     // no-wait
	// 	nil,       // arguments
	// )
	// failOnError(err, "Failed to declare a queue")
	// message := Message{
	// 	Work:   "extension", // or "sip_config"
	// 	Method: "PUT",       // or "PUT" or "DELETE"
	// 	Payload: Payload{
	// 		CompanyID: "12",
	// 		Type:      "owner",
	// 		Secret:    "pratikpassword",
	// 		Name:      "pra",
	// 		Extension: "201",
	// 		Host:      "cg.voip.ims.bsnl.in",
	// 		Number:    "1234567890",
	// 	},
	// }
	// // Publish a message to the queue

	// body, _ := json.Marshal(message)
	// err = ch.Publish(
	// 	"",        // exchange
	// 	queueName, // routing key (queue name)
	// 	false,     // mandatory
	// 	false,     // immediate
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        body,
	// 	})
	// failOnError(err, "Failed to publish a message")

	// fmt.Printf(" [x] Sent '%s'\n", body)
	// return true
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func DeleteExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	extensionId := Helper.StrToInt(r.URL.Query().Get("id"))

	if extensionId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	isOk, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "403"
		response.Message = "Unauthorized access, you are not allowed to make this request!"
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}
	// get extension which user is trying to delete
	extension, err := models.Extensions{}.GetExtensionById(int64(extensionId))
	if err != nil {
		log.Println("err fetching extension: ", err)
		response.Status = "400"
		response.Message = "Unable to fetch desired extension for deletion or it doesn't exists!"
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}
	// check if extension doesnt belong to same user and user isn't super admin -> 403
	if extension.UserId != userDetails.ID && userDetails.AccountType != "super-admin" {
		response.Status = "403"
		response.Message = "Unauthorized access, you are not allowed to make this request!"
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}
	boolType, err := models.Extensions{}.DeleteExtensionDetail(extensionId)
	if !boolType || err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			// MySQL error code 1451 indicates a foreign key constraint
			if driverErr.Number == 1451 {
				response.Message = Helper.GetSqlErrorString(err)
			}
		}
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}
	response.Status = "200"
	response.Message = "Extension deleted successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func GetExtensionData(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var para models.ExtensionCondition

	// var userPayload models.Users

	isOk, userPayload := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "403"
		response.Message = "Unauthorized access, User not found!"
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	// if userPayload.ID == 0 || userPayload.CompanyID == 0 {
	// 	response.Status = "403"
	// 	response.Message = "Unauthorized access, UserId or companyId doesn't match."
	// 	Helper.RenderJsonResponse(w, r, response, 403)
	// 	return true
	// }

	// para.AccountType = userPayload.AccountType

	para.Id, _ = Helper.StrToInt64(r.URL.Query().Get("id")) // take id for url
	// para.CompanyId, _ = Helper.StrToInt64(r.URL.Query().Get("companyId")) // take company_id for url

	// if userPayload.ID != int64(para.CompanyId) {
	// 	response.Status = "403"
	// 	response.Message = "You are not authorized for this request."
	// 	Helper.RenderJsonResponse(w, r, response, 403)
	// 	return true
	// }

	// if userPayload.ID != 0 && para.AccountType == "user" {
	// 	para.Id = int64(userPayload.ID)
	// }
	para.UserId = userPayload.ID
	parameters := models.Extensions{}.GetParaForFilterExtension(para)
	result, err := models.Extensions{}.GetExtensions(parameters)
	if err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	if len(result) == 0 {
		response.Status = "200"
		response.Message = "No result were found for this search."
	} else {
		response.Status = "200"
		response.Message = "Returns all matching  Extensions."
		response.Payload = result // Set the extension data in the response payload
	}
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}
