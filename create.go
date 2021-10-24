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

    mkDirErr := os.MkdirAll(config.WorkingDir, 0777)
    check(mkDirErr)

    mainStruct := Main{Token: instances[0].Token}
    mainReceiver := TemplateReceiver{
        FilePath:     config.WorkingDir + "/main.tf",
        TemplateName: "main",
        TemplateArgs: mainStruct,
    }
    _, mainCreateErr := writeFile(mainReceiver)
    if mainCreateErr != nil {
        return 0, mainCreateErr
    } else {
        createdFiles++
    }

    for _, instance := range instances {
        args := TemplateReceiver{
            FilePath:     fmt.Sprintf("%s/%s.tf", config.WorkingDir, instance.Name),
            TemplateName: "vpn",
            TemplateArgs: instance,
        }
        _, createErr := writeFile(args)
        if createErr != nil {
            return createdFiles, createErr
        }
        createdFiles++
    }

    return createdFiles, nil
}

func create() error {
    fmt.Println("Creating VPN server...")
    var instances []Instance

    sshFile, openErr := os.Open(config.SshPath)
    if openErr != nil { return openErr }
    sshReader := bufio.NewReader(sshFile)
    publicKey, readErr := sshReader.ReadString('\n')
    if readErr != nil { return readErr }
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

    createdFiles, createErr := createFiles(instances)
    if createErr != nil {
        return createErr
    }

    if createdFiles == 0 {
        return nil
    }

    initErr := tfInit()
    if initErr != nil {
        return initErr
    }

    planErr := tfPlan()
    if planErr != nil {
        return planErr
    }

    applyErr := tfApply()
    if applyErr != nil {
        return applyErr
    }

    return nil
}
