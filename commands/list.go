package commands

import (
	"fmt"

	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/urfave/cli/v2"
)

func List(cCtx *cli.Context) error {
	for idx, name := range utils.GetSSHConfigs(false) {
		if cCtx.Bool("count") {
			fmt.Printf("%v. %v\n", idx+1, name)
		} else {
			fmt.Printf("%v\n", name)
		}
	}
	return nil
}
