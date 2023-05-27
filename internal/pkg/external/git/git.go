package git

import (
	"fmt"
	"go-gituser/internal/pkg/logger"
	"os/exec"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) execCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
		return "", err
	}
	return string(output), nil
}

func (s *Service) ReadConfig(key string) (string, error) {
	cmdStr := fmt.Sprintf("account config --global %s", key)
	value, err := s.execCommand(cmdStr)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *Service) SetConfig(key, value string) {
	cmdStr := fmt.Sprintf("account config --global %s %s", key, value)
	_, err := s.execCommand(cmdStr)
	if err != nil {
		logger.PrintErrorExecutingMode()
		return
	}
}
