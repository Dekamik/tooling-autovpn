package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func createFile(fileName string, templateName TemplateName, args interface{}) error {
    err := os.MkdirAll(config.WorkingDir, 0777)
    if err != nil { return err }

    receiver := TemplateReceiver{
        FilePath: fmt.Sprintf("%s/%s.tf", config.WorkingDir, fileName),
        TemplateName: templateName,
        TemplateArgs: args,
    }

    _, err = writeFile(receiver)
    if err != nil { return err }

    return nil
}

func createFiles(instances []LinodeInstanceArgs) (int, error) {
    var createdFiles = 0

    mainStruct := LinodeMainArgs{Token: instances[0].Token}
    err := createFile("main", LinodeMain, mainStruct)
    if err != nil { return 0, err }
    createdFiles++

    for _, instance := range instances {
        err = createFile(instance.Name, LinodeVpn, instance)
        if err != nil { return createdFiles, err }
        createdFiles++
    }

    return createdFiles, nil
}

func create() error {
    fmt.Println("Creating VPN server...")
    var instances []LinodeInstanceArgs

    sshFile, err := os.Open(config.SshPath)
    if err != nil { return err }
    sshReader := bufio.NewReader(sshFile)
    publicKey, err := sshReader.ReadString('\n')
    if err != nil { return err }
    publicKey = strings.TrimSuffix(publicKey, "\n")

    instances = append(instances, LinodeInstanceArgs{
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
