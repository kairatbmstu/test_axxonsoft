package database

import (
	"context"
	"database/sql"
	"time"
)

/*
Package database provides functions and utilities for working with a SQL database.
*/

/*
The DB variable represents the database connection pool. It is a pointer to the sql.DB struct, which provides a pool of database connections.
*/
var DB *sql.DB

/*
The NewDB function is used to create a new database connection pool.
It takes a dsn (Data Source Name) string as a parameter, which specifies
the details of the database connection. The function returns a pointer
to the sql.DB struct representing the connection pool and an error
if any occurred during the connection establishment.
*/
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
