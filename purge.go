package main

import "os"

func purge(token string) error {
    homeDir, _ := os.UserHomeDir()
    files := find(homeDir + "/.autovpn/", ".tf")
    if len(files) == 0 {
        return nil
    }
    summary := make([][]string, len(files))
    if options.Verbose {
        defer printTable(summary)
    }

    for i, file := range files {
        if _, err := os.Stat(file); err == nil {
            removeErr := os.Remove(file)
            if removeErr != nil {
                summary[i] = []string { file, "Error" }
                return removeErr
            }

            summary[i] = []string { file, "Removed" }
        } else {

            summary[i] = []string { file, "Not found" }
        }
    }

    return nil
}
