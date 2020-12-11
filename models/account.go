package models

// Account is the account structure
type Account struct {
	username string
	email    string
}

// SetWorkMode sets work git username and email
func (a *Account) SetWorkMode() {
	a.username = "lucas.pereiraneves"
	a.email = "lucas.pereiraneves@grtgaz.com"
}

// SetSchoolMode sets school git username and email
func (a *Account) SetSchoolMode() {
	a.username = "lucas.nevespereira"
	a.email = "lucas.nevespereira@eemi.com"
}

// SetPersonalMode sets personal git username and email
func (a *Account) SetPersonalMode() {
	a.username = "lucasnevespereira"
	a.email = "pereiraneveslucas@gmail.com"
}

// GetAccountUsername is..
func (a *Account) GetAccountUsername() string {
	return a.username
}

// GetAccountEmail is..
func (a *Account) GetAccountEmail() string {
	return a.email
}
