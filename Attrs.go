package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type AttrDB struct {
	db *sql.DB
}

func InitAttrDB(dbPath string) AttrDB {
	database, _ := sql.Open("sqlite3", dbPath)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS boxes (id TEXT PRIMARY KEY, icon TEXT, name TEXT, exec TEXT)")
	statement.Exec()
	return AttrDB{database}
}

func (attrDB *AttrDB) AddBox(id string, icon string, name string, exec string) {
	statement, _ := attrDB.db.Prepare("INSERT INTO boxes (id, icon, name, exec) VALUES (?, ?, ?, ?)")
	statement.Exec(id, icon, name, exec)
}

type Attrs struct {
	Id string
	Icon string
	Name string
	Exec string
}

func (attrDB *AttrDB) GetBoxAttrs(id string) (Attrs, error) {
	rows, _ := attrDB.db.Query("SELECT id, icon, name, exec FROM boxes WHERE id=?", id)
	if rows.Next() {
		var a Attrs
		rows.Scan(&a.Id, &a.Icon, &a.Name, &a.Exec)
		return a, nil
	} else {
		return Attrs{}, errors.New("id not found")
	}
}

func (attrDB *AttrDB) GetAllBoxes() []Attrs {
	rows, _ := attrDB.db.Query("SELECT * FROM boxes")
	attrsList := make([]Attrs, 0)
	for rows.Next() {
		var a Attrs
		rows.Scan(&a.Id, &a.Icon, &a.Name, &a.Exec)
		attrsList = append(attrsList, a)
	}
	return attrsList
}

func (attrDB *AttrDB) SetBoxAttrs(id, icon, name, exec string) {
	statement, _ := attrDB.db.Prepare("UPDATE boxes SET icon=?, name=?, exec=? WHERE id=?")
	statement.Exec(icon, name, exec, id)
}

func (attrDB *AttrDB) SetIcon(id string, icon string) {
	statement, _ := attrDB.db.Prepare("UPDATE boxes SET icon=? WHERE id=?")
	statement.Exec(icon, id)
}

func (attrDB *AttrDB) SetName(id string, name string) {
	statement, _ := attrDB.db.Prepare("UPDATE boxes SET name=? WHERE id=?")
	statement.Exec(name, id)
}

func (attrDB *AttrDB) SetExec(id string, exec string) {
	statement, _ := attrDB.db.Prepare("UPDATE boxes SET exec=? WHERE id=?")
	statement.Exec(exec, id)
}

func (attrDB *AttrDB) DeleteBox(id string) {
	statement, _ := attrDB.db.Prepare("DELETE FROM boxes WHERE id=?")
	statement.Exec(id)
}

func (attrDB *AttrDB) Close() {
	attrDB.db.Close()
}