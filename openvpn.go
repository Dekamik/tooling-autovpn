package main

import (
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "syscall"
)

func ovpnConnect(configPath string) error {
    cmd := exec.Command("sudo", "openvpn", configPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    fmt.Println("Opening VPN tunnel...")
    if startErr := cmd.Start(); startErr != nil {
        return startErr
    }

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT,
        os.Interrupt)

    var waiting = true

    go func() {
        s := <-sigc
        fmt.Printf("\n%s signal recieved, killing session...\n", s)
        _ = cmd.Process.Kill()
        waiting = false
    }()

    for waiting {}

    return nil
}
