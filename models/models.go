package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"reakgo/utility"
	"strings"
)

var Utility utility.Helper

var (
	// ErrCode is a config or an internal error
	ErrCode = errors.New("Case statement in code is not correct.")
	// ErrNoResult is a not results error
	ErrNoResult = errors.New("Result not found.")
	// ErrUnavailable is a database not available error
	ErrUnavailable = errors.New("Database is unavailable.")
	// ErrUnauthorized is a permissions violation
	ErrUnauthorized = errors.New("User does not have permission to perform this operation.")
)

// standardizeErrors returns the same error regardless of the database used
func standardizeError(err error) error {
	if err == sql.ErrNoRows {
		return ErrNoResult
	}

	return err
}

func GenerateCache() {
	allRows, err := Authentication{}.GetAllAuthRecords()
	if err != nil {
		log.Println(err)
	} else {
		for _, val := range allRows {
			jsonData, err := json.Marshal(val)
			if err != nil {
				log.Println("Error encoding JSON:", err)
				break
			}
			utility.Cache.Set(val.Token, jsonData)
		}
	}
}

func VerifyToken(r *http.Request, w http.ResponseWriter) error {
	authToken := r.Header.Get("Authorization")
	userToken := strings.Split(authToken, " ")
	if len(userToken) != 2 || strings.ToLower(userToken[0]) != "bearer" {
		return fmt.Errorf("invalid authorization header format")
	}

	if entry, err := utility.Cache.Get(userToken[1]); err == nil {

		// Set JSON Payload to the header so the users can use the same
		r.Header.Add("tokenPayload", string(entry))
		return err
	} else {
		// Pull Record from DB and add to Cache

		// PS : Adding DB failsafe opens up a DDoS security issue that people can keep trying with random tokens
		// and crash the server easily by blocking DB pool connections

		data, err := Authentication.GetAuthenticationByToken(Authentication{}, userToken[1])
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(data)
		if err == nil {
			// Rehydrate if we got the JSON conversion done
			// Fails would be rare, but if it happens kind of defeat the purpose as JSON unmarshal would also crash
			utility.Cache.Set(data.Token, jsonData)
			r.Header.Add("tokenPayload", string(jsonData))
			// w.Header().Add("tokenPayload", string(jsonData))
			return err
		} else {
			return err
		}
	}
}

type FrontendUserStruct struct {
	Id               int64  `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	Email            string `json:"email" db:"email"`
	AccountType      string `json:"accountType" db:"accountType"`
	Dob              string `json:"dob" db:"dob"`
	IsWizardComplete string `json:"iswizardcomplete" db:"iswizardcomplete"`
	CompanyId        int64  `json:"companyId" db:"companyId"`
}

func getNonConfDataForFrontEnd(entry []byte) ([]byte, error) {
	var data Authentication
	var dataToSend FrontendUserStruct
	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(entry, &data); err != nil {
		log.Println("err in unmarshal: ", err)
		return nil, err
	}
	// setting data to be sent in front end
	dataToSend.Name = data.Name
	dataToSend.Email = data.Email
	dataToSend.AccountType = data.AccountType
	dataToSend.Dob = data.DOB
	dataToSend.IsWizardComplete = data.IsWizardComplete
	dataToSend.CompanyId = data.CompanyID
	dataToSend.Id = data.ID

	jsonResponse, err := json.Marshal(dataToSend)
	if err != nil {
		// Handle marshal error
		log.Println("err in marshal: ", err)
		return nil, err
	}
	return jsonResponse, nil
}
func VerifyTokenFrontend(r *http.Request, w http.ResponseWriter) error {
	authToken := r.Header.Get("Authorization")
	userToken := strings.Split(authToken, " ")
	if len(userToken) != 2 || strings.ToLower(userToken[0]) != "bearer" {
		return fmt.Errorf("invalid authorization header format")
	}

	if entry, err := utility.Cache.Get(userToken[1]); err == nil {
		jsonResponse, err := getNonConfDataForFrontEnd(entry)
		if err != nil {
			log.Println("err in marshal during tokencheck for frontend: ", err)
			return err
		}
		// Set JSON Payload to the header so the users can use the same
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("tokenPayload", string(jsonResponse))
		w.Header().Set("Access-Control-Expose-Headers", "tokenPayload")
		// r.Header.Add("tokenPayload", string(entry))
		return err
	} else {
		// Pull Record from DB and add to Cache

		// PS : Adding DB failsafe opens up a DDoS security issue that people can keep trying with random tokens
		// and crash the server easily by blocking DB pool connections

		data, err := Authentication.GetAuthenticationByToken(Authentication{}, userToken[1])
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(data)
		if err == nil {
			// Rehydrate if we got the JSON conversion done
			// Fails would be rare, but if it happens kind of defeat the purpose as JSON unmarshall would also crash
			utility.Cache.Set(data.Token, jsonData)
			jsonResponse, err := getNonConfDataForFrontEnd(jsonData)
			if err != nil {
				log.Println("err in marshal during tokencheck for frontend: ", err)
				return err
			}
			// r.Header.Add("tokenPayload", string(jsonData))
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("tokenPayload", string(jsonResponse))
			w.Header().Set("Access-Control-Expose-Headers", "tokenPayload")
			return err
		} else {
			return err
		}
	}
}

func JsonStringToAuthentication(jsonStr string) (*Authentication, error) {
	var auth Authentication

	// Unmarshal the JSON string into the Authentication struct
	if err := json.Unmarshal([]byte(jsonStr), &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}
