package main

import (
    "bufio"
    "io/fs"
    "os"
    "path/filepath"
    "text/template"
)

type TemplateName int

const(
    Undefined TemplateName = iota
    LinodeMain
    LinodeVpn
)

var templates = map[TemplateName]*template.Template {
    LinodeMain: template.Must(template.New("main").Parse(linodeMainTemplate)),
    LinodeVpn: template.Must(template.New("vpn").Parse(linodeVpnTemplate)),
}

type TemplateReceiver struct {
    FilePath string
    TemplateName TemplateName
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
