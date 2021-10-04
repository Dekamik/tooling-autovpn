package main

import (
    "fmt"
    "os"
)

func check(e error) {
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
        check(validationErr)
        _, tokenErr := findToken()
        check(tokenErr)

        os.Exit(0)
    }

    if options.DestroyCmd {
        validationErr := validateRegions(options.Regions)
        check(validationErr)

        os.Exit(0)
    }

    if options.PurgeCmd {
        os.Exit(0)
    }

    if options.RegionsCmd {
        err := showRegions()
        check(err)

        os.Exit(0)
    }
}
