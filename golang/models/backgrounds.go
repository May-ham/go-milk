package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Background struct {
	IdBackground int `json:"idBackground,omitempty"`
	Created      int `json:"created"`
	Name         int `json:"name"`
}

func GetAllBackgrounds() ([]Background, error) {
	rows, err := DB.Query("SELECT id_background, created, name, FROM backgrounds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	backgrounds := []Background{}

	for rows.Next() {
		var b Background
		err = rows.Scan(&b.IdBackground, &b.Created, &b.Name)
		if err != nil {
			return nil, err // return partial data and error if needed
		}
		backgrounds = append(backgrounds, b)
	}

	// check for issues after we're done iterating over the result set.
	if err = rows.Err(); err != nil {
		return backgrounds, err
	}

	return backgrounds, nil
}

func GetBackground(idBackground int) (Background, error) {
	row := DB.QueryRow("SELECT id_background, created, name, FROM backgrounds WHERE id_background = ?", idBackground)

	var b Background
	b.IdBackground = idBackground
	err := row.Scan(&b.IdBackground, &b.Created, &b.Name)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return b, err
	case err != nil:
		log.Fatal(err)
		return b, err
	default:
		return b, nil
	}
}

func CreateBackground(b Background) error {
	_, err := DB.Exec("INSERT INTO characters (id_char, created, name, character_picture) VALUES (?, ?)", b.IdBackground, b.Created, b.Name)
	return err
}
