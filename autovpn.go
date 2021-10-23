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

    createErr := create()
    check(createErr)

    ovpnErr := ovpnConnect(fmt.Sprintf("%s/%s-%s.ovpn", config.WorkingDir, config.Hostname, options.Region))
    check(ovpnErr)

    purgeErr := purge()
    check(purgeErr)

    rmErr := os.RemoveAll(config.WorkingDir)
    check(rmErr)
}
