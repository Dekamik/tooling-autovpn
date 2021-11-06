package main

import (
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "strings"
)

type Software int

const (
    Ansible Software = iota
    OpenSSH          = 1
    OpenVPN          = 2
    Terraform        = 3
)

func tfApply() error {
    return run(fmt.Sprintf("%s -chdir=%s apply %s/.terraform/tfplan", softwareMap[Terraform], config.WorkingDir, config.WorkingDir))
}

func tfInit() error {
    return run(fmt.Sprintf("%s -chdir=%s init", softwareMap[Terraform], config.WorkingDir))
}

func tfPlan() error {
    return run(fmt.Sprintf("%s -chdir=%s plan -out %s/.terraform/tfplan", softwareMap[Terraform], config.WorkingDir, config.WorkingDir))
}

func commandExists(cmd string) bool {
    _, err := exec.LookPath(cmd)
    return err == nil
}

func ovpnStart(configPath string, stdin bool) error {
    cmd := ovpnConnect(configPath)
    if stdin {
        cmd.Stdin = os.Stdin
    }
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    fmt.Println("Opening VPN tunnel...")
    if err := cmd.Start(); err != nil {
        return err
    }

    var waiting = true

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        os.Interrupt,
        os.Kill)

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
    if err != nil { return err }

    return nil
}
