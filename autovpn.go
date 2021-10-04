package main

import (
    "fmt"
    "os"
)

func handle(e error) {
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
        handle(validationErr)
        token, tokenErr := findToken()
        handle(tokenErr)

        createErr := create(token)
        check(createErr)

        os.Exit(0)
    }

    if options.DestroyCmd {
        validationErr := validateRegions(options.Regions)
        handle(validationErr)
        token, tokenErr := findToken()
        handle(tokenErr)

        destroyErr := destroy(token)
        check(destroyErr)

        os.Exit(0)
    }

    if options.PurgeCmd {
        token, tokenErr := findToken()
        handle(tokenErr)

        purgeErr := purge(token)
        check(purgeErr)

        os.Exit(0)
    }

    if options.RegionsCmd {
        err := showRegions()
        check(err)

        os.Exit(0)
    }
}
