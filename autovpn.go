package main

import (
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        fmt.Println(e)
        os.Exit(1)
    }
}

func main() {
    usage := `Provisions and destroys VPN servers.

Usage: 
  autovpn create [-cvy] REGION ...
  autovpn destroy [-kvy] REGION ...
  autovpn purge [-vy]
  autovpn regions [-v] [--json]
  autovpn -h | --help
  autovpn --version

Commands:
  create   Create server(s) in region(s)
  destroy  Destroy server(s) in region(s)
  purge    Destroy all servers across all regions
  regions  List all available regions

Arguments:
  REGION  Linode region for server. Find avaiable regions by running "autovpn regions"

Options:
  -c --connect    Auto-connect with OpenVPN. (requires root privileges)
  -k --keep-ovpn  Keep .ovpn-config.
  --json          Print as JSON.
  -y              Auto-approve.
  -v --verbose    Print more text.
  -h --help       Show this screen.
  --version       Show version.`

    bindErr := bindConfig(usage, os.Args[1:], "v1.0.0")
    check(bindErr)

    if config.CreateCmd {
        validationErr := validateRegions(config.Regions)
        check(validationErr)

        os.Exit(0)
    }

    if config.DestroyCmd {
        validationErr := validateRegions(config.Regions)
        check(validationErr)

        os.Exit(0)
    }

    if config.PurgeCmd {
        os.Exit(0)
    }

    if config.RegionsCmd {
        err := showRegions()
        check(err)

        os.Exit(0)
    }
}
