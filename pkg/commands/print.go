package commands

import (
	"fmt"

	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/urfave/cli/v2"
)

func Print(cCtx *cli.Context) error {
	configName := cCtx.Args().First()
	if configName == "" {
		configName = utils.SelectMenu("Select Config", utils.GetSSHConfigs(false), false)
	}
	// get config
	config := utils.GetSSHConfig(configName)
	// print config
	fmt.Println(config)
	return nil
}
