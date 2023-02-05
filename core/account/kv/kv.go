package kv

import (
	"fmt"

	"github.com/bytehubplus/fusion/core/account"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
	ACCOUNT_VERSION = "1.0"
	ACCOUNT_TYPE    = "Personal"
)

type KVAccount struct {
	db        *leveldb.DB
	storeName string
	meta      *AccountMeta
}

type AccountMeta struct {
	version     string
	AccountType string
	pubKey      []byte
}

func (am *AccountMeta) Version() string {
	return ACCOUNT_VERSION
}

func (am *AccountMeta) Type() string {
	return ACCOUNT_TYPE
}

func (acc *KVAccount) Version() string {
	return acc.meta.Version()
}

func (acc *KVAccount) AccountType() string {
	return acc.meta.Type()
}

func (acc *KVAccount) PubKey() interface{} {
	panic("not implemented")
}

// CreatePartition creates a new partition, return partition ID, nil if seccuss otherwise return nil and an error
// partition the other part account id or data factory id
func (k *KVAccount) CreatePartition(partition []byte, signature []byte) (account.Partition, error) {
	panic("not implemented") // TODO: Implement
}

// DeletePartition deletes a partition
func (k *KVAccount) DeletePartition(partition []byte, signature []byte) error {
	panic("not implemented") // TODO: Implement
}

type KVAccountProvider struct {
}

// CreateAccount creates a new account, the creator's signature of account ID and controllers are required.
func (k *KVAccountProvider) CreateAccount(config account.AccountConfig) (account.Account, error) {

	db, err := leveldb.OpenFile(config.LevelDBConfig.DBPath, &opt.Options{ErrorIfExist: true})

	if err != nil {
		return nil, err
	}

	account := KVAccount{db: db, storeName: config.LevelDBConfig.DBPath}

	return &account, nil
}

// Open opens an existing account
func (k *KVAccountProvider) Open(id string) (account.Account, error) {
	panic("not implemented") // TODO: Implement
}

const (
	KEY_FORMAT = "%s%s%s" // partition block key
	PUB_KEY    = "PUBKEY" // personal account has only one public key
)

func (acc KVAccount) accountKey(prefix string, block []byte, key []byte) []byte {
	rawKey := fmt.Sprintf(KEY_FORMAT, prefix, block, key)
	return []byte(rawKey)
}

func (acc KVAccount) accountMetaKey(block, key []byte) []byte {
	return acc.accountKey("META", block, key)
}
