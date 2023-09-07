# SSM - Simple SSH Manager

```
 _____ _____ _____
|   __|   __|     |
|__   |__   | | | |
|_____|_____|_|_|_|

SSM - SIMPLE SSH MANAGER
                      - Elliot
```

## About

I am learning Go and wanted to make something useful. I have a lot of servers that I need to SSH into and I wanted a way to manage them, so I decided to make a simple SSH manager.

> NOTE
> This is still a work in progress so there may be bugs. Please feel free to report them

## Directory structure

```
/home/elliot/.ssh
├── config *
├── config.d *
│   ├── compute
│   └── homelab
├── known_hosts
├── id_rsa // private key **
└── id_rsa.pub // public key **
```

> `*` - Required

> `**` SSH keys are not required if you use password authorization

> SSH keys can be generated with [ssh-keygen](https://www.ssh.com/academy/ssh/config)

## Config file `~/.ssh/config`

```
# WORK SERVERS
Include config.d/compute

# HOME SERVERS
Include config.d/homelab
```

## Example file in config.d `~/.ssh/config.d/compute`

```
Host compute
  HostName 192.168.0.110
  IdentityFile ~/.ssh/id_rsa
  User ubuntu
  Port 4444
```

## Usage

```
NAME:
   ssm - ssh config manager

USAGE:
   ssm [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   Elliot <admin@elliot404.com>

COMMANDS:
   list, ls  list available ssh configs
   add       add new ssh config
   print     print ssh config
   edit, e   edit ssh config
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## TODO

-   [ ] Remove config
-   [ ] Add support for password authentication
-   [ ] ssh-keygen wrapper

## License

MIT License
