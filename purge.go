package main

import (
    "fmt"
    "strings"
)

func purge() error {
    fmt.Println("Destroying VPN server...")
    files := find(config.WorkingDir, ".tf")
    if len(files) == 0 {
        return nil
    }

    for i, file := range files {
        if strings.Contains(file, config.WorkingDir + "/main.tf") {
            if i == len(files) - 1 {
                files = files[:len(files) - 1]
                continue
            }
            files[i] = files[len(files) - 1]
            files = files[:len(files) - 1]
        }
    }

    filesRemoved, err := removeFiles(files)
    if err != nil { return err }

    if filesRemoved == 0 {
        return nil
    }

    err = tfPlan()
    if err != nil { return err }

    err = tfApply()
    if err != nil { return err }

    return nil
}
