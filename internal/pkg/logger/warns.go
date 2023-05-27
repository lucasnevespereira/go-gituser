package logger

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func PrintWarningReadingAccount(mode string) {
	fmt.Println("")
	fmt.Fprintf(os.Stderr, color.YellowString("Warning: ")+"%v \n", "You have no "+mode+" account 🧐")
	fmt.Println("")
	color.Cyan("Tips:")
	fmt.Printf("To add a %v account try to run gituser config \n", mode)
	fmt.Println("")
}

func PrintRemeberToActiveMode(mode string) {
	info := color.New(color.Bold).PrintfFunc()
	info(color.BlueString("Remember to run <gituser %v> to activate this mode \n"), mode)
}

func PrintNoActiveMode() {
	fmt.Fprint(os.Stderr, color.YellowString("Active account not found \n"))
	fmt.Println("")
	info := color.New(color.Bold).PrintfFunc()
	info(color.BlueString("Run <gituser %v> to setup accounts \n"), "config")
	info(color.BlueString("Run <gituser %v> to activate a mode \n"), "(work,pesonal,school)")
}

func PrintUnsavedActiveMode() {
	fmt.Fprint(os.Stderr, color.YellowString("Active account not found \n"))
	fmt.Println("")
	info := color.New(color.Bold).PrintfFunc()
	info(color.BlueString("Run <gituser %v> to setup accounts \n"), "config")
	info(color.BlueString("Run <gituser %v> to activate a mode \n"), "(work,pesonal,school)")
}