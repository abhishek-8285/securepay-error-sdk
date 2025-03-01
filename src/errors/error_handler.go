package errors

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ErrorCode struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"httpStatus"`
}

var errorCodes map[string]ErrorCode

func init() {
	data, err := ioutil.ReadFile("src/errors/errorCodes.json")
	if err != nil {
		log.Fatal("Failed to load error mappings: ", err)
	}
	json.Unmarshal(data, &errorCodes)
}

func GetError(errorKey string) ErrorCode {
	if err, exists := errorCodes[errorKey]; exists {
		return err
	}
	return errorCodes["SYSTEM_INTERNAL_ERROR"]
}

func ErrorResponse(w http.ResponseWriter, errorKey string) {
	errorDetail := GetError(errorKey)
	w.WriteHeader(errorDetail.HttpStatus)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "FAILED",
		"error":  errorDetail,
	})
}
