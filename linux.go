// +build linux

package main

import (
    "errors"
    "fmt"
    "os/exec"
)

var softwareMap = map[Software]string {
    Ansible: "ansible-playbook",
    OpenSSH: "ssh",
    OpenVPN: "openvpn",
    Terraform: "terraform",
}

func ovpnConnect(configPath string) *exec.Cmd {
    return exec.Command("sudo", softwareMap[OpenVPN], configPath)
}

func checkPrerequisites() error {
    for program, executable := range softwareMap {
        if !commandExists(executable) {
            return errors.New(fmt.Sprintf("%s not found on system, is %s installed?", executable, program))
        }
    }
    return nil
}
