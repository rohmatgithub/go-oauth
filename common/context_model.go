package common

type ContextModel struct {
	LoggerModel          LoggerModel
	AuthAccessTokenModel AuthAccessTokenModel
	PermissionHave       string
}

type AuthAccessTokenModel struct {
	ResourceUserID int64 `json:"rid"`
	CompanyID      int64 `json:"cpid"`
	// Scope          string `json:"scp"`
	// ClientID       string `json:"cid"`
	Locale string `json:"lang"`
}
