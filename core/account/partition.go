package account

// Account Partition interface
type Partition interface {

	// GrantRead allows accountid read data from Partition
	GrantRead(partition []byte, accountID []byte, signature []byte) error
	RevokeRead(partition []byte, accountid []byte, signature []byte) error

	// GrantWrite allows accountID write data to Partition
	GrantWrite(partition []byte, accountID []byte, signature []byte) error
	RevokeWrite(partition []byte, accountID []byte, signature []byte) error

	// GrantUpdate grant accountID update existing data in Partition
	GrantUpdate(partition []byte, accountid []byte, signature []byte) error
	RevokeUpdate(partition []byte, accountid []byte, signature []byte) error
}

type PartitionReader interface {
	// Read reads key's value in Partition partiion
	Read(partition []byte, key []byte) ([]byte, error)
}

type PartitionWriter interface {
	// Write write value for key in Partition partiion
	Write(partition []byte, key []byte) error
}
