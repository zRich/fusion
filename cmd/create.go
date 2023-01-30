/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/ed25519"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/bytehubplus/fusion/did"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		pubKey, _ := cmd.Flags().GetString("public")
		didFile, _ := cmd.Flags().GetString("output")
		CreateDIDDoc(pubKey, didFile)
	},
}

func init() {
	didCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	var pubKey string
	var didFile string
	createCmd.Flags().StringVarP(&pubKey, "public", "p", "public.pem", "public key in JsonWebKey2020 format.")
	createCmd.MarkFlagRequired("public")
	createCmd.Flags().StringVarP(&didFile, "output", "o", "", "DID document output file.")
	createCmd.MarkFlagRequired("output")
}

func CreateDIDDoc(pubKey string, didFile string) error {
	raw, _ := os.ReadFile(pubKey)
	block, rest := pem.Decode(raw)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal(rest)
	}

	pub := ed25519.PublicKey(block.Bytes)

	didID, _ := did.ParseDID("did:example:123")
	doc := &did.Document{
		Context: []did.URI{did.DIDContextV1URI()},
		ID:      *didID,
	}

	keyID, _ := did.ParseDIDURL("did:example:123#key-1")
	vm, _ := did.NewVerificationMethod(*keyID, did.JsonWebKey2020, did.DID{}, pub)
	doc.AddAssertionMethod(vm)
	didJson, _ := json.MarshalIndent(doc, "", "  ")

	err := os.WriteFile(didFile, didJson, 0644)
	return err
}
