package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MachadoMichael/pet/config"
	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	connStr := config.Variables.ConnStrDb
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		println("dbError", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migration(db *sql.DB) error {
	filePath := filepath.Join(config.Variables.MigrationsPath, "create_db_tables.up.sql")
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	sqlString := string(sqlBytes)

	for i, v := range strings.Split(sqlString, ";") {
		v = v + ";"
		_, err = db.Exec(v)
		if err != nil {
			fmt.Println("Migrations error: ", err, i, v)

			return fmt.Errorf("failed to execute SQL commands: %w", err)
		}
	}

	return nil
}
