package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAccountsPreserveBuiltInModesAndListCustomModes(t *testing.T) {
	var accounts Accounts
	if err := json.Unmarshal([]byte(`{"personal":{"username":"alice","email":"alice@example.com"},"work":{"username":"bob","email":"bob@example.com"}}`), &accounts); err != nil {
		t.Fatal(err)
	}
	accounts.Set("freelance", Account{Username: "carol", Email: "carol@example.com"})
	accounts.Set("open-source", Account{Username: "dave", Email: "dave@example.com"})

	var modes []string
	accounts.ForEachConfigured(func(mode string, _ Account) bool {
		modes = append(modes, mode)
		return true
	})
	if want := []string{"personal", "work", "freelance", "open-source"}; !reflect.DeepEqual(modes, want) {
		t.Fatalf("configured modes = %v, want %v", modes, want)
	}
	if account, ok := accounts.Get("freelance"); !ok || account.Email != "carol@example.com" {
		t.Fatalf("custom account = %+v, found = %v", account, ok)
	}
	if _, ok := accounts.Get("missing"); ok {
		t.Fatal("missing mode was found")
	}

	data, err := json.Marshal(&accounts)
	if err != nil {
		t.Fatal(err)
	}
	var restored Accounts
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(accounts, restored) {
		t.Fatalf("round trip changed accounts: %+v", restored)
	}
}

func TestValidateCustomMode(t *testing.T) {
	for _, mode := range []string{"freelance", "open-source", "client2"} {
		if err := ValidateCustomMode(mode); err != nil {
			t.Errorf("%q should be valid: %v", mode, err)
		}
	}
	for _, mode := range []string{"", "Work", "2client", "open_source", "open source", "setup", "personal", "help"} {
		if err := ValidateCustomMode(mode); err == nil {
			t.Errorf("%q should be rejected", mode)
		}
	}
}
