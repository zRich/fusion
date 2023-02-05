package vault

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/bytehubplus/fusion/did"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var (
	// Vault meta data prefix
	MetaPrefix []byte = []byte("VM")
	// Vault data entry prefix
	EntryPrefix []byte = []byte("VE")

	PartitionPrefix []byte = []byte("SEC")
)

// KvVault, Vault in KV database
type KvVault struct {
	db      *leveldb.DB
	metaDoc did.DID
	lock    sync.RWMutex
}

// Read reads key's value in Partition partiion
func (k *KvVault) Read(Partition []byte, key []byte, signature []byte) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// Write write value for key in Partition partiion
func (k *KvVault) Write(Partition []byte, key []byte, value []byte, signature []byte) error {
	panic("not implemented") // TODO: Implement
}

// Update updates value for key in Partition partiion
func (k *KvVault) Update(Partition []byte, key []byte, value []byte, signature []byte) error {
	panic("not implemented") // TODO: Implement
}

type vaultPartition struct {
	readDoc   did.Document
	writeDoc  did.Document
	updateDoc did.Document
}

// CreatePartition creates a new partition, return partition ID, nil if seccuss otherwise return nil and an error
// purpose the other part vault id or data factory id
func (k *KvVault) CreatePartition(purpose []byte) ([]byte, error) {
	if k.PartitionExist(purpose) {
		return nil, errors.New("partition already exists")
	} else {
		partition := fmt.Sprintf("%s%s", PartitionPrefix, purpose)
		k.Put(partition, []byte("PARTITION"))
		return []byte(partition), nil
	}
}

func (v *KvVault) PartitionExist(partition []byte) bool {
	exist, err := v.Get(string(fmt.Sprintf("%s%s", PartitionPrefix, partition)))
	if err != nil && exist != nil {
		return true
	}
	return false
}

// CleanPartition cleans data for partition
func (k *KvVault) CleanPartition(partition []byte) error {
	iter := k.db.NewIterator(util.BytesPrefix(partition), nil)

	for iter.Next() {
		key := iter.Key()
		k.db.Delete(key, nil)
	}
	iter.Release()

	return nil
}

// lockpartition locks the partition, which causes only controllers can access the partition, others cannot
func (k *KvVault) LockPartition(partition []byte) {
	k.Put(fmt.Sprintf("%s%s", PartitionPrefix, partition), []byte("LOCK"))
}

// GrantRead allows vaultid read data from partition
func (k *KvVault) GrantRead(partition []byte, vaultID []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", PartitionPrefix, partition, vaultID), []byte("READ"))
}

func (k *KvVault) RevokeRead(partition []byte, vaultid []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", PartitionPrefix, partition, vaultid), []byte("NOREAD"))
}

// GrantWrite allows vaultID write data to partition
func (k *KvVault) GrantWrite(partition []byte, vaultID []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", PartitionPrefix, partition, vaultID), []byte("WRITE"))
}

func (k *KvVault) RevokeWrite(partition []byte, vaultID []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", PartitionPrefix, partition, vaultID), []byte("NOWRITE"))
}

// GrantUpdate grant vaultID update existing data in partition
func (k *KvVault) GrantUpdate(partition []byte, vaultid []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", PartitionPrefix, partition, vaultid), []byte("UPDATE"))
}

func (k *KvVault) RevokeUpdate(partition []byte, vaultid []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", PartitionPrefix, partition, vaultid), []byte("NOUPDATE"))
}

// Controllers returns vault's controller DIDs
func (k *KvVault) Controllers() []string {
	var result []string
	rawData, err := k.Get("doc")
	if err != nil {
		return result
	}

	var doc did.Document
	json.Unmarshal(rawData, &doc)
	// doc, err := did.ParseDocument(rawData)
	for _, v := range doc.Controller {
		result = append(result, v.String())
	}
	return result
}

// PutEntry saves an entry data into vault, return entry's unique id if successful, otherwise return error
func (k *KvVault) PutEntry(entry []byte) ([]byte, error) {
	hash := sha256.Sum256(entry)
	key := fmt.Sprintf("%s%s", EntryPrefix, hash[:])
	err := k.Put(key, entry)
	if err != nil {
		return hash[:], nil
	}
	return nil, err
}

func (k *KvVault) GetEntry(Id string) ([]byte, error) {
	key := fmt.Sprintf("%s%s", EntryPrefix, Id)
	return k.Get(key)
}

func (k *KvVault) Get(key string) ([]byte, error) {
	return k.db.Get([]byte(key), nil)

}

func (k *KvVault) Put(Key string, value []byte) error {
	return k.db.Put([]byte(Key), value, nil)

}

func (k *KvVault) Delete(key string) error {
	return k.db.Delete([]byte(key), nil)
}

func (k *KvVault) VaultID() string {
	hash := sha256.Sum256([]byte(k.metaDoc.String()))
	return string(hash[:])
}

type Provider struct {
	RootFSPath string
	// Config     Config
}

func NewProvider(path string) *Provider {
	p := &Provider{RootFSPath: path}
	return p
}

func (p *Provider) Open(id string) (Vault, error) {
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/%s", p.RootFSPath, id), &opt.Options{ErrorIfMissing: true})
	if err != nil {
		return nil, err
	}

	vault := KvVault{db: db}
	return &vault, nil
}

func (p *Provider) OpenWithDid(did did.DID) (Vault, error) {
	return p.Open(did.String())
}

// CreateVault creates a new vault
// param
func (p *Provider) CreateVault(doc did.Document) (Vault, error) {

	//public key must not be empty
	if doc.IsAuthenticationEmpty() {
		return nil, errors.New("authentication cannot be empty")
	}

	//create but not open existing
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/%s", p.RootFSPath,
		p.createVaultID(did.DID(doc.ID))), &opt.Options{ErrorIfExist: true})
	if err != nil {
		return nil, err
	}

	vault := KvVault{db: db}
	didValue := doc.ID.String()
	vault.Put("did", []byte(didValue))
	raw, _ := doc.MarshalJSON()

	vault.Put("metaDoc", raw)

	return &vault, nil
}

func (p *Provider) createVaultID(did did.DID) string {
	hash := sha256.Sum256([]byte(did.String()))
	hexStr := hex.EncodeToString(hash[:])
	return hexStr
}
