package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/slices"
)

func BasicPrompt(prompt string, nullable bool) string {
	var input string
	for {
		fmt.Println(prompt)
		fmt.Print("> ")
		fmt.Scanln(&input)
		if input != "" || nullable {
			break
		}
	}
	return input
}

func CheckExisting(name string) bool {
	return slices.Contains(GetSSHConfigs(false), name)
}

func SelectMenu(promptText string, items []string, addNew bool) string {
	options := items
	if addNew {
		options = append(options, "Add New")
	}
	prompt := promptui.Select{
		Label:        promptText,
		Items:        options,
		HideSelected: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Something went wrong")
	}
	return result
}

func AskYesNo(promptText string, defaultValue string) bool {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "" {
		input = defaultValue
	}
	return input == "y" || input == "yes"
}
