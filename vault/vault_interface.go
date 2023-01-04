package vault

import "github.com/bytehubplus/fusion/did"

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

type VaultSection interface {
	//CreateSecion creates a new section, return section ID, nil if seccuss otherwise return nil and an error
	// purpose the other part vault id or data factory id
	CreateSecion(purpose []byte) ([]byte, error)

	//CleanSecion cleans data for section
	CleanSecion(section []byte) error

	//locksection locks the section, which causes only controllers can access the section, others cannot
	LockSection(section []byte)
	//GrantRead allows vaultid read data from section
	GrantRead(section []byte, vaultID []byte) error
	RevokeRead(section []byte, vaultid []byte) error

	//GrantWrite allows vaultID write data to section
	GrantWrite(section []byte, vaultID []byte) error
	RevokeWrite(section []byte, vaultID []byte) error

	//GrantUpdate grant vaultID update existing data in section
	GrantUpdate(section []byte, vaultid []byte) error
	RevokeUpdate(section []byte, vaultid []byte) error
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
