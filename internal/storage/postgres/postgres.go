package storage

import "database/sql"

type Database struct {
	db *sql.DB
}

func (d *Database) GetValue(string) (string, error) {

}
func (d *Database) LookUp(string) (bool, error) {

}
func (d *Database) WriteValue(string, string) error {

}
