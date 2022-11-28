package models

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseOK struct {
	Message string `json:"message"`
}
