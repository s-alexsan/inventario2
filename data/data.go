package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var error error

	db, error = sql.Open("sqlite3", "./inventario2.db")
	if error != nil {
		return error
	}

	return db.Ping()
}

func CrateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS inventario (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"produto" TEXT,
		"categoria" TEXT,
		"preco" REAL,
		"quantidade" REAL
	)`

	statemend, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	statemend.Exec()
	log.Println("inventario criado com sucesso!")
}

func Query(sql string, args ...any) {
	statement, err := db.Prepare(sql)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(args...)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Query executado com sucesso")

}

func Select(sql string) {
	row, err := db.Query(sql)
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var id int
		var produto string
		var categoria string
		var preco float64
		var quantidade float64

		row.Scan(&id, &produto, &categoria, &preco, &quantidade)
		log.Println("[", id, "] ", produto, " - ", categoria, "(R$", preco, ", ", quantidade, ")")
	}

}
