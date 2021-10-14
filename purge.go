package main

import (
    "os"
)

func purge(token string) error {
    homeDir, _ := os.UserHomeDir()
    files := find(homeDir + "/.autovpn/", ".tf")
    if len(files) == 0 {
        return nil
    }

    filesRemoved, removeErr := removeFiles(files)
    if removeErr != nil {
        return removeErr
    }

    if filesRemoved != 0 {
        err := tfApply()
        if err != nil {
            return err
        }
    }

    return nil
}
