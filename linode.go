package main

var linodeMainTemplate = `terraform {
  required_providers {
    linode = {
      source = "linode/linode"
      version = "1.16.0"
    }
  }
}

provider "linode" {
  token = "{{.Token}}"
}
`

var linodeVpnTemplate = `module "{{.Name}}" {
  source = "git@github.com:Dekamik/vpn-modules.git//vpn-server?ref=v0.2.1"

  token = "{{.Token}}"
  public_keys = {
    "{{.Hostname}}" = "{{.PublicKey}}"
  }

  name = "{{.Hostname}}-{{.Name}}"
  region = "{{.Region}}"
  type = "{{.Type}}"
  download_dir = "{{.DownloadDir}}"
}
`
