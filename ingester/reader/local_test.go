package reader

import (
	"testing"

	"github.com/b4t3ou/csv-ingester/ingester"
)

type handlerMock struct {
	data []ingester.Record
}

func (hm *handlerMock) Read(record ingester.Record) {
	hm.data = append(hm.data, record)
}

func TestLocalCSVReader_Receive(t *testing.T) {
	reader, err := NewLocalCSVReader("test-data.csv")
	if err != nil {
		t.Errorf("failed to load file, error: %+v", err)
	}

	handler := &handlerMock{
		data: []ingester.Record{},
	}

	reader.Receive(handler.Read)

	if len(handler.data) != 4 {
		t.Errorf("failed to get the righ amount of data, expected 4, got %d", len(handler.data))
	}
}
