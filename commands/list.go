package commands

import (
	"fmt"

	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/urfave/cli/v2"
)

func List(cCtx *cli.Context) error {
	var output string
	for _, name := range utils.GetSSHConfigs(false) {
		output += fmt.Sprintf("%v\n", name)
	}
	fmt.Print(output)
	return nil
}
