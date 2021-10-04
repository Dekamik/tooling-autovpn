package main

import (
    "io/ioutil"
    "os"
)

func create(token string) error {
    var instances []Instance
    hostName, hostNameErr := os.Hostname()
    if hostNameErr != nil { return hostNameErr }
    homeDir, _ := os.UserHomeDir()
    publicKey, readErr := ioutil.ReadFile(homeDir + "/.ssh/id_rsa.pub")
    if readErr != nil { return readErr }
    pubKeyContent := string(publicKey)

    for _, region := range options.Regions {
        instances = append(instances, Instance{
            Name:      region,
            Hostname:  hostName,
            Token:     token,
            PublicKey: pubKeyContent,
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
