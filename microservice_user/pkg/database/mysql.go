package database

import (
	"database/sql"
	"log"

	"microservice_user/config"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL faz a conexão com BD e retorna uma pool de conexão
func MySQL(conf *config.Config) *dabase_pool {
	// Verifica se já tem uma pool de conexão (precisa só de uma), caso tiver apenas retorna
	if dbpool != nil && dbpool.DB != nil {
		return dbpool
	} else {
		// Abre conexão com BD tendo conf do drive e DSN do banco
		db, err := sql.Open(conf.DB_DRIVE, conf.DB_DSN)
		// Verifica se teve algum erro
		if err != nil {
			log.Fatal(err)
		}

		// Testa conexão com o BD, dando um Ping
		err = db.Ping()
		// Verifica se teve algum erro
		if err != nil {
			log.Fatal(err)
		}

		// Atribui a conexão a variável pool de conexão
		dbpool = &dabase_pool{
			DB: db,
		}
	}

	// retorna pool de conexão
	return dbpool
}
