package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

const _layout = "20060102150405"

func main() {
	var command string
	var name string

	flag.StringVar(&command, "command", "", "migrate command (up|down|new|status|seed)")
	flag.StringVar(&name, "name", "", "name for new migration")
	flag.Parse()

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations", // パスを修正
	}

	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "up":
		n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d migrations\n", n)

	case "down":
		n, err := migrate.Exec(db, "mysql", migrations, migrate.Down)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rolled back %d migrations\n", n)

	case "new":
		if name == "" {
			log.Fatal("Migration name is required")
		}
		filename := fmt.Sprintf("%s_%s.sql", time.Now().Format(_layout), name)
		f, err := os.Create(fmt.Sprintf("db/migrations/%s", filename))
		if err != nil {
			log.Fatal("failed to create file:", err)
		}
		defer f.Close()

		f.WriteString("-- +migrate Up\n\n\n") //nolint:all
		f.WriteString("-- +migrate Down\n\n") //nolint:all
		fmt.Printf("Created new migration: %s\n", filename)

	case "status":
		records, err := migrate.GetMigrationRecords(db, "mysql")
		if err != nil {
			log.Fatal(err)
		}
		for _, record := range records {
			fmt.Printf("Applied: %s\n", record.Id)
		}
	case "seed":
		if err := seedData(db); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Seed data inserted successfully")

	default:
		log.Fatal("Invalid command. Use up, down, new, status, or seed")
	}
}

func setupDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	return sql.Open("mysql", dsn)
}

func seedData(db *sql.DB) error {
	seedFiles, err := filepath.Glob("db/seed/*.sql")
	if err != nil {
		return fmt.Errorf("failed to find seed files: %w", err)
	}

	for _, file := range seedFiles {
		if err := executeSeedFile(db, file); err != nil {
			return fmt.Errorf("failed to execute seed file %s: %w", file, err)
		}
		fmt.Printf("Executed seed file: %s\n", file)
	}
	return nil
}

func executeSeedFile(db *sql.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var queries []string
	var currentQuery strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// コメントと空行をスキップ
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}

		currentQuery.WriteString(line)
		currentQuery.WriteString(" ")

		// セミコロンで終わる行を1つのクエリとして扱う
		if strings.HasSuffix(line, ";") {
			queries = append(queries, strings.TrimSpace(currentQuery.String()))
			currentQuery.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// トランザクション内でクエリを実行
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			tx.Rollback() //nolint:all
			return fmt.Errorf("failed to execute query: %w", err)
		}
	}

	return tx.Commit()
}
