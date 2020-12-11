package main

import (
	"flag"
	"go-gituser/helpers"
	"go-gituser/models"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		helpers.PrintErrorInvalidArguments()
	}

	help := flag.Bool("help", false, "Print usage information")
	flag.Parse()
	if *help {
		helpers.PrintHelp()
	}

	argValue := strings.ToUpper(os.Args[1])
	gitAccount := models.Account{}

	switch argValue {
	case "WORK":
		gitAccount.SetWorkMode()
		gitUsername := gitAccount.GetAccountUsername()
		gitEmail := gitAccount.GetAccountEmail()
		helpers.RunModeConfig(gitUsername, gitEmail)
	case "SCHOOL":
		gitAccount.SetSchoolMode()
		gitUsername := gitAccount.GetAccountUsername()
		gitEmail := gitAccount.GetAccountEmail()
		helpers.RunModeConfig(gitUsername, gitEmail)
	case "PERSONAL":
		gitAccount.SetPersonalMode()
		gitUsername := gitAccount.GetAccountUsername()
		gitEmail := gitAccount.GetAccountEmail()
		helpers.RunModeConfig(gitUsername, gitEmail)
	}

}
