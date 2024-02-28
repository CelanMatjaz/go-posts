package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/CelanMatjaz/go-posts/database"
)

func main() {
	database.SetupEnv()
    database.CreateConnection()

	const migrationsDir = "./migrations"

	dirEntries, err := os.ReadDir(migrationsDir)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

    fmt.Println("Starting migrations...")

	for _, e := range dirEntries {
		ext := filepath.Ext(e.Name())

		if ext == ".sql" {
			file, err := os.ReadFile(filepath.Join(migrationsDir, e.Name()))

			if err != nil {
				log.Fatal(err)
				os.Exit(-1)
			}

            fmt.Println("Running migration", e.Name())

            _, err = database.DB.Exec(string(file))

			if err != nil {
				log.Fatal(err)
				os.Exit(-1)
			}
		}
	}

    fmt.Println("Done")

    database.DestroyConnection()
}
