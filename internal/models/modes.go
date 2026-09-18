package models

import (
	"fmt"
	"regexp"
)

const (
	WorkMode     = "work"
	SchoolMode   = "school"
	PersonalMode = "personal"
)

var customModeName = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

var reservedModes = map[string]bool{
	WorkMode: true, SchoolMode: true, PersonalMode: true,
	"setup": true, "now": true, "info": true, "ssh": true,
	"help": true, "manual": true, "quickstart": true, "version": true,
	"completion": true,
}

func ValidateCustomMode(mode string) error {
	if !customModeName.MatchString(mode) {
		return fmt.Errorf("mode name must start with a lowercase letter and contain only lowercase letters, numbers, or hyphens")
	}
	if reservedModes[mode] {
		return fmt.Errorf("%q is already a gituser command", mode)
	}
	return nil
}
