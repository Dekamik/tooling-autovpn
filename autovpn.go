package main

import (
    "fmt"
    "os"
)

var version = "DEVELOPMENT_BUILD" // Gets changed on GitHub Actions

func main() {
    err := bindOptions(os.Args[1:], version)
    check(err)
    err = readConfig()
    check(err)

    if options.ShowRegions {
        err = showRegions()
        check(err)
        os.Exit(0)
    }

    if options.ShowTypes {
        err = showTypes()
        check(err)
        os.Exit(0)
    }

    err = checkPrerequisites()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    if len(options.ConnectTo) != 0 {
        var path string
        if profile, ok := config.Profiles[options.ConnectTo]; ok {
            path = profile.Path
        } else {
            fmt.Printf("Profile %s not found in config, exiting...\n", options.ConnectTo)
            os.Exit(1)
        }

        fmt.Printf("Connecting with file in %s\n", path)

        err = ovpnStart(path, true)
        check(err)
        os.Exit(0)
    }

    if regionValid, err := isRegionValid(options.Region); err != nil || !regionValid {
        check(err)
        fmt.Printf("Region %s is not a valid Linode region, exiting...\n", options.Region)
        os.Exit(1)
    }

    if typeValid, err := isTypeValid(options.LinType); err != nil || !typeValid {
        check(err)
        fmt.Printf("Type %s is not a valid Linode instance type, exiting...\n", options.LinType)
        os.Exit(1)
    }

    err = create()
    check(err)

    err = ovpnStart(fmt.Sprintf("%s/%s-%s.ovpn", config.WorkingDir, config.Hostname, options.Region), false)
    check(err)

    err = purge()
    check(err)

    fmt.Printf("Deleting %s...\n", config.WorkingDir)
    err = os.RemoveAll(config.WorkingDir)
    check(err)
}
