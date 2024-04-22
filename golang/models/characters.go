package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Character struct {
	IdCharacter      int `json:"idCharacter,omitempty"`
	Created          int `json:"created"`
	Name             int `json:"name"`
	CharacterPicture int `json:"characterPicture"`
}

func GetAllCharacters() ([]Character, error) {
	rows, err := DB.Query("SELECT id_char, created, name, character_picture FROM characters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	characters := []Character{}

	for rows.Next() {
		var ch Character
		err = rows.Scan(&ch.IdCharacter, &ch.Created, &ch.Name, &ch.CharacterPicture)
		if err != nil {
			return nil, err // return partial data and error if needed
		}
		characters = append(characters, ch)
	}

	// check for issues after we're done iterating over the result set.
	if err = rows.Err(); err != nil {
		return characters, err
	}

	return characters, nil
}

func GetCharacter(idCharacter int) (Character, error) {
	row := DB.QueryRow("SELECT id_char, created, name, character_picture FROM characters WHERE id_chat = ?", idCharacter)

	var ch Character
	ch.IdCharacter = idCharacter
	err := row.Scan(&ch.IdCharacter, &ch.Created, &ch.Name, &ch.CharacterPicture)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return ch, err
	case err != nil:
		log.Fatal(err)
		return ch, err
	default:
		return ch, nil
	}
}

func CreateCharacter(ch Character) error {
	_, err := DB.Exec("INSERT INTO characters (id_char, created, name, character_picture) VALUES (?, ?)", ch.IdCharacter, ch.Created, ch.Name, ch.CharacterPicture)
	return err
}
