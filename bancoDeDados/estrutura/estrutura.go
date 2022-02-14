package main

import (
	// Importa com o _, pois é um import que não será acessado diretamente.
	// Ele é necessário, mas não será chamado.
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Recebe uma refrência da conexão aberta com o banco de dados e a query que
// será executada. Retornando o resutado.
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	// Ao utilizar o nome "mysql" para abrir a conexão, automaticamente será
	// utilizado o driver que foi importado do mysql
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Cria o banco cursogo se ele não existir
	exec(db, "create database if not exists cursogo")
	// Define a utilização do banco cursogo
	exec(db, "use cursogo")
	// Drop a tabela de usuários, se ela existir
	exec(db, "drop table if exists usuarios")
	// Cria a tabela de usuários
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)

}
