package resp

import "encoding/json"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

func SuccessResponse(data interface{}) *Response {
	bytes, err := json.Marshal(data)
	if err != nil {
		return ErrorResponse("INTERNAL_ERROR", "failed to serialize response data")
	}

	return &Response{
		Success: true,
		Data:    bytes,
	}
}

func ErrorResponse(code, message string) *Response {
	return &Response{Success: false, Error: &Error{Code: code, Message: message}}
}
