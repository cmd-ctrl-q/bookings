package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// constants that are used to define properties connection pool
const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

// ConnectSQL creates database pool for postgres
func ConnectSQL(dsn string) (*DB, error) {

	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxIdleConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	// test db again
	err = testDB(d)
	if err != nil {
		panic(err)
	}

	return dbConn, nil
}

// testDB tries to ping the database
func testDB(d *sql.DB) error {

	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
