# SSM - THE SECURE SHELL MANAGER

```                 
 _____ _____ _____ 
|   __|   __|     |
|__   |__   | | | |
|_____|_____|_|_|_|

SSM - THE SECURE SHELL MANAGER
                      - Elliot
```

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
