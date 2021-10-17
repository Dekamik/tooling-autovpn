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
  source = "git@github.com:Dekamik/vpn-modules.git//vpn-server?ref=v0.2.1"

  token = "{{.Token}}"
  public_keys = {
    "{{.Hostname}}" = "{{.PublicKey}}"
  }

  name = "{{.Hostname}}-{{.Name}}"
  region = "{{.Region}}"
  type = "{{.Type}}"
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
    file, fileErr := os.Create(receiver.FilePath)
    if fileErr != nil { return receiver.FilePath, fileErr }
    writer := bufio.NewWriter(file)

    writeErr := templates[receiver.TemplateName].Execute(writer, receiver.TemplateArgs)
    if writeErr != nil {
        return receiver.FilePath, writeErr
    }

    flushErr := writer.Flush()
    if flushErr != nil {
        return receiver.FilePath, flushErr
    }

    return receiver.FilePath, nil
}

func removeFiles(files []string) (int, error) {
    var filesRemoved = 0
    summary := make([][]string, len(files))
    if options.Verbose {
        defer printTable(summary)
    }

    for i, file := range files {
        if _, err := os.Stat(file); err == nil {
            removeErr := os.Remove(file)
            if removeErr != nil {
                summary[i] = []string { file, "Error" }
                return filesRemoved, removeErr
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
