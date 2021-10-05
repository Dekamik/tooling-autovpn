package main

import (
    "bufio"
    "os"
    "strings"
)

func create(token string) error {
    var instances []Instance
    hostName, _ := os.Hostname()
    homeDir, _ := os.UserHomeDir()

    sshFile, openErr := os.Open(homeDir + "/.ssh/id_ed25519.pub")
    if openErr != nil { return openErr }
    sshReader := bufio.NewReader(sshFile)
    publicKey, readErr := sshReader.ReadString('\n')
    if readErr != nil { return readErr }
    publicKey = strings.TrimSuffix(publicKey, "\n")

    for _, region := range options.Regions {
        instances = append(instances, Instance{
            Name:      region,
            Hostname:  hostName,
            Token:     token,
            PublicKey: publicKey,
            Region:    region,
            Type:      "g6-nanode-1",
        })
    }

    for _, instance := range instances {
        createErr := createFile(instance)
        if createErr != nil { return createErr }
    }

    return nil
}
