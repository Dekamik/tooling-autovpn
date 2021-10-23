package main

import (
    "strings"
)

func purge() error {
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
