package vault

import (
	"context"

	"github.com/bytehubplus/fusion/did"
)

// vault is a sensetive data storage unit which can store one or more data entry in KV format
// did used as vault for being identification, authentication/authorization etc.
// each entry maybe encrypted
// a read reqeust contains:
// 1. vault ID,
// 2. entry id
// 3. encryption key, which can be used to encryp it before sending to requestor
type Vault interface {
	Get(key string) ([]byte, error)
	Put(Key string, value []byte) error
	Delete(key string) error
	VaultID() string
	Controllers() []string
	GetEntry(Id string) ([]byte, error)
	PutEntry(entry []byte) ([]byte, error)
}

type Depositer interface {
	DepositVerifiableCredential(ctx context.Context, vcJson []byte) (entryID string, err error)
}

type Withdrawer interface {
	WithdrawVerifiableCredential(ctx context.Context, vcID string) error
}

// / A vault Partition interface
type Partition interface {
	// CreatePartition creates a new Partition, return Partition ID, nil if seccuss otherwise return nil and an error
	// purpose the other part vault id or data factory id
	CreatePartition(purpose []byte, signature []byte) ([]byte, error)

	// CleanPartition cleans data for Partition
	CleanPartition(partition []byte, signature []byte) error

	// lockPartition locks the Partition, which causes only controllers can access the Partition, others cannot
	LockPartition(partition []byte, signature []byte)

	// GrantRead allows vaultid read data from Partition
	GrantRead(partition []byte, vaultID []byte, signature []byte) error
	RevokeRead(partition []byte, vaultid []byte, signature []byte) error

	// GrantWrite allows vaultID write data to Partition
	GrantWrite(partition []byte, vaultID []byte, signature []byte) error
	RevokeWrite(partition []byte, vaultID []byte, signature []byte) error

	// GrantUpdate grant vaultID update existing data in Partition
	GrantUpdate(partition []byte, vaultid []byte, signature []byte) error
	RevokeUpdate(partition []byte, vaultid []byte, signature []byte) error

	// Read reads key's value in Partition partiion
	Read(partition []byte, key []byte, signature []byte) ([]byte, error)

	// Write write value for key in Partition partiion
	Write(partition []byte, key []byte, value []byte, signature []byte) error

	// Update updates value for key in Partition partiion
	Update(partition []byte, key []byte, value []byte, signature []byte) error
}

type VaultProvider interface {
	//CreateVault creates a new vault, the creator's signature of vault ID and controllers are required.
	CreateVault(config Config) (Vault, error)
	//Open opens an existing vault from filesystem
	Open(id string) (Vault, error)

	//obsolete, do not use
	OpenWithDid(did did.DID) (Vault, error)
}

type Config struct {
	// RootFSPath is the directory where vault is stored
	RootFSPath string
	// DID        did.DID
	// DBConfig   *DBConfig
}

type DBConfig struct {
	DBPath string
}
