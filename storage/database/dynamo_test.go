package database

import (
	"os"
	"testing"

	"github.com/satori/go.uuid"

	"github.com/b4t3ou/csv-ingester/storage"
)

func TestDynamo_Save_Get(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "foo")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "bar")

	db, err := NewDynamoDB(
		uuid.NewV4().String(),
		WithDynamoEndpoint("http://localhost:4569"),
	)

	if err != nil {
		t.Errorf("failed to create db instance")
		return
	}

	err = db.CreateTable()
	if err != nil {
		t.Errorf("failed to create test local table")
	}

	email := uuid.NewV4().String()

	err = db.Save(storage.Record{
		ID:           "1",
		Name:         "foo bar",
		Email:        email,
		MobileNumber: "1235",
	})
	if err != nil {
		t.Errorf("failed to save record, error: %+v", err)
		return
	}
}
