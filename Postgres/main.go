package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "product"
)

func dbConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() string {
	dat, err := ioutil.ReadFile("quotes_all.csv")
	check(err)
	return string(dat)
}

func showAllTables(db *sql.DB) *sql.Rows {
	Result, err := db.Query("SELECT * FROM Field")
	if err != nil {
		panic(err)
	}
	return Result
}

func insertTable(db *sql.DB, field1 string, field2, field3 string) {
	Insert := fmt.Sprintf("INSERT INTO Field (f1, f2, f3) VALUES (%s, %s, %s);", field1, field2, field3)
	_, err := db.Exec(Insert)

	if err != nil {
		panic(err)
	}
}
func makeTables(db *sql.DB) {
	SQL := "CREATE TABLE Field ( Field1 TEXT, Field2 TEXT, Field3 TEXT);"
	_, err := db.Exec(SQL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created table")
}
func main() {
	db := dbConnect()
	//makeTables(db)
	defer db.Close()
	insertTable(db, "hello", "hello", "hello")
	//Rows := showAllTables(db)
	//data := readFile()
	//fmt.Println(data)

}
