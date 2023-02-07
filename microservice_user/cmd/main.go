package main

import (
	"encoding/json"
	"log"
	"os"

	// Import interno de packages do próprio sistema
	"crm/config"
	"crm/pkg/database"
	"crm/pkg/routes"
	"crm/pkg/server"
	"crm/pkg/service"
	"crm/webui"
)

// Função principal (primeira executada) - chama config para fazer conexão BD, service, server, router e roda servidor http
func main() {
	// Atribui o endereço da estrutura de uma configuração padrão do sistema
	default_conf := &config.Config{}

	// Atribui arquivo Json com configurações externas para file_config, verifica se não está vazia para virar novas configurações
	if file_config := "test_db.json"; file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	// Atribui para conf as novas configurações do sistema
	conf := config.NewConfig(default_conf)

	// Pega pool de conexão do Database (config passadas anteriorente pelo Json para Database)
	dbpool := database.NewDB(conf)
	// Se criou uma conexão com as configurações passadas para o Database - Imprima essa mensagem de sucesso
	if dbpool != nil {
		log.Print("Successfully connected")
	}

	// Cria serviços de um produto (CRUD) com a pool de conexão passada por parâmetro
	service := service.NewProdutoService(dbpool)

	// Cria servidor HTTP com as config passadas por parâmetro
	serv := server.NewServer(conf)

	// Cria rotas passsando o servidor HTTP e os serviços do produto (CRUD)
	router := routes.ConfigRoutes(serv.SERVER, service)

	// Se tiver ativada a interface de usuário, criar as rotas para o front end (WEB UI)
	if conf.WEB_UI {
		webui.RegisterUIHandlers(router)
	}

	// Coloca servidor para rodar passando as rotas, servidor HTTP e serviços do produto (CRUD) como parâmetro
	server.Run(router, serv, service)
}
