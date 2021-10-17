package main

import (
    "fmt"
    "os"
    "strings"
)

func purge() error {
    homeDir, _ := os.UserHomeDir()
    files := find(homeDir + "/.autovpn/", ".tf")
    if len(files) == 0 {
        return nil
    }

    fmt.Println(files)

    for i, file := range files {
        if strings.Contains(file, "/.autovpn/main.tf") {
            if i == len(files) - 1 {
                files = files[:len(files) - 1]
                continue
            }
            files[i] = files[len(files) - 1]
            files = files[:len(files) - 1]
        }
    }

    filesRemoved, removeErr := removeFiles(files)
    if removeErr != nil {
        return removeErr
    }

    if filesRemoved == 0 {
        return nil
    }

    planErr := tfPlan()
    if planErr != nil {
        return planErr
    }

    applyErr := tfApply()
    if applyErr != nil {
        return applyErr
    }

    return nil
}
