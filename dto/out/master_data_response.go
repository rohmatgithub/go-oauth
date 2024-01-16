package out

import "time"

type Header struct {
	RequestID string    `json:"request_id"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type Status struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

type CompanyBranchResponse struct {
	Header  Header `json:"header"`
	Payload struct {
		Status Status `json:"status"`
		Data   []struct {
			ID        int       `json:"id"`
			Code      string    `json:"code"`
			Name      string    `json:"name"`
			Address1  string    `json:"address_1"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"data"`
		Other interface{} `json:"other"`
	} `json:"payload"`
}
