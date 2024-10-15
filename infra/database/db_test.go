package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Open(database, connStr string) (*sql.DB, error) {
	args := m.Called(database, connStr)
	return args.Get(0).(*sql.DB), args.Error(1)
}

func (m *MockDB) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func TestOpenDB(t *testing.T) {
	// Create an instance of MockDB
	mockDB := new(MockDB)

	// Define expected arguments and return values
	database := "test_db"
	connStr := "test_connection_string"
	expectedDB := &sql.DB{} // This can be a real DB or a mock object, as needed

	// Set up expectations on the mock
	mockDB.On("Open", database, connStr).Return(expectedDB, nil)

	// Call the Open method
	db, err := mockDB.Open(database, connStr)

	// Assert that the return values are as expected
	assert.NoError(t, err)
	assert.Equal(t, expectedDB, db)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestPingDB(t *testing.T) {
	// Create an instance of MockDB
	mockDB := new(MockDB)

	// Set up expectations on the mock
	mockDB.On("Ping").Return(nil)

	// Call the Ping method
	err := mockDB.Ping()

	// Assert that the return value is as expected
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}
