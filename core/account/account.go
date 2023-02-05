package account

const (
	TOPIC_LEN = 128 // partition length
)

// account is a sensetive data storage unit which can store one or more data entry in KV format
// did used as account for being identification, authentication/authorization etc.
// each entry maybe encrypted
// a read reqeust contains:
// 1. account ID,
// 2. entry id
// 3. encryption key, which can be used to encryp it before sending to requestor
type Account interface {
	// CreatePartition creates a new partition, return partition ID, nil if seccuss otherwise return nil and an error
	// partition the other part account id or data factory id
	CreatePartition(partition []byte, signature []byte) (Partition, error)

	// DeletePartition deletes a partition
	DeletePartition(partition []byte, signature []byte) error
}

type AccountProvider interface {
	//CreateAccount creates a new account, the creator's signature of account ID and controllers are required.
	CreateAccount(config AccountConfig) (Account, error)
	//Open opens an existing account
	Open(id string) (Account, error)
}

type AccountConfig struct {
	// LevelDBConfig is the directory where account data is stored
	LevelDBConfig *LevelDBConfig
	PubKey        []byte
}

type LevelDBConfig struct {
	DBPath string
}

type AccountMeta interface {
	Version() string
	AccountType() string
	PubKey() interface{}
}
