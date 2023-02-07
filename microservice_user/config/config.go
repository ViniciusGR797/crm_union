package config

import (
	// Package "os" - usado para manipulação de arquivos
	"os"
	// Package "strconv" - implementa conversões de tipos primitivos, ex: String para int
	"strconv"
)

const (
	// DEVELOPER - Modo de desenvolvimento - Ambiente programável
	DEVELOPER = "developer"
	// HOMOLOGATION - Modo de homologação - Fase de testes
	HOMOLOGATION = "homologation"
	// PRODUCTION - Modo de produção - Usuário final
	PRODUCTION = "production"
)

// Estrutura para armazenar as configurações da aplicação - Config
type Config struct {
	// Porta do servidor - Ex: 8080
	SRV_PORT string `json:"srv_port"`
	// Web UI (Interface de Usuário Web) - Atividada/Desativada
	WEB_UI bool `json:"web_ui"`
	// Modo de uso da API - DEVELOPER, HOMOLOGATION ou PRODUCTION
	Mode string `json:"mode"`
	// Abrir o navegador - Atividada/Desativada
	OpenBrowser bool `json:"open_browser"`
	// Configurações do DataBase
	DBConfig `json:"dbconfig"`
}

// Estrutura para armazenar as configurações do banco de dados - DBConfig
type DBConfig struct {
	// Drive do DataBase - Ex: MySql
	DB_DRIVE string `json:"db_drive"`
	// Host do Database - Ex: LocalHost
	DB_HOST string `json:"db_host"`
	// Porta do Database - Ex: 3306
	DB_PORT string `json:"db_port"`
	// Usuário do Database - Ex: root
	DB_USER string `json:"db_user"`
	// Senha do Database - Ex: ******
	DB_PASS string `json:"db_pass"`
	// Nome do Database - Ex: golangdb
	DB_NAME string `json:"db_name"`
	// Data source name (Nome da Fonte de Dados) - Converter nome do site em IP - Ex: Google.com em 8.8.8.8
	DB_DSN string `json:"-"`
}

// NewConfig - Cria uma nova configuração - passada por parâmetro
func NewConfig(config *Config) *Config {

	// Variável que armazenará as novas configurações
	var conf *Config

	// Verifica se a config e a porta do servidor está vazia (caso estejá pega a config padrão)
	if config == nil || config.SRV_PORT == "" {
		conf = DefaultConfig()
	} else {
		conf = config
	}

	// Atribui uma variável de ambiente para porta do servidor
	SRV_PORT := os.Getenv("SRV_PORT")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	// Atribui uma variável de ambiente para modo de uso da API
	SRV_MODE := os.Getenv("SRV_MODE")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	// Atribui uma variável de ambiente para interface de Usuário Web
	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_WEB_UI != "" {
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	// Atribui uma variável de ambiente para drive do DataBase
	SRV_DB_DRIVE := os.Getenv("SRV_DB_DRIVE")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_DB_DRIVE != "" {
		conf.DBConfig.DB_DRIVE = SRV_DB_DRIVE
	}

	// Atribui uma variável de ambiente para host do Database
	SRV_DB_HOST := os.Getenv("SRV_DB_HOST")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_DB_HOST != "" {
		conf.DBConfig.DB_HOST = SRV_DB_HOST
	}

	// Atribui uma variável de ambiente para porta do Database
	SRV_DB_PORT := os.Getenv("SRV_DB_PORT")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_DB_PORT != "" {
		conf.DBConfig.DB_PORT = SRV_DB_PORT
	}

	// Atribui uma variável de ambiente para usuário do Database
	SRV_DB_USER := os.Getenv("SRV_DB_USER")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_DB_USER != "" {
		conf.DBConfig.DB_USER = SRV_DB_USER
	}

	// Atribui uma variável de ambiente para senha do Database
	SRV_DB_PASS := os.Getenv("SRV_DB_PASS")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_DB_PASS != "" {
		conf.DBConfig.DB_PASS = SRV_DB_PASS
	}

	// Atribui uma variável de ambiente para nome do Database
	SRV_DB_NAME := os.Getenv("SRV_DB_NAME")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_DB_NAME != "" {
		conf.DBConfig.DB_NAME = SRV_DB_NAME
	}

	// Retorna a nova configuração
	return config
}

// Configurações padrão da aplicação - defaultConf
func DefaultConfig() *Config {
	// Cria e atribui já valores para a configuração padrão
	default_config := Config{
		SRV_PORT:    "8080",
		WEB_UI:      true,
		OpenBrowser: true,
		DBConfig: DBConfig{
			DB_DRIVE: "mysql",
			DB_HOST:  "localhost",
			DB_PORT:  "3306",
			DB_USER:  "grupob",
			DB_PASS:  "Rapadura745",
			DB_NAME:  "golangdb",
		},
		Mode: PRODUCTION,
	}

	// retorna o endereço de memória da configuração padrão (aumentar eficiência evitando copias)
	return &default_config
}
