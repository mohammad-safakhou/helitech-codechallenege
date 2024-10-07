package utils

type StandardHttpResponse struct {
	error
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
