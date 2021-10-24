package main

import (
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "strings"
    "syscall"
)

func tfApply() error {
    return run(fmt.Sprintf("terraform -chdir=%s apply %s/.terraform/tfplan", config.WorkingDir, config.WorkingDir))
}

func tfInit() error {
    return run(fmt.Sprintf("terraform -chdir=%s init", config.WorkingDir))
}

func tfPlan() error {
    return run(fmt.Sprintf("terraform -chdir=%s plan -out %s/.terraform/tfplan", config.WorkingDir, config.WorkingDir))
}

func ovpnConnect(configPath string, stdin bool) error {
    cmd := exec.Command("sudo", "openvpn", configPath)
    if stdin {
        cmd.Stdin = os.Stdin
    }
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    fmt.Println("Opening VPN tunnel...")
    if startErr := cmd.Start(); startErr != nil {
        return startErr
    }

    var waiting = true

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT,
        os.Interrupt)

    go func() {
        s := <-sigc
        fmt.Printf("%s signal recieved, killing session...\n", strings.Title(s.String()))
        _ = cmd.Process.Kill()
        waiting = false
    }()

    for waiting {}

    return nil
}

func run(command string) error {
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
