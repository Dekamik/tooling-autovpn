package main

import (
    "encoding/json"
    "fmt"
    "github.com/docopt/docopt-go"
    "os"
)

var config struct {
    CreateCmd  bool `docopt:"create"`
    DestroyCmd bool `docopt:"destroy"`
    PurgeCmd   bool `docopt:"purge"`
    RegionsCmd bool `docopt:"regions"`

    Regions	[]string `docopt:"REGION"`

    AutoConnect	bool `docopt:"-c,--connect"`
    KeepOvpn    bool `docopt:"-k,--keep-ovpn"`
    AutoApprove bool `docopt:"-y"`

    PrintHelp    bool `docopt:"-h,--help"`
    PrintVersion bool `docopt:"--version"`
    Verbose		 bool `docopt:"-v,--verbose"`
}

func main() {
    usage := `Provisions and destroys VPN servers.

Usage: 
  autovpn create [-cvy] REGION ...
  autovpn destroy [-kvy] REGION ...
  autovpn purge [-vy]
  autovpn regions [-v]
  autovpn -h | --help
  autovpn --version

Commands:
  create    Create server(s) in region(s)
  destroy   Destroy server(s) in region(s)
  purge     Destroy all servers across all regions
  regions   List all available regions as JSON

Arguments:
  REGION    Linode region for server. Find avaiable regions by running "autovpn regions"

Options:
  -c --connect      Auto-connect with OpenVPN. (requires root privileges)
  -k --keep-ovpn    Keep .ovpn-config.
  -v --verbose      Print more text.
  -y                Auto-approve.
  -h --help         Show this screen.
  --version         Show version.`

    opts, _ := docopt.ParseArgs(usage, os.Args[1:], "v1.0.0")
    bindErr := opts.Bind(&config)
    check(bindErr)

    if config.RegionsCmd {
        verbose("Fetching regions...")
        regions := getRegions()
        verboseln("OK")
        jsonStr, jsonErr := json.Marshal(regions)
        check(jsonErr)

        fmt.Println(string(jsonStr))
        os.Exit(0)
    }

    if config.CreateCmd || config.DestroyCmd {
        if len(config.Regions) == 0 {
            fmt.Println("No region specified.")
            os.Exit(1)
        }

        verbose("Fetching regions...")
        regions := getRegions()
        verboseln("OK")
        for _, region := range config.Regions {
            if !isRegion(region, regions) {
                fmt.Printf("Illegal region %s.", region)
                os.Exit(1)
            }
        }
    }
}
