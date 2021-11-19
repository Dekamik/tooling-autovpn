package main

import (
    "bufio"
    "io/fs"
    "os"
    "path/filepath"
    "text/template"
)

var tfMainTemplate = `terraform {
  required_providers {
    linode = {
      source = "linode/linode"
      version = "1.16.0"
    }
  }
}

provider "linode" {
  token = "{{.Token}}"
}
`

var tfVpnTemplate = `module "{{.Name}}" {
  source = "git@github.com:Dekamik/vpn-modules.git//vpn-server?ref=v0.2.3"

  token = "{{.Token}}"
  public_keys = {
    "{{.Hostname}}" = "{{.PublicKey}}"
  }

  name = "{{.Hostname}}-{{.Name}}"
  region = "{{.Region}}"
  type = "{{.Type}}"
  download_dir = "{{.DownloadDir}}"
}
`

var templates = map[string]*template.Template {
    "main": template.Must(template.New("main").Parse(tfMainTemplate)),
    "vpn": template.Must(template.New("vpn").Parse(tfVpnTemplate)),
}

type TemplateReceiver struct {
    FilePath string
    TemplateName string
    TemplateArgs interface{}
}

func writeFile(receiver TemplateReceiver) (string, error) {
    file, err := os.Create(receiver.FilePath)
    if err != nil { return receiver.FilePath, err }
    writer := bufio.NewWriter(file)

    err = templates[receiver.TemplateName].Execute(writer, receiver.TemplateArgs)
    if err != nil { return receiver.FilePath, err }

    err = writer.Flush()
    if err != nil { return receiver.FilePath, err }

    return receiver.FilePath, nil
}

func removeFiles(files []string) (int, error) {
    var filesRemoved = 0
    summary := make([][]string, len(files))

    for i, file := range files {
        if _, err := os.Stat(file); err == nil {
            err = os.Remove(file)
            if err != nil {
                summary[i] = []string { file, "Error" }
                return filesRemoved, err
            }
            summary[i] = []string { file, "Removed" }
            filesRemoved++
        } else {
            summary[i] = []string { file, "Not found" }
        }
    }

    return filesRemoved, nil
}

func find(root, ext string) []string {
    var a []string
    _ = filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
        if e != nil {
            return e
        }
        if filepath.Ext(d.Name()) == ext {
            a = append(a, s)
        }
        return nil
    })
    return a
}
