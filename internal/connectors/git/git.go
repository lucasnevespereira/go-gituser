package git

import (
	"fmt"
	"go-gituser/internal/logger"
	"os/exec"
)

type IConnector interface {
	ReadConfig() (name string, email string)
	SetConfig(name string, email string)
}

type Connector struct{}

func NewGitConnector() IConnector {
	return &Connector{}
}

func (c *Connector) ReadConfig() (name, email string) {
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

	return string(emailBytes), string(nameBytes)
}

func (c *Connector) SetConfig(name, email string) {
	c.setConfigName(name)
	c.setConfigEmail(email)
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
