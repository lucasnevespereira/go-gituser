package models

import "sort"

type Accounts struct {
	Personal Account            `json:"personal"`
	Work     Account            `json:"work"`
	School   Account            `json:"school"`
	Custom   map[string]Account `json:"custom,omitempty"`
}

type Account struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	SigningKeyID string `json:"signingkeyid"`
	SSHKeyPath   string `json:"sshkeypath"`
}

func (a *Accounts) Get(mode string) (Account, bool) {
	switch mode {
	case PersonalMode:
		return a.Personal, a.Personal.Username != ""
	case WorkMode:
		return a.Work, a.Work.Username != ""
	case SchoolMode:
		return a.School, a.School.Username != ""
	default:
		account, ok := a.Custom[mode]
		return account, ok && account.Username != ""
	}
}

func (a *Accounts) Set(mode string, account Account) {
	switch mode {
	case PersonalMode:
		a.Personal = account
	case WorkMode:
		a.Work = account
	case SchoolMode:
		a.School = account
	default:
		if a.Custom == nil {
			a.Custom = make(map[string]Account)
		}
		a.Custom[mode] = account
	}
}

// ForEachConfigured visits saved accounts in a stable order.
func (a *Accounts) ForEachConfigured(visit func(mode string, account Account) bool) {
	for _, mode := range []string{PersonalMode, SchoolMode, WorkMode} {
		if account, ok := a.Get(mode); ok && !visit(mode, account) {
			return
		}
	}
	modes := make([]string, 0, len(a.Custom))
	for mode := range a.Custom {
		modes = append(modes, mode)
	}
	sort.Strings(modes)
	for _, mode := range modes {
		if account, ok := a.Get(mode); ok && !visit(mode, account) {
			return
		}
	}
}
