package out

type StandardResponse struct {
	Header  HeaderResponse `json:"header"`
	Payload Payload        `json:"payload"`
}

type HeaderResponse struct {
	RequestID string `json:"request_id"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

type Payload struct {
	Status StatusPayload `json:"status"`
	Data   interface{}   `json:"data"`
	Other  interface{}   `json:"other"`
}

type StatusPayload struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}
