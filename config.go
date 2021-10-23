package main

import (
    "github.com/docopt/docopt-go"
    "github.com/spf13/viper"
    "strings"
)

var options struct {
    ConfigPath string `docopt:"--config"`
    Region     string `docopt:"REGION"`

    ShowRegions bool `docopt:"--show-regions"`
    NoHeaders   bool `docopt:"--no-headers"`
    PrintJson   bool `docopt:"--json"`

    PrintHelp    bool `docopt:"-h,--help"`
    PrintVersion bool `docopt:"--version"`
}

var config struct {
    Hostname   string `mapstructure:"hostname"`
    Token      string `mapstructure:"token"`
    WorkingDir string `mapstructure:"workingdir"`
    SshPath    string `mapstructure:"sshpath"`
}

var usage = `Tool for provisioning and connecting to temporary VPN servers.
These servers get deleted when the connection is exited by pressing 'q' during the session.

Usage: 
  autovpn [--config=<config>] REGION
  autovpn --show-regions [--json | --no-headers]
  autovpn -h | --help
  autovpn --version

Arguments:
  REGION  Linode region for server. Find avaiable regions by running "autovpn regions"

Options:
  --show-regions  Show available regions.
  --json          Print as JSON.
  --no-headers	  Suppress printout headers
  -h --help       Show this screen.
  --version       Show version.`

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
    viper.AddConfigPath(options.ConfigPath)
    viper.SetConfigName("config")
    viper.SetConfigType("toml")
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
