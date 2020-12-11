package helpers

import (
	"fmt"
	"os/exec"
)

// RunModeConfig sets git mode configuration
func RunModeConfig(name string, email string) {
	execGitConfigName(name)
	execGitConfigEmail(email)
}

func execGitConfigName(name string) {
	cmdStr := "git config --global user.name " + name
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		PrintErrorExecutingMode()
	}
	fmt.Println(name + " was set with success 😀")
}

func execGitConfigEmail(email string) {
	cmdStr := "git config --global user.email " + email
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		PrintErrorExecutingMode()
	}
	fmt.Println(email + " was set with success 😀")
}
