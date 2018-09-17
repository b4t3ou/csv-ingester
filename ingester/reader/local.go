package reader

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/b4t3ou/csv-ingester/ingester"
)

// LocalCSVReader representing the local csv reader
type LocalCSVReader struct {
	reader *csv.Reader
}

// NewLocalCSVReader returning with a new local csv reader
func NewLocalCSVReader(filePath string) (*LocalCSVReader, error) {
	r := &LocalCSVReader{}

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	r.reader = csv.NewReader(bufio.NewReader(csvFile))

	return r, nil
}

// Receive reading the file and passing to the callback handler
func (r *LocalCSVReader) Receive(handler ingester.Callback) {
	for {
		line, error := r.reader.Read()
		if error == io.EOF {
			break
		}

		if error != nil {
			log.Fatal().Err(error).Msg("failed to read line")
		}

		if line[0] == "id" {
			continue
		}

		if len(line) < 4 {
			log.Error().Interface("line", line).Msg("a few fields are missing")
			continue
		}

		handler(ingester.Record{
			ID:           line[0],
			Name:         line[1],
			Email:        line[2],
			MobileNumber: line[3],
		})
	}
}
