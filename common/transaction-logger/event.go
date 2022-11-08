package transaction_logger

// Event represents a transaction log entry
type Event struct {
	sequence  uint64    // unique id
	EventType EventType // action taken
	Key       string    // key affected by action
	Value     string    // value of a PUT operation
}

type EventType uint64

const (
	EVENT_TYPE_PUT EventType = iota + 1
	EVENT_TYPE_DELETE
)
