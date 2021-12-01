package main

import (
    "github.com/docopt/docopt-go"
    "github.com/spf13/viper"
    "strings"
)

type Options struct {
    Region     string `docopt:"REGION"`
    ConnectTo  string `docopt:"-c"`
    LinType    string `docopt:"-t"`

    ShowRegions bool `docopt:"--show-regions"`
    ShowTypes   bool `docopt:"--show-types"`
    NoHeaders   bool `docopt:"--no-headers"`
    PrintJson   bool `docopt:"--json"`

    PrintHelp    bool `docopt:"-h,--help"`
    PrintVersion bool `docopt:"--version"`
}

type Profile struct {
    Path string `mapstructure:"path"`
}

type Config struct {
    Hostname   string             `mapstructure:"hostname"`
    Token      string             `mapstructure:"token"`
    WorkingDir string             `mapstructure:"workingdir"`
    SshPath    string             `mapstructure:"sshpath"`
    Profiles   map[string]Profile `mapstructure:"profiles"`
}

var options Options
var config Config

var usage = `Tool for provisioning and connecting to a temporary VPN server.
This server gets destroyed when the connection is terminated.

Usage: 
  autovpn [-t TYPE_ID] REGION
  autovpn -c PROFILE
  autovpn --show-regions [--json | --no-headers]
  autovpn --show-types
  autovpn -h | --help
  autovpn --version

Arguments:
  REGION  Linode region for server. Find avaiable regions by running "autovpn regions"

Options:
  -c PROFILE      Connect to pre-defined VPN profile
  -t TYPE_ID      Linode instance type to spawn [default: g6-dedicated-2]
  --show-regions  Show available regions.
  --show-types    Show available linode types
  --json          Print as JSON.
  --no-headers	  Suppress printout headers
  -h --help       Show this screen.
  --version       Show version.`

func bindOptions(argv []string, semver string) error {
    opts, _ := docopt.ParseArgs(usage, argv, semver)
    err := opts.Bind(&options)
    if err != nil {
        return err
    }
    return nil
}

func readConfig() error {
    viper.AddConfigPath("$HOME/.autovpn")
    viper.SetConfigName("config")
    viper.SetConfigType("toml")
    err := viper.ReadInConfig()
    if err != nil {
        if strings.Contains(err.Error(), "Not Found") {
            return nil
        }
        return err
    }

    err = viper.Unmarshal(&config)
    if err != nil { return err }
    return nil
}
