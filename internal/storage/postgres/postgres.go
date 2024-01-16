package postgres

import (
	"database/sql"
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"github.com/sirupsen/logrus"
)

type linkSearcher struct {
	id       int
	old_link string
	new_link string
}

type Database struct {
	db *sql.DB
}

// gets full value
func (d *Database) GetValue(new string) (string, error) {
	link := linkSearcher{}
	query := `SELECT *
FROM links 
WHERE new_link = $1;`
	rows, err := d.db.Query(query, new)
	if err != nil {
		logrus.Errorf("%v", err)
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&link.id, &link.old_link, &link.new_link)
	}
	if err != nil {
		logrus.Errorf("%v", err)
		return "", err
	}
	return link.old_link, nil
}
func (d *Database) LookUp(old string) (bool, error) {
	link := linkSearcher{}
	query := `SELECT * 
FROM links 
where old_link = $1`
	rows, err := d.db.Query(query, old)
	if err != nil {
		logrus.Errorf("%v", err)
		return false, err
	}
	for rows.Next() {
		err = rows.Scan(&link.id, &link.old_link, &link.new_link)
	}
	return link.old_link == old, nil
}

func (d *Database) WriteValue(short string, full string) error {
	query := `INSERT INTO links(old_link,new_link) 
VALUES($1,$2)`
	_, err := d.db.Query(query, full, short)
	if err != nil {
		logrus.Error("%v: %v", storage.ErrNotWritten, err)
		return err
	}
	return nil
}

func Create(db *sql.DB) *Database {
	return &Database{db: db}
}
