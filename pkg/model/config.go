package model

import (
	"fmt"

	"github.com/elliot40404/ssm/pkg/utils"
)

type Config struct {
	Name     string
	Hostname string
	Port     string
	User     string
	Category string
	SshKey   string
}

func (c *Config) Save(useKey bool) {
	var config string
	if useKey {
		config = fmt.Sprintf("Host %v\n  HostName %v\n  IdentityFile %v\n  User %v\n  Port %v", c.Name, c.Hostname, c.SshKey, c.User, c.Port)
	} else {
		config = fmt.Sprintf("Host %v\n  HostName %v\n  User %v\n  Port %v", c.Name, c.Hostname, c.User, c.Port)
	}
	utils.CreateConfig(c.Name, config)
	utils.LinkConfig(c.Name, c.Category)
	fmt.Println("Config saved! You can now access it using: ssh", c.Name)
}
