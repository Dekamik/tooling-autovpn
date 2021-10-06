package main

import (
    "fmt"
    "os"
)

func destroy(token string) error {
    homeDir, _ := os.UserHomeDir()
    files := make([]string, len(options.Regions))
    for i, r := range options.Regions {
        files[i] = fmt.Sprintf(homeDir + "/.autovpn/%s.tf", r)
    }

    removeErr := removeFiles(files)
    if removeErr != nil {
        return removeErr
    }

    return nil
}
