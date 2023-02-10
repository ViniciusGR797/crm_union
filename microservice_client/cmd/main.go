package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	// Import interno de packages do próprio sistema
	"microservice_client/config"
	"microservice_client/pkg/database"
	"microservice_client/pkg/routes"
	"microservice_client/pkg/server"
	"microservice_client/pkg/service"
)

// Função principal (primeira executada) - chama config para fazer conexão BD, service, server, router e roda servidor http
func main() {
	// Atribui o endereço da estrutura de uma configuração padrão do sistema
	default_conf := &config.Config{}

	// Abre o arquivo JSON com as variáveis de ambiente
	file, err := os.Open("microservice_client/env.json") // file.json has the json content
	if err != nil {
		log.Print(err)
	}

	// Lé todo JSON e transforma em um JSON byte
	jsonByte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
	}

	// Converte JSON byte em uma struct, no caso a struct default_conf
	if err := json.Unmarshal(jsonByte, &default_conf); err != nil {
		log.Print(err)
	}

	// Atribui para conf as novas configurações do sistema
	conf := config.NewConfig(default_conf)

	// Pega pool de conexão do Database (config passadas anteriorente pelo Json para Database)
	dbpool := database.NewDB(conf)
	// Se criou uma conexão com as configurações passadas para o Database - Imprima essa mensagem de sucesso
	if dbpool != nil {
		log.Print("Successfully connected")
	}

	// Cria serviços de um client (CRUD) com a pool de conexão passada por parâmetro
	service := service.NewClientService(dbpool)

	// Cria servidor HTTP com as config passadas por parâmetro
	serv := server.NewServer(conf)

	// Cria rotas passsando o servidor HTTP e os serviços do user (CRUD)
	router := routes.ConfigRoutes(serv.SERVER, service)

	// Se tiver ativada a interface de usuário, criar as rotas para o front end (WEB UI)
	// if conf.WEB_UI {
	// 	webui.RegisterUIHandlers(router)
	// }

	// Coloca servidor para rodar passando as rotas, servidor HTTP e serviços do user (CRUD) como parâmetro
	server.Run(router, serv, service)
}
