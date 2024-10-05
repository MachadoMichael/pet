package route

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStart(t *testing.T) {
	srv, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("server error: %v", err)
		}
	}()

	resp, err := http.Get("http://localhost:4000")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	if err := srv.Close(); err != nil {
		t.Errorf("server close error: %v", err)
	}
}
