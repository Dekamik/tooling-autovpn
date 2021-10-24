package main

import (
    "fmt"
    "os"
)

func main() {
    bindErr := bindOptions(os.Args[1:], "v1.0.0")
    check(bindErr)
    readErr := readConfig()
    check(readErr)

    if options.ShowRegions {
        err := showRegions()
        check(err)

        os.Exit(0)
    }

    if len(options.ConnectTo) != 0 {
        var path string
        for key, profile := range config.Profiles {
            if key == options.ConnectTo {
                path = profile.Path
                break
            }
        }

        if path == "" {
            fmt.Printf("Couldn't find path for profile %s in config, exiting...\n", options.ConnectTo)
            os.Exit(1)
        }
        fmt.Printf("Connecting with file in %s\n", path)

        connectToErr := ovpnConnect(path, true)
        check(connectToErr)
        os.Exit(0)
    }

    createErr := create()
    check(createErr)

    ovpnErr := ovpnConnect(fmt.Sprintf("%s/%s-%s.ovpn", config.WorkingDir, config.Hostname, options.Region), true)
    check(ovpnErr)

    purgeErr := purge()
    check(purgeErr)

    fmt.Printf("Deleting %s...", config.WorkingDir)
    rmErr := os.RemoveAll(config.WorkingDir)
    check(rmErr)
}
