package model

import (
	"fmt"
	"github.com/elliot40404/ssm/pkg/utils"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Config struct {
	Name     string `validate:"required"`
	Hostname string `validate:"required,fqdn|ip"`
	Port     int    `validate:"required,min=1,max=65535"`
	User     string `validate:"required"`
	Category string `validate:"required"`
	SshKey   string `validate:"omitempty,filepath"`
}

func (c *Config) Save(useKey bool, dry bool) {
	var config string
	if useKey {
		config = fmt.Sprintf("Host %v\n  HostName %v\n  IdentityFile %v\n  User %v\n  Port %v", c.Name, c.Hostname, c.SshKey, c.User, c.Port)
	} else {
		config = fmt.Sprintf("Host %v\n  HostName %v\n  User %v\n  Port %v", c.Name, c.Hostname, c.User, c.Port)
	}
	if dry {
		fmt.Println("Dry run: \n", config)
		return
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
			config.Port = utils.ParseInt(strings.TrimSpace(strings.TrimPrefix(line, "Port ")))
		}
	}
	return *config
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(c)
}
