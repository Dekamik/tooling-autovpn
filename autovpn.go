package main

import (
    "encoding/json"
    "fmt"
    "github.com/docopt/docopt-go"
    "os"
    "sort"
    "strings"
)

var config struct {
    CreateCmd  bool `docopt:"create"`
    DestroyCmd bool `docopt:"destroy"`
    PurgeCmd   bool `docopt:"purge"`
    RegionsCmd bool `docopt:"regions"`

    Regions	[]string `docopt:"REGION"`

    AutoConnect	bool `docopt:"-c,--connect"`
    KeepOvpn    bool `docopt:"-k,--keep-ovpn"`
    PrintJson   bool `docopt:"--json"`
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

    opts, _ := docopt.ParseArgs(usage, os.Args[1:], "v1.0.0")
    bindErr := opts.Bind(&config)
    check(bindErr)

    if config.RegionsCmd {
        verbose("Fetching regions...")
        regions := getRegions()
        verboseln("OK")

        var str = ""
        if config.PrintJson {
            jsonBytes, jsonErr := json.Marshal(regions)
            str = string(jsonBytes)
            check(jsonErr)
        } else {
            var regionStrings []string
            for _, region := range regions {
                regionStrings = append(regionStrings, fmt.Sprintf("%s: %s", region.Id, region.Country))
            }
            sort.Strings(regionStrings)
            str = strings.Join(regionStrings, "\n")
        }

        fmt.Println(str)
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
