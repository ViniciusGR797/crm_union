package database

import (
	"database/sql"
	"fmt"
	"microservice_customer/config"
)

var (
	db *sql.DB
)

// DatabaseInterface Estrutura interface para padronizar comportamento de Database (tudo que tiver retorna DB e chega conexão DB é um Database)
type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

// dabase_pool Estrutura para pool de conexão no Database (reutilizar uma conexão do Database)
type dabase_pool struct {
	DB *sql.DB
}

// Cria variável que armazena o endereço da pool de conexão
var dbpool = &dabase_pool{}

// NewDB Cria nova conexão com Database, de acordo com as config passadas por parâmetro
func NewDB(conf *config.Config) *dabase_pool {
	conf.DBConfig.DB_DSN = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.DBConfig.DB_USER, conf.DBConfig.DB_PASS, conf.DBConfig.DB_HOST, conf.DBConfig.DB_PORT, conf.DBConfig.DB_NAME)
	dbpool = Mysql(conf)

	// retorna pool de conexão
	return dbpool
}

// Close Método que fecha conexão com Database
func (d *dabase_pool) Close() error {

	// Chama função para fechar a conexão que retorna feedback (erro ou não)
	err := d.DB.Close()
	if err != nil {
		return err
	}

	dbpool = &dabase_pool{}

	// retorna null, apenas por retornar já que deu certo a conexão com o Database
	return err
}

// GetDB retorna o Database conectado (em uso)
func (d *dabase_pool) GetDB() (DB *sql.DB) {
	return d.DB
}
