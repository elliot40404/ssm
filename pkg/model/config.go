package model

import (
	"fmt"
	"strings"

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

func (c *Config) Parse(input string) Config {
	lines := strings.Split(input, "\n")
	config := &Config{}
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "Host ") {
			config.Name = strings.TrimSpace(strings.TrimPrefix(line, "Host "))
		} else if strings.HasPrefix(line, "HostName ") {
			config.Hostname = strings.TrimSpace(strings.TrimPrefix(line, "HostName "))
		} else if strings.HasPrefix(line, "User ") {
			config.User = strings.TrimSpace(strings.TrimPrefix(line, "User "))
		} else if strings.HasPrefix(line, "IdentityFile ") {
			config.SshKey = strings.TrimSpace(strings.TrimPrefix(line, "IdentityFile "))
		} else if strings.HasPrefix(line, "Port ") {
			config.Port = strings.TrimSpace(strings.TrimPrefix(line, "Port "))
		}
	}
	return *config
}
