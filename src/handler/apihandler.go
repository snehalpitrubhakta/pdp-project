package handler

import (
	"data"
	"encoding/json"
)

const (
	REGISTER = "register"
	LOGIN    = "login"
)

func HandleApiRequest(request *data.ApiRequest) []byte {
	var response *data.ApiResponse = nil
	switch *request.Api {
	case LOGIN:
		response = &data.ApiResponse{Status: 200, Data: "Login Successful", Error: ""}
	default:
		response = &data.ApiResponse{Status: 200, Data: nil, Error: "Invalid Api"}
	}

	resp, err := json.Marshal(response)
	if nil != err {
		return nil
	}
	return resp
}
