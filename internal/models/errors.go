package models

import "errors"

var (
	ErrNoAccountFound = errors.New("no account found")
	ErrExecutingMode  = errors.New("could not execute this mode")
	ErrReadingInput   = errors.New("couldn't read input")
	ErrSetupAccounts  = errors.New("could not setup accounts")
)
