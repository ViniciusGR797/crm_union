package database

import (
	"database/sql"
	"fmt"

	// Import interno de packages do próprio sistema
	"microservice_user/config"
)

// Cria ponteiro como variável global - armazena Database
var (
	db *sql.DB
)

// DatabaseInterface padroniza comportamento de Database (tudo que tiver retorna DB e chega conexão DB é um Database)
type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

// Estrutura para pool de conexão no Database (reutilizar uma conexão do Database)
type dabase_pool struct {
	DB *sql.DB
}

// Cria variável que armazena o endereço da pool de conexão
var dbpool = &dabase_pool{}

// NewDB cria nova conexão com Database, de acordo com as config passadas por parâmetro
func NewDB(conf *config.Config) *dabase_pool {
	// Atribui endereço DNS do Database passando URL do Database
	conf.DBConfig.DB_DSN = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.DBConfig.DB_USER, conf.DBConfig.DB_PASS, conf.DBConfig.DB_HOST, conf.DBConfig.DB_PORT, conf.DBConfig.DB_NAME)
	// Cria pool de conexão com Database, através das config
	dbpool = MySQL(conf)

	// retorna pool de conexão
	return dbpool
}

// Close fecha conexão com Database
func (d *dabase_pool) Close() error {

	// Chama função para fechar a conexão que retorna feedback (erro ou não)
	err := d.DB.Close()
	// Verifica se tem erro (não está vazio) - retorna a mensagem de erro
	if err != nil {
		return err
	}

	// Atribui o endereço dessa pool de conexão para variável global (para ser reutilizada)
	dbpool = &dabase_pool{}

	// retorna null, apenas por retornar já que deu certo a conexão com o Database
	return err
}

// GetDB pega/retorna o Database conectado (em uso)
func (d *dabase_pool) GetDB() (DB *sql.DB) {
	return d.DB
}
