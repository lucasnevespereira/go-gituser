package main

import (
	"encoding/json"
	"flag"
	"go-gituser/config"
	"go-gituser/helpers"
	"go-gituser/models"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		helpers.PrintManual()
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		helpers.PrintErrorInvalidArguments()
		os.Exit(1)
	}

	argValue := strings.ToLower(os.Args[1])

	gitAccount := models.Account{}

	configFilePath, err := helpers.GetConfigFilePath()
	if err != nil {
		helpers.PrintError(err)
	}

	data, err := helpers.GetDataFromJSON(configFilePath)
	if err != nil {
		helpers.PrintError(err)
	}

	err = json.Unmarshal(data, &gitAccount)
	if err != nil {
		helpers.PrintError(err)
	}

	// Flags
	manual := flag.Bool("manual", false, "Print informations about the program")
	info := flag.Bool("info", false, "Print informations about the accounts")
	now := flag.Bool("now", false, "Print current git user account")

	flag.Parse()

	if *manual {
		helpers.PrintManual()
		os.Exit(1)
	}

	if *info {
		helpers.ReadAccountsData(gitAccount)
		os.Exit(1)
	}

	if *now {
		currEmail, currName := helpers.RunCurrentAccount()
		currName = strings.TrimSuffix(currName, "\n")
		currEmail = strings.TrimSuffix(currEmail, "\n")

		if gitAccount.PersonalUsername == (currName) && gitAccount.PersonalEmail == (currEmail) {
			helpers.ReadCurrentAccountData(currName, currEmail, "personal")
		}

		if gitAccount.SchoolUsername == (currName) && gitAccount.SchoolEmail == (currEmail) {
			helpers.ReadCurrentAccountData(currName, currEmail, "school")
		}

		if gitAccount.WorkUsername == (currName) && gitAccount.WorkEmail == (currEmail) {
			helpers.ReadCurrentAccountData(currName, currEmail, "work")
		}

		os.Exit(1)
	}

	switch argValue {
	case helpers.WorkMode:
		if gitAccount.WorkUsername == "" {
			helpers.PrintWarningReadingAccount(helpers.WorkMode)
			os.Exit(1)
		}
		helpers.RunExecSetAccount(gitAccount.WorkUsername, gitAccount.WorkEmail)
	case helpers.SchoolMode:
		if gitAccount.SchoolUsername == "" {
			helpers.PrintWarningReadingAccount(helpers.SchoolMode)
			os.Exit(1)
		}
		helpers.RunExecSetAccount(gitAccount.SchoolUsername, gitAccount.SchoolEmail)
	case helpers.PersonalMode:
		if gitAccount.PersonalUsername == "" {
			helpers.PrintWarningReadingAccount(helpers.PersonalMode)
			os.Exit(1)
		}
		helpers.RunExecSetAccount(gitAccount.PersonalUsername, gitAccount.PersonalEmail)
	case helpers.ConfigMode:
		config.InitSetData()
	}

}
