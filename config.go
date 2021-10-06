package main

import (
    "errors"
    "github.com/docopt/docopt-go"
    "github.com/spf13/viper"
    "strings"
)

var options struct {
    CreateCmd  bool `docopt:"create"`
    DestroyCmd bool `docopt:"destroy"`
    PurgeCmd   bool `docopt:"purge"`
    RegionsCmd bool `docopt:"regions"`
    StatusCmd  bool `docopt:"status"`

    Regions	[]string `docopt:"REGION"`

    Token       string `docopt:"-t,--token"`
    AutoConnect	bool   `docopt:"-c,--connect"`
    KeepOvpn    bool   `docopt:"-k,--keep-ovpn"`
    ApplyOnAll  bool   `docopt:"-a,--all"`
    PrintJson   bool   `docopt:"--json"`
    NoHeaders   bool   `docopt:"--no-headers"`
    AutoApprove bool   `docopt:"-y,--auto-approve"`

    PrintHelp    bool `docopt:"-h,--help"`
    PrintVersion bool `docopt:"--version"`
    Verbose		 bool `docopt:"-v,--verbose"`
}

var config struct {
    ApiToken string `mapstructure:"api-token"`
}

var usage = `Provisions and destroys VPN servers.

Usage: 
  autovpn create [-cvy] [-t TOKEN | --token TOKEN] REGION ...
  autovpn destroy [-kvy] [-t TOKEN | --token TOKEN] REGION ...
  autovpn purge [-avy] [-t TOKEN | --token TOKEN]
  autovpn regions [-v] [--json | --no-headers]
  autovpn status [-av] [--json | --no-headers]
  autovpn -h | --help
  autovpn --version

Commands:
  create   Create server(s) in region(s)
  destroy  Destroy server(s) in region(s)
  purge    Destroy all servers across all regions
  regions  List all available regions
  status   VPN server status

Arguments:
  REGION  Linode region for server. Find avaiable regions by running "autovpn regions"

Options:
  -t TOKEN --token TOKEN  API token for your Linode account
  -c --connect            Auto-connect with OpenVPN. (requires root privileges)
  -k --keep-ovpn          Keep .ovpn-options.
  -a --all				  Run command on all servers on your account, not only those associated with your computer.
  --json                  Print as JSON.
  --no-headers			  Suppress printout headers
  -y --auto-approve       Approve changes automatically.
  -v --verbose            Print more text.
  -h --help               Show this screen.
  --version               Show version.`

func bindOptions(argv []string, semver string) error {
    opts, _ := docopt.ParseArgs(usage, argv, semver)
    bindErr := opts.Bind(&options)
    if bindErr != nil {
        return bindErr
    }
    return nil
}

func readConfig() error {
    viper.AddConfigPath("$HOME/.autovpn")
    viper.SetConfigName("config")
    viper.SetConfigType("ini")
    readErr := viper.ReadInConfig()
    if readErr != nil {
        if strings.Contains(readErr.Error(), "Not Found") {
            return nil
        }
        return readErr
    }

    marshalErr := viper.Unmarshal(&config)
    if marshalErr != nil { return marshalErr }
    return nil
}

func findToken() (string, error) {
    if len(config.ApiToken) != 0 {
        return config.ApiToken, nil
    }
    if len(options.Token) != 0 {
        return options.Token, nil
    }
    return "", errors.New("No API token found. Try autovpn --help or check README.md ")
}
