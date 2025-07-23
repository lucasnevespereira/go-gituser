package git

import (
	"fmt"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"os/exec"
	"runtime"
	"strings"
)

type IConnector interface {
	ReadConfig() *models.Account
	SetConfig(account *models.Account)
}

type Connector struct{}

func NewGitConnector() IConnector {
	return &Connector{}
}

func (c *Connector) ReadConfig() *models.Account {
	var cmdName, cmdEmail, cmdSigningKeyID *exec.Cmd

	if runtime.GOOS == "windows" {
		cmdName = exec.Command("cmd", "/C", "git config --global user.name")
		cmdEmail = exec.Command("cmd", "/C", "git config --global user.email")
		cmdSigningKeyID = exec.Command("cmd", "/C", "git config --global user.signingkey")
	} else {
		cmdName = exec.Command("/bin/sh", "-c", "git config --global user.name")
		cmdEmail = exec.Command("/bin/sh", "-c", "git config --global user.email")
		cmdSigningKeyID = exec.Command("/bin/sh", "-c", "git config --global user.signingkey")
	}

	emailBytes, emailBytesErr := cmdEmail.CombinedOutput()
	if emailBytesErr != nil {
		logger.PrintErrorExecutingMode()
	}
	nameBytes, nameBytesErr := cmdName.CombinedOutput()
	if nameBytesErr != nil {
		logger.PrintErrorExecutingMode()
	}

	// we ignore gpg signing key err, since it is optional
	cmdSigningKeyIDBytes, _ := cmdSigningKeyID.CombinedOutput()

	return &models.Account{
		Username:     strings.TrimSpace(string(nameBytes)),
		Email:        strings.TrimSpace(string(emailBytes)),
		SigningKeyID: strings.TrimSpace(string(cmdSigningKeyIDBytes)),
	}
}

func (c *Connector) SetConfig(account *models.Account) {
	c.setConfigName(account.Username)
	c.setConfigEmail(account.Email)
	c.setConfigSigningKey(account.SigningKeyID)

}

func (c *Connector) setConfigName(name string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git config --global user.name "+name)
	} else {
		cmdStr := "git config --global user.name " + name
		cmd = exec.Command("/bin/sh", "-c", cmdStr)
	}
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("👤 " + name + " was set as username")
}

func (c *Connector) setConfigEmail(email string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git config --global user.email "+email)
	} else {
		cmdStr := "git config --global user.email " + email
		cmd = exec.Command("/bin/sh", "-c", cmdStr)
	}
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("📧 " + email + " was set as email")
}

func (c *Connector) setConfigSigningKey(SigningKeyID string) {
	// If SigningKeyID is empty, we're not using GPG signing
	if SigningKeyID == "" {
		c.setCommitSign(false)
		return
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git config --global user.signingkey "+SigningKeyID)
	} else {
		cmdStr := "git config --global user.signingkey " + SigningKeyID
		cmd = exec.Command("/bin/sh", "-c", cmdStr)
	}
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("🔑 " + SigningKeyID + " was set as gpg signing key")

	// Enable commit signing when setting a signing key
	c.setCommitSign(true)
}

func (c *Connector) setCommitSign(enable bool) {
	value := "false"
	if enable {
		value = "true"
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git config --global commit.gpgsign "+value)
	} else {
		cmdStr := "git config --global commit.gpgsign " + value
		cmd = exec.Command("/bin/sh", "-c", cmdStr)
	}
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}

	if enable {
		fmt.Println("✅ Commit signing enabled")
	} else {
		fmt.Println("❌ Commit signing disabled")
	}
}
