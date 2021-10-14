package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "text/template"
)

type Instance struct {
    Name      string
    Hostname  string
    Token     string
    PublicKey string
    Region    string
    Type      string
}

var tfTemplate = `module "{{.Name}}" {
  source = "git@github.com:Dekamik/vpn-modules.git//vpn-server?ref=v0.1.1"

  token = "{{.Token}}"
  public_keys = [
    "{{.PublicKey}}"
  ]

  name = "{{.Hostname}}-{{.Name}}"
  region = "{{.Region}}"
  type = "{{.Type}}"
}
`

func writeFile(instance Instance) (string, error) {
    homeDir, _ := os.UserHomeDir()
    mkDirErr := os.MkdirAll(homeDir + "/.autovpn", 0777)
    check(mkDirErr)

    tmpl, tmplErr := template.New("tfmodule").Parse(tfTemplate)
    if tmplErr != nil { return "", tmplErr }

    filePath := fmt.Sprintf("%s/.autovpn/%s.tf", homeDir, instance.Name)
    file, fileErr := os.Create(filePath)
    if fileErr != nil { return filePath, tmplErr }
    writer := bufio.NewWriter(file)

    execErr := tmpl.Execute(writer, instance)
    if execErr != nil { return filePath, execErr }
    flushErr := writer.Flush()
    if flushErr != nil { return filePath, flushErr }

    return filePath, nil
}

func createFiles(instances []Instance) (int, error) {
    var createdFiles = 0
    summary := make([][]string, len(options.Regions))
    if options.Verbose {
        defer printTable(summary)
    }

    for i, instance := range instances {
        fileName, createErr := writeFile(instance)
        if createErr != nil {
            summary[i] = []string { fileName, "Error" }
            return createdFiles, createErr
        }
        summary[i] = []string { fileName, "Created" }
        createdFiles++
    }

    return createdFiles, nil
}

func create(token string) error {
    var instances []Instance
    hostName, _ := os.Hostname()
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
            Hostname:  hostName,
            Token:     token,
            PublicKey: publicKey,
            Region:    region,
            Type:      "g6-nanode-1",
        })
    }

    createdFiles, createErr := createFiles(instances)
    if createErr != nil {
        return createErr
    }

    if createdFiles != 0 {
        err := tfApply()
        if err != nil {
            return err
        }
    }

    return nil
}
