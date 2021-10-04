package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
    "os"
    "strings"
)

var config struct {
    CreateMode		bool 	    `docopt:"create"`
    DestroyMode     bool        `docopt:"destroy"`
    PurgeMode       bool        `docopt:"purge"`
    RegionsMode		bool	    `docopt:"regions"`

    Regions			[]string    `docopt:"REGION"`

    AutoConnect		bool	    `docopt:"-c,--connect"`
    KeepOvpn        bool        `docopt:"-k,--keep-ovpn"`
    AutoApprove     bool        `docopt:"-y"`

    PrintHelp 		bool 	    `docopt:"-h,--help"`
    PrintVersion 	bool 	    `docopt:"--version"`
    Verbose			bool	    `docopt:"-v,--verbose"`
}

func main() {
    usage := `Provisions and destroys VPN servers.

Usage: 
  autovpn create [-vcy] REGION ...
  autovpn destroy [-vky] REGION ...
  autovpn purge [-vy]
  autovpn regions [-v]
  autovpn -h | --help
  autovpn --version

Commands:
  create    Create server(s) in region(s)
  destroy   Destroy server(s) in region(s)
  purge     Destroy all servers across all regions
  regions   List all available regions

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

    if config.RegionsMode {
        regions := getRegions()
        fmt.Println(strings.Join(regionStrings(regions), ", "))
        os.Exit(0)
    }

    if config.CreateMode || config.DestroyMode {
        if len(config.Regions) == 0 {
            fmt.Println("No region selected")
            os.Exit(1)
        }

        regions := getRegions()
        for _, region := range config.Regions {
            if !isRegion(region, regions) {
                fmt.Printf("Illegal region %s", region)
                os.Exit(1)
            }
        }
    }
}
