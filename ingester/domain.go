package ingester

// ReaderInterface represents the data reader interface
type ReaderInterface interface {
	Receive(handler Callback)
}

// Callback is representing the data callback handler function
type Callback func(Record)

// Record is representing a record int the database
type Record struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Email        string `json:"email,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
}
