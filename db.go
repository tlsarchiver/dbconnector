package dbconnector

import (
	"database/sql"
	"fmt"
	// This import is blank to use the PostgreSQL driver
	_ "github.com/lib/pq"
	"os"
)

// SetupDB initializes the DB
func SetupDB(dbConfig DatabaseConfig) *sql.DB {
	var err error

	// Open the database, being sure it has been installed
	db, err := sql.Open(dbConfig.dbType, dbConfig.buildDbURL())
	checkErr(err)

	// defer db.Close()

	// Limit the number of open connections to the database, even if this means we
	// need to wait for an existing transaction to finish and free its connection.
	db.SetMaxOpenConns(dbConfig.maxOpenConns)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open a connection to the database!")
		panic(err.Error())
	}

	return db
}
