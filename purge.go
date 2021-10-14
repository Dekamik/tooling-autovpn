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

    if filesRemoved == 0 {
        return nil
    }

    initErr := tfInit()
    if initErr != nil {
        return initErr
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
