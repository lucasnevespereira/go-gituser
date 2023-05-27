package account

import (
	"go-gituser/internal/pkg/external/git"
	logger2 "go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/models"
)

type Service struct {
	git *git.Service
}

func NewAccountService(gitService *git.Service) *Service {
	return &Service{git: gitService}
}

func (s *Service) CurrentAccount() (string, string) {
	name, nameErr := s.git.ReadConfig("user.name")
	if nameErr != nil {
		logger2.PrintError(nameErr)
	}
	email, emailErr := s.git.ReadConfig("user.email")
	if emailErr != nil {
		logger2.PrintError(emailErr)
	}

	return email, name
}

func (s *Service) SetAccount(name, email string) {
	s.git.SetConfig("user.name", name)
	s.git.SetConfig("user.email", email)

	logger2.ReadAccountSetted("👤", "user.name", name)
	logger2.ReadAccountSetted("📧", "user.email", email)
}

func (s *Service) UsernameIsUnsaved(savedAccounts models.Accounts, name string) bool {
	return savedAccounts.PersonalUsername != name ||
		savedAccounts.WorkUsername != name || savedAccounts.SchoolUsername != name
}

func (s *Service) EmailIsUnsaved(savedAccounts models.Accounts, email string) bool {
	return savedAccounts.PersonalEmail != email ||
		savedAccounts.WorkEmail != email || savedAccounts.SchoolEmail != email
}
