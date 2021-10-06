package main

import (
    "fmt"
    "os"
)

func destroy(token string) error {
    homeDir, _ := os.UserHomeDir()
    summary := make([][]string, len(options.Regions))
    if options.Verbose {
        defer printTable(summary)
    }

    for i, r := range options.Regions {
        fileName := fmt.Sprintf(homeDir + "/.autovpn/%s.tf", r)
        if _, err := os.Stat(fileName); err == nil {
            removeErr := os.Remove(fileName)
            if removeErr != nil {
                summary[i] = []string { fileName, "Error" }
                return removeErr
            }

            summary[i] = []string { fileName, "Removed" }
        } else {

            summary[i] = []string { fileName, "Not found" }
        }
    }

    return nil
}
