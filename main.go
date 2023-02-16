package main

import "os/exec"

func main() {

	cmds := []string{
		"microservice_business/cmd/main.go",
		"microservice_client/cmd/main.go",
		"microservice_customer/cmd/main.go",
		"microservice_group/cmd/main.go",
		"microservice_planner/cmd/main.go",
		"microservice_release/cmd/main.go",
		"microservice_remark/cmd/main.go",
		"microservice_subject/cmd/main.go",
		"microservice_user/cmd/main.go",
	}

	for _, cmd := range cmds {
		command := exec.Command("go", "run", cmd)
		command.Start()
	}

}
