// +build windows

package main

import "os/exec"

func ovpnCommand(configPath string) *exec.Cmd {
    return exec.Command("openvpn", configPath)
}
