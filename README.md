# tooling-autovpn
Tool for automatically creating, connecting to and destroying VPN servers on-the-fly.

## Getting started
### Prerequisites
This tool requires the following programs:
* OpenSSH
* Terraform
* Ansible
* OpenVPN

These can be installed on linux by running `sudo bash ./install.sh`

You also need an account on Linode and a generated API token for your account.

### Configuring
You also need to create a config file with this path: `~/.autovpn/config`.
Inside that file you will define some settings as follows:

```toml
hostname = "<your-hostname>"
sshpath = "<path/to/public/ssh/key>"
token = "<your-linode-token>"
workingdir = "/home/<youruser>/.autovpn/tmp"
```

### Usage
When all above is done, you can start using the command.

Create and connect to a vpn server by running `autovpn <region>`.
When the session is created, simply send SIGTERM (Ctrl-C) to disconnect and destroy the server.

You can view a list of available regions by running `autovpn --show-regions`.

For additional help, run `autovpn -h`.
