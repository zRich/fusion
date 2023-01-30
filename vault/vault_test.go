package vault

import (
	"crypto/sha256"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/bytehubplus/fusion/did"
)

var (
	//go:embed testdata/valid_doc.jsonld
	validDoc string
	//go:embed testdata/did1.json
	did1Json string
)

func TestCreateHash(t *testing.T) {
	hash := sha256.Sum256([]byte("did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff"))
	t.Logf("hash: %s", hash[:])
}

func TestCreateVault(t *testing.T) {
	p := NewProvider("./data/")
	// doc, _ := did.ParseDocument([]byte(validDoc))
	var doc did.Document
	json.Unmarshal([]byte(did1Json), &doc)
	// json.Unmarshal([]byte(validDoc), &doc)
	v, _ := p.CreateVault(doc)
	didValue, _ := v.Get("did")
	t.Logf("did value: %s", didValue)
}

func TestUnmashalJSON(t *testing.T) {
	var doc did.Document
	json.Unmarshal([]byte(did1Json), &doc)
	t.Logf("did value: %s", doc.ID.ID)
}
