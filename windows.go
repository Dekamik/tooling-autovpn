// +build windows

package main

import "os/exec"

var softwareMap = map[Software]string {
    Ansible: "ansible-playbook",
    OpenSSH: "ssh",
    OpenVPN: "C:\\Program Files\\OpenVPN\\bin\\openvpn-gui.exe",
    Terraform: "terraform",
}

func ovpnConnect(configPath string) *exec.Cmd {
    return exec.Command(softwareMap[OpenVPN], "--command", "connect", configPath)
}

func ovpnDisconnect(configPath string) *exec.Cmd {
    return exec.Command(softwareMap[OpenVPN], "--command", "disconnect", configPath)
}

func checkPrerequisites() error {
    for program, executable := range softwareMap {
        if !commandExists(executable) {
            return errors.New(fmt.Sprintf("%s not found on system, is %s installed?", executable, program))
        }
    }
    return nil
}
