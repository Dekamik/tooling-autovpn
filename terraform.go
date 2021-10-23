package main

import (
    "fmt"
)

func tfApply() error {
    return run(fmt.Sprintf("terraform -chdir=%s apply %s/.terraform/tfplan", config.WorkingDir, config.WorkingDir))
}

func tfInit() error {
    return run(fmt.Sprintf("terraform -chdir=%s init", config.WorkingDir))
}

func tfPlan() error {
    return run(fmt.Sprintf("terraform -chdir=%s plan -out %s/.terraform/tfplan", config.WorkingDir, config.WorkingDir))
}
