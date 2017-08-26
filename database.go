package main

import (
	"database/sql"
	"log" // not required, but handy

	_ "github.com/lib/pq"
)

func main() {
	// lazy evaluation:
	db, err := sql.Open("postgres", "user=sarah password=123456 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Add a new record
	res, err := db.Exec("INSERT INTO CUSTOMER (FirstName, LastName ) VALUES ('Paul', 'Dechering');")
	if err != nil {
		panic(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	// use rowCount somehow, or replace with _
	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT ID, LastName, FirstName FROM CUSTOMER;")
	if err != nil {
		panic(err)
	}

	var (
		customerID int
		FirstName  string
		LastName   string
	)

	for rows.Next() {
		err = rows.Scan(&customerID, &FirstName, &LastName)
		if err != nil {
			panic(err)
		}
		log.Printf("Values: %d, %q, %q", customerID, FirstName, LastName)
	}
	rows.Close()
}
