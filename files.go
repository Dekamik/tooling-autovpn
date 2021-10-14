package main

import (
    "io/fs"
    "os"
    "path/filepath"
)

func removeFiles(files []string) (int, error) {
    var filesRemoved = 0
    summary := make([][]string, len(files))
    if options.Verbose {
        defer printTable(summary)
    }

    for i, file := range files {
        if _, err := os.Stat(file); err == nil {
            removeErr := os.Remove(file)
            if removeErr != nil {
                summary[i] = []string { file, "Error" }
                return filesRemoved, removeErr
            }
            summary[i] = []string { file, "Removed" }
            filesRemoved++
        } else {
            summary[i] = []string { file, "Not found" }
        }
    }

    return filesRemoved, nil
}

func find(root, ext string) []string {
    var a []string
    _ = filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
        if e != nil {
            return e
        }
        if filepath.Ext(d.Name()) == ext {
            a = append(a, s)
        }
        return nil
    })
    return a
}
