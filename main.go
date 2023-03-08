package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmds := []string{
		"microservice_business",
		"microservice_client",
		"microservice_customer",
		"microservice_group",
		"microservice_planner",
		"microservice_release",
		"microservice_remark",
		"microservice_user",
		"microservice_subject",
	}

	for _, cmd := range cmds {
		command := exec.Command("go", "build", "-o", cmd, fmt.Sprintf("%s/cmd/main.go", cmd))
		err := command.Start()
		if err != nil {
			fmt.Println("Erro ao executar o comando:", err)
			return
		}
		err = command.Wait()
		if err != nil {
			fmt.Println("Erro ao esperar pelo término do processo:", err)
			return
		}
	}

}
