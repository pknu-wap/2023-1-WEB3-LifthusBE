package db

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
	"log"
	"os"
)

// ConncectToHusAuth returns hus_auth_db's ent client.
// you've got to close it with Close() in defer out of this function.
func ConnectToLifthusUser() (*ent.Client, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DB connection
	connectionPhrase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	client, err := ent.Open("mysql", connectionPhrase)
	if err != nil {
		log.Print("[F] opening connection to mysql failed: %w", err)
		return nil, err
	}

	// Running the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Print("[F] creating schema resources failed: %w", err)
		return nil, err
	}

	return client, nil
}
