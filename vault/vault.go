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

	SectionPrefix []byte = []byte("SEC")
)

// KvVault, Vault in KV database
type KvVault struct {
	db      *leveldb.DB
	metaDoc did.DID
	lock    sync.RWMutex
}

type vaultSection struct {
	readDoc   did.Document
	writeDoc  did.Document
	updateDoc did.Document
}

// CreateSecion creates a new section, return section ID, nil if seccuss otherwise return nil and an error
// purpose the other part vault id or data factory id
func (k *KvVault) CreateSecion(purpose []byte) ([]byte, error) {
	if k.SectionExist(purpose) {
		return nil, errors.New("section already exists")
	} else {
		section := fmt.Sprintf("%s%s", SectionPrefix, purpose)
		k.Put(section, []byte("SECTION"))
		return []byte(section), nil
	}
}

func (v *KvVault) SectionExist(section []byte) bool {
	exist, err := v.Get(string(fmt.Sprintf("%s%s", SectionPrefix, section)))
	if err != nil && exist != nil {
		return true
	}
	return false
}

// CleanSecion cleans data for section
func (k *KvVault) CleanSecion(section []byte) error {
	iter := k.db.NewIterator(util.BytesPrefix(section), nil)

	for iter.Next() {
		key := iter.Key()
		k.db.Delete(key, nil)
	}
	iter.Release()

	return nil
}

// locksection locks the section, which causes only controllers can access the section, others cannot
func (k *KvVault) LockSection(section []byte) {
	k.Put(fmt.Sprintf("%s%s", SectionPrefix, section), []byte("LOCK"))
}

// GrantRead allows vaultid read data from section
func (k *KvVault) GrantRead(section []byte, vaultID []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", SectionPrefix, section, vaultID), []byte("READ"))
}

func (k *KvVault) RevokeRead(section []byte, vaultid []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", SectionPrefix, section, vaultid), []byte("NOREAD"))
}

// GrantWrite allows vaultID write data to section
func (k *KvVault) GrantWrite(section []byte, vaultID []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", SectionPrefix, section, vaultID), []byte("WRITE"))
}

func (k *KvVault) RevokeWrite(section []byte, vaultID []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", SectionPrefix, section, vaultID), []byte("NOWRITE"))
}

// GrantUpdate grant vaultID update existing data in section
func (k *KvVault) GrantUpdate(section []byte, vaultid []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", SectionPrefix, section, vaultid), []byte("UPDATE"))
}

func (k *KvVault) RevokeUpdate(section []byte, vaultid []byte) error {
	return k.Put(fmt.Sprintf("%s%s%s", SectionPrefix, section, vaultid), []byte("NOUPDATE"))
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
