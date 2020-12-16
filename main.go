package main

import (
	"encoding/json"
	"flag"
	"go-gituser/helpers"
	"go-gituser/models"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		helpers.PrintHelp()
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		helpers.PrintErrorInvalidArguments()
		os.Exit(1)
	}

	argValue := strings.ToUpper(os.Args[1])

	gitAccount := models.Account{}

	configFilePath := os.Getenv("PATH_TO_GITUSER_CONFIG")
	if configFilePath == "" {
		configFilePath = "data/config.json"
	}

	data, _ := helpers.GetDataFromJSON(configFilePath)
	_ = json.Unmarshal(data, &gitAccount)

	// Flags
	help := flag.Bool("help", false, "Print informations about the program")
	info := flag.Bool("info", false, "Print informations about the accounts")

	flag.Parse()

	if *help {
		helpers.PrintHelp()
		os.Exit(1)
	}

	if *info {
		helpers.ReadAccountsData(gitAccount)
		os.Exit(1)
	}

	switch argValue {
	case "WORK":
		if gitAccount.WorkUsername == "" {
			helpers.PrintWarningReadingAccount("work")
			os.Exit(1)
		}
		helpers.RunModeConfig(gitAccount.WorkUsername, gitAccount.WorkEmail)
	case "SCHOOL":
		if gitAccount.SchoolUsername == "" {
			helpers.PrintWarningReadingAccount("school")
			os.Exit(1)
		}
		helpers.RunModeConfig(gitAccount.SchoolUsername, gitAccount.SchoolEmail)
	case "PERSONAL":
		if gitAccount.PersonalUsername == "" {
			helpers.PrintWarningReadingAccount("personal")
			os.Exit(1)
		}
		helpers.RunModeConfig(gitAccount.PersonalUsername, gitAccount.PersonalEmail)
	}

}
