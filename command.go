package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func run(command string) error {
    fmt.Printf("+ %s\n", command)
    cmdArgs := strings.Fields(command)
    cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        return err
    }

    return nil
}
