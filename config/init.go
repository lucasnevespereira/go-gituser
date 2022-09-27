package config

import (
	"encoding/json"
	"fmt"
	"go-gituser/helpers"
	"go-gituser/models"
	"io/ioutil"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

var (
	pUsername            string
	pEmail               string
	wUsername            string
	wEmail               string
	sUsername            string
	sEmail               string
	shouldConfigureAgain string
)

// InitSetData starts git account data setup*
func InitSetData() {

	for {
		prompt := promptui.Select{
			Label: "Please choose an account to configure",
			Items: []string{
				helpers.WorkSelectLabel,
				helpers.SchoolSelectLabel,
				helpers.PersonalSelectLabel,
				helpers.CancelSelectLabel,
			},
		}

		_, choice, err := prompt.Run()
		if err != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		switch choice {
		case helpers.WorkSelectLabel:
			getUserAccount(helpers.WorkMode)
			helpers.PrintRemeberToActiveMode(helpers.WorkMode)
		case helpers.SchoolSelectLabel:
			getUserAccount(helpers.SchoolMode)
			helpers.PrintRemeberToActiveMode(helpers.PersonalMode)
		case helpers.PersonalSelectLabel:
			getUserAccount(helpers.PersonalMode)
			helpers.PrintRemeberToActiveMode(helpers.PersonalMode)
		case helpers.CancelSelectLabel:
			os.Exit(1)
		}

		fmt.Println("Would you like to configure another account ? (y/n)")
		_, err = fmt.Scanln(&shouldConfigureAgain)
		if err != nil {
			helpers.PrintErrorReadingInput()
		}

		shouldConfigureAgain = strings.ToUpper(strings.TrimSpace(shouldConfigureAgain))

		if shouldConfigureAgain != "Y" {
			fmt.Println("Okay. Bye there!")
			break
		}
	}

	checkForEmptyAccountData()
	writeAccountData()

}

func getUserAccount(mode string) {
	switch mode {
	case helpers.WorkMode:
		fmt.Println("What is your work username ?")
		_, errUsername := fmt.Scanln(&wUsername)
		if errUsername != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your work email ?")
		_, errEmail := fmt.Scanln(&wEmail)
		if errEmail != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

	case helpers.SchoolMode:
		fmt.Println("What is your school username ?")
		_, errUsername := fmt.Scanln(&sUsername)
		if errUsername != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your school email ?")
		_, errEmail := fmt.Scanln(&sEmail)
		if errEmail != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}
	case helpers.PersonalMode:
		fmt.Println("What is your personal username ?")
		_, errUsername := fmt.Scanln(&pUsername)
		if errUsername != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your personal email ?")
		_, errEmail := fmt.Scanln(&pEmail)
		if errEmail != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

	case helpers.CancelSelectLabel:
		os.Exit(1)
	}

}

// WriteAccountData writes data to json file
func writeAccountData() {
	configFile, err := helpers.GetConfigFilePath()
	if err != nil {
		helpers.PrintError(err)
	}

	data := models.Account{
		PersonalUsername: pUsername,
		PersonalEmail:    pEmail,
		WorkUsername:     wUsername,
		WorkEmail:        wEmail,
		SchoolUsername:   sUsername,
		SchoolEmail:      sEmail,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	ioutil.WriteFile(configFile, file, 666)

}

// checkForEmptyAccountData checks if there is no overrides with empty accounts.
func checkForEmptyAccountData() {
	currAccount := models.Account{}

	configFilePath, err := helpers.GetConfigFilePath()
	if err != nil {
		helpers.PrintError(err)
	}

	data, _ := helpers.GetDataFromJSON(configFilePath)
	_ = json.Unmarshal(data, &currAccount)

	if pEmail == "" || pUsername == "" {
		pEmail = currAccount.PersonalEmail
		pUsername = currAccount.PersonalUsername
	}

	if wEmail == "" || wUsername == "" {
		wEmail = currAccount.WorkEmail
		wUsername = currAccount.WorkUsername
	}

	if sEmail == "" || sUsername == "" {
		sEmail = currAccount.SchoolEmail
		sUsername = currAccount.SchoolUsername
	}

}
