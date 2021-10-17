package main

import (
    "fmt"
    "os"
)

func destroy() error {
    homeDir, _ := os.UserHomeDir()
    files := make([]string, len(options.Regions))
    for i, r := range options.Regions {
        files[i] = fmt.Sprintf(homeDir + "/.autovpn/%s.tf", r)
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
