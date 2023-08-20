package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/elliot40404/ssm/pkg/model"
	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

func Add(cCtx *cli.Context) error {
	var name string
	for {
		name = utils.BasicPrompt("Name your config", false)
		if !checkExisting(name) {
			break
		}
		fmt.Println("Config with same name already exists")
	}
	// TODO: ADD CONFIG VALIDATION
	var config model.Config
	config.Name = name
	config.Hostname = utils.BasicPrompt("Hostname", false)
	config.User = utils.BasicPrompt("User", false)
	config.Port = utils.BasicPrompt("Port (Enter to use Default 22)", true)
	if config.Port == "" {
		config.Port = "22"
	}
	// ADD SSH KEY
	usekey := askYesNo("Use SSH Key [Y|n]: ", "y")
	fmt.Println(usekey)
	if usekey {
		result := selectMenu("Select SSH Key", utils.GetSSHKeys(), true)
		if result == "Add New" {
			config.SshKey = utils.BasicPrompt("Full path to ssh key", false)
		} else {
			config.SshKey = "~/.ssh/" + result
		}
	}
	// ADD CATEGORY
	result := selectMenu("Select Category", utils.GetCategories(), true)
	if result == "Add New" {
		config.Category = utils.BasicPrompt("Enter New Category (without #)", false)
	} else {
		config.Category = result
	}
	if !cCtx.Bool("dry-run") {
		config.Save(usekey)
	}
	return nil
}

func checkExisting(name string) bool {
	return slices.Contains(utils.GetSSHConfigs(false), name)
}

func selectMenu(promptText string, items []string, addNew bool) string {
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

func askYesNo(promptText string, defaultValue string) bool {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "" {
		input = defaultValue
	}
	return input == "y" || input == "yes"
}
