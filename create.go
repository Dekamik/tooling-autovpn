package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Main struct {
    Token string
}

type Instance struct {
    Name        string
    Hostname    string
    Token       string
    PublicKey   string
    Region      string
    Type        string
    DownloadDir string
}

func createFiles(instances []Instance) (int, error) {
    var createdFiles = 0

    err := os.MkdirAll(config.WorkingDir, 0777)
    check(err)

    mainStruct := Main{Token: instances[0].Token}
    mainReceiver := TemplateReceiver{
        FilePath:     config.WorkingDir + "/main.tf",
        TemplateName: LinodeMain,
        TemplateArgs: mainStruct,
    }
    _, err = writeFile(mainReceiver)
    if err != nil {
        return 0, err
    } else {
        createdFiles++
    }

    for _, instance := range instances {
        args := TemplateReceiver{
            FilePath:     fmt.Sprintf("%s/%s.tf", config.WorkingDir, instance.Name),
            TemplateName: LinodeVpn,
            TemplateArgs: instance,
        }
        _, err := writeFile(args)
        if err != nil {
            return createdFiles, err
        }
        createdFiles++
    }

    return createdFiles, nil
}

func create() error {
    fmt.Println("Creating VPN server...")
    var instances []Instance

    sshFile, err := os.Open(config.SshPath)
    if err != nil { return err }
    sshReader := bufio.NewReader(sshFile)
    publicKey, err := sshReader.ReadString('\n')
    if err != nil { return err }
    publicKey = strings.TrimSuffix(publicKey, "\n")

    instances = append(instances, Instance{
        Name:        options.Region,
        Hostname:    config.Hostname,
        Token:       config.Token,
        PublicKey:   publicKey,
        Region:      options.Region,
        Type:        "g6-nanode-1",
        DownloadDir: config.WorkingDir,
    })

    createdFiles, err := createFiles(instances)
    if err != nil { return err }

    if createdFiles == 0 {
        return nil
    }

    err = tfInit()
    if err != nil { return err }

    err = tfPlan()
    if err != nil { return err }

    err = tfApply()
    if err != nil { return err }

    return nil
}
