package main

import (
    "bufio"
    "fmt"
    "html/template"
    "os"
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

func createFile(instance Instance) error {
    homeDir, _ := os.UserHomeDir()
    mkDirErr := os.MkdirAll(homeDir + "/.autovpn", 0777)
    check(mkDirErr)

    tmpl, tmplErr := template.New("tfmodule").Parse(tfTemplate)
    if tmplErr != nil { return tmplErr }

    filePath := fmt.Sprintf("%s/.autovpn/%s.tf", homeDir, instance.Name)
    fmt.Println(filePath)
    file, fileErr := os.Create(filePath)
    if fileErr != nil { return tmplErr }
    writer := bufio.NewWriter(file)

    execErr := tmpl.Execute(writer, instance)
    if execErr != nil { return execErr }
    flushErr := writer.Flush()
    if flushErr != nil { return flushErr }

    return nil
}
