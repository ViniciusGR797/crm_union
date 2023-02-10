package main

import (
<<<<<<< HEAD
=======
	"encoding/json"
	"io/ioutil"
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
	"log"

	// Import interno de packages do próprio sistema
	"microservice_user/config"
	"microservice_user/pkg/database"
	"microservice_user/pkg/routes"
	"microservice_user/pkg/server"
	"microservice_user/pkg/service"
)

// Função principal (primeira executada) - chama config para fazer conexão BD, service, server, router e roda servidor http
func main() {
	// Atribui o endereço da estrutura de uma configuração padrão do sistema
<<<<<<< HEAD
	conf := &config.Config{}
=======
	default_conf := &config.Config{}

	// Abre o arquivo JSON com as variáveis de ambiente
	file, err := os.Open("microservice_user/env.json") // file.json has the json content
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

	// Converte JSON byte em uma struct, no caso a struct default_conf
	if err := json.Unmarshal(jsonByte, &default_conf); err != nil {
		log.Print(err)
	}
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f

	// Atribui para conf as novas configurações do sistema
	conf = config.NewConfig()

	// Pega pool de conexão do Database (config passadas anteriorente pelo Json para Database)
	dbpool := database.NewDB(conf)
	// Se criou uma conexão com as configurações passadas para o Database - Imprima essa mensagem de sucesso
	if dbpool != nil {
		log.Print("Successfully connected")
	}

	// Cria serviços de um user (CRUD) com a pool de conexão passada por parâmetro
	service := service.NewUserService(dbpool)

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
