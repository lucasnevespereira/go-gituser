package config

import (
	"encoding/json"
	"fmt"
	"go-gituser/helpers"
	"go-gituser/models"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	work     = "💻 Work Account"
	school   = "📚 School Account"
	personal = "🏠 Personal Account"
)

var (
	pUsername            string
	pEmail               string
	wUsername            string
	wEmail               string
	sUsername            string
	sEmail               string
	shouldConfigureAgain string
)

// InitSetData starts git account data setup*
func InitSetData() {

	for {
		prompt := promptui.Select{
			Label: "Please choose an account to configure",
			Items: []string{work, school, personal},
		}

		_, choice, err := prompt.Run()
		if err != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		switch choice {
		case work:
			getUserAccount("work")
		case school:
			getUserAccount("school")
		case personal:
			getUserAccount("personal")
		}

		fmt.Println("Would you like to configure another account ? (y/n)")
		_, err = fmt.Scanln(&shouldConfigureAgain)
		if err != nil {
			helpers.PrintErrorReadingInput()
		}

		shouldConfigureAgain = strings.ToUpper(strings.TrimSpace(shouldConfigureAgain))

		if shouldConfigureAgain != "Y" {
			fmt.Println("Okay. Bye there!")
			break
		}
	}

	writeAccountData()

}

func getUserAccount(mode string) {
	switch mode {
	case "work":
		fmt.Println("What is your work username ?")
		_, errUsername := fmt.Scanln(&wUsername)
		if errUsername != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your work email ?")
		_, errEmail := fmt.Scanln(&wEmail)
		if errEmail != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}
	case "school":
		fmt.Println("What is your school username ?")
		_, errUsername := fmt.Scanln(&sUsername)
		if errUsername != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your school email ?")
		_, errEmail := fmt.Scanln(&sEmail)
		if errEmail != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}
	case "personal":
		fmt.Println("What is your personal username ?")
		_, errUsername := fmt.Scanln(&pUsername)
		if errUsername != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your personal email ?")
		_, errEmail := fmt.Scanln(&pEmail)
		if errEmail != nil {
			helpers.PrintErrorReadingInput()
			os.Exit(1)
		}
	}

}

// WriteAccountData writes data to json file
func writeAccountData() {
	filename := os.Getenv("PATH_TO_GITUSER_CONFIG")
	err := checkFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	data := models.Account{
		PersonalUsername: pUsername,
		PersonalEmail:    pEmail,
		WorkUsername:     wUsername,
		WorkEmail:        wEmail,
		SchoolUsername:   sUsername,
		SchoolEmail:      sEmail,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	ioutil.WriteFile(filename, file, 666)

}

// checkFile makes sure a files exists
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
