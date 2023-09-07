package commands

import (
	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/urfave/cli/v2"
)

func Edit(cCtx *cli.Context) error {
	configName := cCtx.Args().First()
	if configName == "" {
		configName = utils.SelectMenu("Select Config", utils.GetSSHConfigs(false), false)
	}
	utils.OpenWithDefaultTextEditor(utils.GetSSHDir(true) + "/" + configName)
	return nil
}
