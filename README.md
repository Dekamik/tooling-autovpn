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
sshpath = "/home/<youruser>/.ssh/id_ed25519.pub"
token = "<your-linode-token>"
workingdir = "/home/<youruser>/.autovpn/tmp"
```

### Additional configurations
You can also define profile blocks for connecting to existing VPN servers with .ovpn files on your computer.
Simply define the blocks as follows:

```toml
[profiles]
  [profiles.home]
  path = "/path/to/home/client.ovpn"
  
  [profiles.<any>]
  path = "/path/to/any.ovpn"
```

With above configurations you simply call `autovpn -c <profile>` 
where `<profile>` is the name you gave the profile block after the dot. 

E.g. to connect to the `home` profile, call `autovpn -c home`. 
The program will open the tunnel and forward stdin, stdout and stderr.

## Usage
When all above is done, you can start using the command.

Create and connect to a vpn server by running `autovpn <region>`.
When the session is created, simply send SIGTERM (Ctrl-C) to disconnect and destroy the server.

You can view a list of available regions by running `autovpn --show-regions`.

For additional help, run `autovpn -h`.
