package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Inicia uma transação
	tx, _ := db.Begin()

	// Chama a preparação da query direto da transação
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(? ,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	// Simulando um erro com a chave duplicada
	_, err = stmt.Exec(2002, "Tiago")

	if err != nil {
		// Realiza o rollback da transação
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
