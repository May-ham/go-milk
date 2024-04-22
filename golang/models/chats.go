package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Chat struct {
	Id_chat int `json:"id_chat,omitempty"`
	Id_char int `json:"id_char"`
	Created int `json:"created"`
}

func GetAllChats() ([]Chat, error) {
	rows, err := DB.Query("SELECT id_chat, id_char, created FROM chats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats := []Chat{}

	for rows.Next() {
		var ch Chat
		err = rows.Scan(&ch.Id_chat, &ch.Id_char, &ch.Created)
		if err != nil {
			return nil, err // return partial data and error if needed
		}
		chats = append(chats, ch)
	}

	// check for issues after we're done iterating over the result set.
	if err = rows.Err(); err != nil {
		return chats, err
	}

	return chats, nil
}

func GetChat(id_chat int) (Chat, error) {
	row := DB.QueryRow("SELECT id_char,created FROM chats WHERE id_chat = ?", id_chat)

	var ch Chat
	ch.Id_chat = id_chat
	err := row.Scan(&ch.Id_char, &ch.Created)
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

func CreateChat(ch Chat) error {
	_, err := DB.Exec("INSERT INTO chats (id_char, created) VALUES (?, ?)", ch.Id_char, ch.Created)
	return err
}
