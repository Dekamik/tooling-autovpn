package main

func tfApply() error {
    _ = run("dpush ~/.autovpn/")
    err := run("terraform apply ~/.autovpn/.terraform/tfplan")
    _ = run("dpop")
    return err
}

func tfInit() error {
    _ = run("dpush ~/.autovpn/")
    err := run("terraform init")
    _ = run("dpop")
    return err
}

func tfPlan() error {
    _ = run("dpush ~/.autovpn/")
    err := run("terraform plan -out ~/.autovpn/.terraform/tfplan")
    _ = run("dpop")
    return err
}
