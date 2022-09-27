package helpers

import (
	"fmt"
	"os/exec"
)

// RunExecSetAccount execs command to set git account
func RunExecSetAccount(name string, email string) {
	execSetGitConfigName(name)
	execSetGitConfigEmail(email)
}

func execSetGitConfigName(name string) {
	cmdStr := "git config --global user.name " + name
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		PrintErrorExecutingMode()
	}
	fmt.Println("👤 " + name + " was set as username")
}

func execSetGitConfigEmail(email string) {
	cmdStr := "git config --global user.email " + email
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		PrintErrorExecutingMode()
	}
	fmt.Println("📧 " + email + " was set as email")
}

// RunCurrentAccount returns current user git account
func RunCurrentAccount() (string, string) {
	cmdName := exec.Command("/bin/sh", "-c", "git config --global user.name")
	cmdEmail := exec.Command("/bin/sh", "-c", "git config --global user.email")

	email, emailErr := cmdEmail.CombinedOutput()
	if emailErr != nil {
		PrintErrorExecutingMode()
	}
	name, nameErr := cmdName.CombinedOutput()
	if nameErr != nil {
		PrintErrorExecutingMode()
	}

	return string(email), string(name)
}
