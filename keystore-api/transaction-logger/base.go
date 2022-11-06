package transaction_logger

// TransactionLogger is used to write transaction logs
// of mutating operations to a data source.
type TransactionLogger interface {
	Run()
	WritePut(key string, value string)
	WriteDelete(key string)
	ReadEvents() (<-chan Event, <-chan error)
	Err() <-chan error
}
