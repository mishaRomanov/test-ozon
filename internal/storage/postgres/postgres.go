package storage

import "database/sql"

type Database struct {
	db *sql.DB
}

func Create() (*sql.DB, error) {
}

func (d *Database) WriteValue(short, full string) error {

}

func (d *Database) GetValue(val string) error {

}
