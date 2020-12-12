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
	}

	if len(os.Args) > 2 {
		helpers.PrintErrorInvalidArguments()
	}

	help := flag.Bool("help", false, "Print informations about the program")

	flag.Parse()

	if *help {
		helpers.PrintHelp()
	}

	argValue := strings.ToUpper(os.Args[1])
	gitAccount := models.Account{}
	data, _ := helpers.GetDataFromJSON("data/config.json")
	_ = json.Unmarshal(data, &gitAccount)

	switch argValue {
	case "WORK":
		helpers.RunModeConfig(gitAccount.WorkUsername, gitAccount.WorkEmail)
	case "SCHOOL":
		helpers.RunModeConfig(gitAccount.SchoolUsername, gitAccount.SchoolEmail)
	case "PERSONAL":
		helpers.RunModeConfig(gitAccount.PersonalUsername, gitAccount.PersonaEmail)
	}

}
