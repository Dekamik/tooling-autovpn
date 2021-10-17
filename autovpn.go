package main

import (
    "fmt"
    "os"
)

func checkGracefully(e error) {
    if e != nil {
        fmt.Println(e)
        os.Exit(1)
    }
}

func main() {
    bindErr := bindOptions(os.Args[1:], "v1.0.0")
    check(bindErr)
    readErr := readConfig()
    check(readErr)

    if options.CreateCmd {
        validationErr := validateRegions(options.Regions)
        checkGracefully(validationErr)

        createErr := create()
        check(createErr)

        os.Exit(0)
    }

    if options.DestroyCmd {
        validationErr := validateRegions(options.Regions)
        checkGracefully(validationErr)

        destroyErr := destroy()
        check(destroyErr)

        os.Exit(0)
    }

    if options.PurgeCmd {
        purgeErr := purge()
        check(purgeErr)

        os.Exit(0)
    }

    if options.RegionsCmd {
        err := showRegions()
        check(err)

        os.Exit(0)
    }

    if options.StatusCmd {
        err := showStatus()
        check(err)

        os.Exit(0)
    }
}
