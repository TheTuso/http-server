package api

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
)

type Database struct {
	*sql.DB
}

func NewDatabase() *Database {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		panic(err)
	}

	return &Database{db}
}

func (database *Database) Disconnect() error {
	if err := database.Close(); err != nil {
		return err
	}

	return nil
}

func (database *Database) GetItems() ([]Item, error) {
	var databaseItems []Item
	items, err := database.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}

	for items.Next() {
		var item Item
		err = items.Scan(&item.ID, &item.Name)
		if err != nil {
			return nil, err
		}
		databaseItems = append(databaseItems, item)
	}

	return databaseItems, nil
}

func (database *Database) AddItem(item Item) error {
	_, err := database.Query("INSERT INTO items VALUES (?, ?)", item.ID, item.Name)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) RemoveItem(id string) error {
	if _, err := database.Query("DELETE FROM items WHERE id = ?", id); err != nil {
		return err
	}

	return nil
}
