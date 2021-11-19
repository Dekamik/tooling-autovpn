package main

import (
    "github.com/docopt/docopt-go"
    "github.com/spf13/viper"
    "strings"
)

type Options struct {
    Region     string `docopt:"REGION"`
    ConnectTo  string `docopt:"-c"`

    ShowRegions bool `docopt:"--show-regions"`
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
  autovpn (linode | digitalocean | aws | azure | gcp) REGION
  autovpn (linode | digitalocean | aws | azure | gcp) --show-regions [--json | --no-headers]
  autovpn -c PROFILE
  autovpn -h | --help
  autovpn --version

Arguments:
  REGION  Linode region for server. Find avaiable regions by running "autovpn regions"

Providers:
  linode        Linode
  digitalocean  Digital Ocean
  aws           Amazon Web Services
  azure         Microsoft Azure
  gcp           Google Cloud Platform

Options:
  -c PROFILE      Connect to pre-defined VPN profile
  --show-regions  Show available regions.
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
