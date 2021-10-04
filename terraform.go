package main

import (
    "bufio"
    "fmt"
    "html/template"
    "os"
)

type Instance struct {
    Name  	  string
    HostName  string
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

func createFile(name string, instance Instance) error {
    tmpl, tmplErr := template.New("tfmodule").Parse(tfTemplate)
    if tmplErr != nil { return tmplErr }

    file, fileErr := os.Create(fmt.Sprintf("~/.autovpn/%s.tf", name))
    if fileErr != nil { return tmplErr }
    writer := bufio.NewWriter(file)

    execErr := tmpl.Execute(writer, instance)
    if execErr != nil { return execErr }

    return nil
}
