package database

import (
	"github.com/b4t3ou/csv-ingester/storage"
	"github.com/rs/zerolog/log"
)

// Memory representing the memory type database
type Memory struct {
	data map[string]storage.Record
}

// NewMemory returning with a new memory type database
func NewMemory() *Memory {
	return &Memory{
		data: map[string]storage.Record{},
	}
}

// Save is saving an item into the database
func (m *Memory) Save(item storage.Record) error {
	m.data[item.Email] = item

	log.Info().Interface("record", item).Msg("record has been saved")

	return nil
}
