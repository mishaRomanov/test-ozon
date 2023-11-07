package postgres

import "database/sql"

type Database struct {
	db *sql.DB
}

func (d *Database) GetValue(string) (string, error) {
	return "", nil
}
func (d *Database) LookUp(string) (bool, error) {
	return true, nil
}
func (d *Database) WriteValue(string, string) error {
	return nil
}

func NewDatabase() *Database {
	var db *Database
	db.db.Exec(`CREATE TABLE links()`)
	return db
}
