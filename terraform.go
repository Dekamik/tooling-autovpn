package main

import (
    "context"
    "github.com/hashicorp/terraform-exec/tfexec"
    "github.com/hashicorp/terraform-exec/tfinstall"
    "io/ioutil"
    "os"
)

func tfApply() error {
    tf := getTerraform()

    initErr := tf.Init(context.Background(), tfexec.Upgrade(true))
    if initErr != nil {
        return initErr
    }

    applyErr := tf.Apply(context.Background())
    if applyErr != nil {
        return applyErr
    }

    return nil
}

func tfPlan() {
    // Use normal exec
}

func getTerraform() *tfexec.Terraform {
    tmpDir, err := ioutil.TempDir("", "tfinstall")
    if err != nil {
        panic(err)
    }
    defer os.RemoveAll(tmpDir)

    execPath, err := tfinstall.Find(context.Background(), tfinstall.LatestVersion(tmpDir, false))
    if err != nil {
        panic(err)
    }

    homeDir, _ := os.UserHomeDir()
    workingDir := homeDir + "/.autovpn/"
    tf, err := tfexec.NewTerraform(workingDir, execPath)
    if err != nil {
        panic(err)
    }

    return tf
}
