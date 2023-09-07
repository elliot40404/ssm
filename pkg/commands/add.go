package commands

import (
	"fmt"

	"github.com/elliot40404/ssm/pkg/model"
	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/urfave/cli/v2"
)

func Add(cCtx *cli.Context) error {
	var name string
	for {
		name = utils.BasicPrompt("Name your config", false)
		if !utils.CheckExisting(name) {
			break
		}
		fmt.Println("Config with same name already exists")
	}
	var config model.Config
	config.Name = name
	config.Hostname = utils.BasicPrompt("Hostname", false)
	config.User = utils.BasicPrompt("User", false)
	port := utils.BasicPrompt("Port (Enter to use Default 22)", true)
	if port != "" {
		config.Port = utils.ParseInt(port)
	}
	config.Port = 22
	// ADD SSH KEY
	usekey := utils.AskYesNo("Use SSH Key [Y|n]: ", "y")
	if usekey {
		result := utils.SelectMenu("Select SSH Key", utils.GetSSHKeys(), true)
		if result == "Add New" {
			config.SshKey = utils.BasicPrompt("Full path to ssh key", false)
		} else {
			config.SshKey = "~/.ssh/" + result
		}
	}
	// ADD CATEGORY
	result := utils.SelectMenu("Select Category", utils.GetCategories(), true)
	if result == "Add New" {
		config.Category = utils.BasicPrompt("Enter New Category (without #)", false)
	} else {
		config.Category = result
	}
	// VALIDATE CONFIG
	err := config.Validate()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	config.Save(usekey, cCtx.Bool("dry-run"))
	return nil
}
