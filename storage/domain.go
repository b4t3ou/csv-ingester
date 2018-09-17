package storage

// DBInterface representing a database interface
type DBInterface interface {
	Save(item Record) error
}

// Record is representing a record int the database
type Record struct {
	ID           string `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
	// TODO add validator
}
