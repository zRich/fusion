package did

import (
	_ "embed"
	"encoding/json"
	"testing"
)

var (
	//go:embed testdata/did1.json
	didJson string
)

func TestUnmarshalJSON(t *testing.T) {
	var doc Document
	// json.Unmarshal([]byte(did1Json), &doc)
	json.Unmarshal([]byte(didJson), &doc)
	// doc, err := Parse("did:a:123:456;service")
	// if err != nil {
	// 	t.Log("failed")
	// }
	t.Logf("did:%s:%s", doc.ID.Method, doc.ID.ID)
	t.Logf("%d vm", len(doc.VerificationMethod))
}
