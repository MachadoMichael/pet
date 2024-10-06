package use_case

import (
	"testing"
)

func TestCreate(t *testing.T) {
	// Create a test database
	db, err := repository.NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Create a new pet DTO
	photos := [5]string{
		"photo1",
		"photo2",
		"photo3",
		"photo4",
		"photo5",
	}

	for _, p := range photos {
		_, err = CreatePhoto(p)
		if err != nil {
			t.Fatal(err)
		}
	}
	// Call the Create function

	// Clean up the test database
	err = db.DropTestDB()
	if err != nil {
		t.Fatal(err)
	}
}
