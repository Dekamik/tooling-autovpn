package main

import "os"

func purge(token string) error {
    homeDir, _ := os.UserHomeDir()
    files := find(homeDir + "/.autovpn/", ".tf")
    if len(files) == 0 {
        return nil
    }

    removeErr := removeFiles(files)
    if removeErr != nil {
        return removeErr
    }

    return nil
}
