package account

type BlockReader interface {
	// Read reads #number block from partition
	Read(partition, number []byte) ([]byte, error)
}

type BlockWriter interface {
	// Write writes a new block in partition, returns block number
	Write(partition, value []byte) ([]byte, error)
}

type BlockDeleter interface {
	// Delete deletes #number block from partition
	Delete(partition, number []byte) error
}
