# Go DB connection module

This module helps establishing a connection to the database, based on the
environment variables.

Quick example:

	// Populates the db configuration from the environment variables
	dbConfig := dbconnector.ParseConfiguration()

	// Setup the DB
	db := dbconnector.SetupDB(dbConfig)

	// Prepare the requests we will be using
	stmtAddFail, err = db.Prepare("INSERT INTO certificates (host, ip, failed, failure_error, timestamp) VALUES ($1, $2, $3, $4, $5)")
	checkErr(err)
