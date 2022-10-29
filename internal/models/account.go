package models

type Accounts struct {
	PersonalUsername string `json:"personalUsername"`
	PersonalEmail    string `json:"personalEmail"`
	WorkUsername     string `json:"workUsername"`
	WorkEmail        string `json:"workEmail"`
	SchoolUsername   string `json:"schoolUsername"`
	SchoolEmail      string `json:"schoolEmail"`
}
