package git

import (
	"fmt"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"os/exec"
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
	cmdName := exec.Command("/bin/sh", "-c", "git config --global user.name")
	cmdEmail := exec.Command("/bin/sh", "-c", "git config --global user.email")

	emailBytes, emailBytesErr := cmdEmail.CombinedOutput()
	if emailBytesErr != nil {
		logger.PrintErrorExecutingMode()
	}
	nameBytes, nameBytesErr := cmdName.CombinedOutput()
	if nameBytesErr != nil {
		logger.PrintErrorExecutingMode()
	}

	return &models.Account{
		Username: string(nameBytes),
		Email:    string(emailBytes),
	}
}

func (c *Connector) SetConfig(account *models.Account) {
	c.setConfigName(account.Username)
	c.setConfigEmail(account.Email)
}

func (c *Connector) setConfigName(name string) {
	cmdStr := "git config --global user.name " + name
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("👤 " + name + " was set as username")
}

func (c *Connector) setConfigEmail(email string) {
	cmdStr := "git config --global user.email " + email
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("📧 " + email + " was set as email")
}
