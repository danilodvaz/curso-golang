package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Executa o select
	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5)
	defer rows.Close()

	fmt.Println(rows)

	// Pecorre tudo o que foi retornado
	for rows.Next() {
		var u usuario

		// Escaneia o resultado do select
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
