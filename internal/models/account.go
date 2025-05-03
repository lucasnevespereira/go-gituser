package models

type Accounts struct {
	Personal Account `json:"personal"`
	Work     Account `json:"work"`
	School   Account `json:"school"`
}

type Account struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	SigningKeyID    string `json:"signingkeyid"`
}
