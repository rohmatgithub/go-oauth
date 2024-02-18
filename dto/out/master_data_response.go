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

type CompanyResponse struct {
	Header  Header `json:"header"`
	Payload struct {
		Status Status `json:"status"`
		Data   struct {
			ID               int64     `json:"id"`
			CompanyProfileID int64     `json:"company_profile_id"`
			Code             string    `json:"code"`
			Name             string    `json:"name"`
			NPWP             string    `json:"npwp"`
			Address1         string    `json:"address_1"`
			CreatedAt        time.Time `json:"created_at"`
			UpdatedAt        time.Time `json:"updated_at"`
		} `json:"data"`
		Other interface{} `json:"other"`
	} `json:"payload"`
}
