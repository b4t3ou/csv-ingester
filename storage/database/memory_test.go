package database

import (
	"testing"

	"github.com/b4t3ou/csv-ingester/storage"
)

func TestMemory_Save(t *testing.T) {
	memory := NewMemory()
	memory.Save(storage.Record{
		ID:           "1",
		Name:         "foo bar",
		Email:        "test@test.com",
		MobileNumber: "1235",
	})

	if _, exists := memory.data["test@test.com"]; !exists {
		t.Errorf("failed to save record")
	}
}
