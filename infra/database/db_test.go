package database

import (
	"testing"

	"github.com/MachadoMichael/pet/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// Test successful initialization
	db, err := Init()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Test failed initialization (e.g. invalid connection string)
	invalidConnStr := "invalid://connection:string"
	config.Variables.ConnStrDb = invalidConnStr
	db, err = Init()
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestMigration(t *testing.T) {
	// Test successful migration
	db, err := Init()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = Migration(db)
	assert.NoError(t, err)

	// Test failed migration (e.g. invalid migration file)
	invalidMigrationFile := "invalid_migration_file.sql"
	config.Variables.MigrationsPath = invalidMigrationFile
	err = Migration(db)
	assert.Error(t, err)
}
