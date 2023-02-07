package config

import (
	// Package "os" - usado para manipulação de arquivos

	"fmt"
	"os"

	// Package "strconv" - implementa conversões de tipos primitivos, ex: String para int
	"strconv"

	"github.com/joho/godotenv"
)

const (
	// DEVELOPER - Modo de desenvolvimento - Ambiente programável
	DEVELOPER = "developer"
	// PRODUCTION - Modo de produção - Usuário final
	PRODUCTION = "production"
)

// Estrutura para armazenar as configurações da aplicação - Config
type Config struct {
	// Porta do servidor - Ex: 8080
	USER_PORT int `json:"user_port"`
	// Modo de uso da API - DEVELOPER, HOMOLOGATION ou PRODUCTION
	Mode string `json:"mode"`
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
func NewConfig() *Config {
	// Variável que armazenará as novas configurações
	var conf *Config

	var err error

	if err = godotenv.Load(); err != nil {
		conf = DefaultConfig()
	}
	fmt.Print(err, " ||||||||||||")

	// Atribui uma variável de ambiente para porta do servidor
	USER_PORT, _ := strconv.Atoi(os.Getenv("USER_PORT"))
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if USER_PORT != 0 {
		conf.USER_PORT = USER_PORT
	}

	// Atribui uma variável de ambiente para modo de uso da API
	SRV_MODE := os.Getenv("SRV_MODE")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	// Atribui uma variável de ambiente para drive do DataBase
	DB_DRIVE := os.Getenv("DB_DRIVE")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if DB_DRIVE != "" {
		conf.DBConfig.DB_DRIVE = DB_DRIVE
	}

	// Atribui uma variável de ambiente para host do Database
	DB_HOST := os.Getenv("DB_HOST")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if DB_HOST != "" {
		conf.DBConfig.DB_HOST = DB_HOST
	}

	// Atribui uma variável de ambiente para porta do Database
	DB_PORT := os.Getenv("DB_PORT")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if DB_PORT != "" {
		conf.DBConfig.DB_PORT = DB_PORT
	}

	// Atribui uma variável de ambiente para usuário do Database
	DB_USER := os.Getenv("DB_USER")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if DB_USER != "" {
		conf.DBConfig.DB_USER = DB_USER
	}

	// Atribui uma variável de ambiente para senha do Database
	DB_PASS := os.Getenv("DB_PASS")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if DB_PASS != "" {
		conf.DBConfig.DB_PASS = DB_PASS
	}

	// Atribui uma variável de ambiente para nome do Database
	DB_NAME := os.Getenv("DB_NAME")
	// Caso tenha essa variável de ambiente (não esteja vazia), atribui as novas configurações
	if DB_NAME != "" {
		conf.DBConfig.DB_NAME = DB_NAME
	}

	// Retorna a nova configuração
	return conf
}

// Configurações padrão da aplicação - defaultConf
func DefaultConfig() *Config {
	// Cria e atribui já valores para a configuração padrão
	default_config := Config{
		USER_PORT: 8081,
		DBConfig: DBConfig{
			DB_DRIVE: "mysql",
			DB_HOST:  "localhost",
			DB_PORT:  "3306",
			DB_USER:  "",
			DB_PASS:  "",
			DB_NAME:  "",
		},
		Mode: PRODUCTION,
	}

	// retorna o endereço de memória da configuração padrão (aumentar eficiência evitando copias)
	return &default_config
}
