package main

import (
	"data"
	"encoding/json"
	"fmt"
	"handler"
	"logger"
	"net/http"
)

type ApiRequest struct {
	Api  *string
	Data *string
}

func handleHttpRequest(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	apiRequest := &data.ApiRequest{}
	err := json.Unmarshal([]byte(query), apiRequest)
	if nil != err {
		errMessage := fmt.Sprintf(err.Error())
		w.Write([]byte(errMessage))
	}
	response := handler.HandleApiRequest(apiRequest)
	w.Write([]byte(response))
}

func main() {
	loggerConfig := &logger.LoggerConfig{}
	logger.Init(loggerConfig)
	http.HandleFunc("/apiserver", handleHttpRequest)
	http.ListenAndServe(":3030", nil)
	logger.Debug("API Server started successfully")
}
