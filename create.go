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
    Name      string
    Hostname  string
    Token     string
    PublicKey string
    Region    string
    Type      string
}

func createFiles(instances []Instance) (int, error) {
    var createdFiles = 0
    summary := make([][]string, len(options.Regions) + 1)
    if options.Verbose {
        defer printTable(summary)
    }

    homeDir, _ := os.UserHomeDir()
    mkDirErr := os.MkdirAll(homeDir + "/.autovpn", 0777)
    check(mkDirErr)

    mainStruct := Main{Token: instances[0].Token}
    mainReceiver := TemplateReceiver{
        FilePath:     fmt.Sprintf("%s/.autovpn/main.tf", homeDir),
        TemplateName: "main",
        TemplateArgs: mainStruct,
    }
    mainFile, mainCreateErr := writeFile(mainReceiver)
    if mainCreateErr != nil {
        summary[0] = []string { mainFile, "Error" }
        return 0, mainCreateErr
    } else {
        summary[0] = []string { mainFile, "Created" }
        createdFiles++
    }

    for i, instance := range instances {
        args := TemplateReceiver{
            FilePath:     fmt.Sprintf("%s/.autovpn/%s.tf", homeDir, instance.Name),
            TemplateName: "vpn",
            TemplateArgs: instance,
        }
        fileName, createErr := writeFile(args)
        if createErr != nil {
            summary[i] = []string { fileName, "Error" }
            return createdFiles, createErr
        }
        summary[i] = []string { fileName, "Created" }
        createdFiles++
    }

    return createdFiles, nil
}

func create() error {
    var instances []Instance
    hostName, _ := os.Hostname()
    hostName = hostName[1:]
    homeDir, _ := os.UserHomeDir()

    sshFile, openErr := os.Open(homeDir + "/.ssh/id_rsa.pub")
    if openErr != nil { return openErr }
    sshReader := bufio.NewReader(sshFile)
    publicKey, readErr := sshReader.ReadString('\n')
    if readErr != nil { return readErr }
    publicKey = strings.TrimSuffix(publicKey, "\n")

    for _, region := range options.Regions {
        instances = append(instances, Instance{
            Name:      region,
            Hostname:  config.Hostname,
            Token:     config.Token,
            PublicKey: publicKey,
            Region:    region,
            Type:      "g6-nanode-1",
        })
    }

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
