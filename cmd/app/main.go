package main

import (
	"fmt"
	"go-task-cli/internal/app"
	"go-task-cli/internal/config"
	"go-task-cli/internal/repository"
	"go-task-cli/internal/service"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		fmt.Printf("Ошибка инициализации конфига: %v\n", err)
		return
	}
	repo := repository.NewTaskRepository(config.TaskFile)
	serv := service.NewTaskService(repo)

	app.Run(serv)
}
