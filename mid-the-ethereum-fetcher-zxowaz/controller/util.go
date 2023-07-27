package controller

import (
	"encoding/json"
	"mid-the-ethereum-fetcher-zxowaz/logger"
	"net/http"
)

func writeResponse(w http.ResponseWriter, err error, result interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		errorMessage := err.Error()
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}
