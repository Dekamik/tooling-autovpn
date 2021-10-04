package main

import (
	"github.com/docopt/docopt-go"
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

func bindConfig(usage string, argv []string, semver string) error {
	opts, _ := docopt.ParseArgs(usage, argv, semver)
	bindErr := opts.Bind(&config)
	if bindErr != nil {
		return bindErr
	}
	return nil
}
