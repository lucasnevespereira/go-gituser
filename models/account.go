package models

// Account is the account structure
type Account struct {
	PersonalUsername string `json:"personalUsername"`
	PersonalEmail    string `json:"personalEmail"`
	WorkUsername     string `json:"workUsername"`
	WorkEmail        string `json:"workEmail"`
	SchoolUsername   string `json:"schoolUsername"`
	SchoolEmail      string `json:"schoolEmail"`
}
