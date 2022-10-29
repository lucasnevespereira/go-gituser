package git

import (
	"fmt"
	"go-gituser/utils/logger"
	"os/exec"
)

func CurrentAccount() (string, string) {
	cmdName := exec.Command("/bin/sh", "-c", "git config --global user.name")
	cmdEmail := exec.Command("/bin/sh", "-c", "git config --global user.email")

	email, emailErr := cmdEmail.CombinedOutput()
	if emailErr != nil {
		logger.PrintErrorExecutingMode()
	}
	name, nameErr := cmdName.CombinedOutput()
	if nameErr != nil {
		logger.PrintErrorExecutingMode()
	}

	return string(email), string(name)
}

func SetAccount(name string, email string) {
	setConfigName(name)
	setConfigEmail(email)
}

func setConfigName(name string) {
	cmdStr := "git config --global user.name " + name
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("👤 " + name + " was set as username")
}

func setConfigEmail(email string) {
	cmdStr := "git config --global user.email " + email
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("📧 " + email + " was set as email")
}
