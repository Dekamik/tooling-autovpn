package main

import (
    "fmt"
    "os"
)

func tfApply() error {
    homeDir, _ := os.UserHomeDir()
    err := run(fmt.Sprintf("terraform -chdir=%s/.autovpn/ apply %s/.autovpn/.terraform/tfplan", homeDir, homeDir))
    return err
}

func tfInit() error {
    homeDir, _ := os.UserHomeDir()
    err := run(fmt.Sprintf("terraform -chdir=%s/.autovpn/ init", homeDir))
    return err
}

func tfPlan() error {
    homeDir, _ := os.UserHomeDir()
    err := run(fmt.Sprintf("terraform -chdir=%s/.autovpn/ plan -out %s/.autovpn/.terraform/tfplan", homeDir, homeDir))
    return err
}
