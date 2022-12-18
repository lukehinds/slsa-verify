/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "crypto/sha256"
	"fmt"
	"os"
	// "encoding/json"
	"encoding/base64"

	"github.com/spf13/cobra"
	cjson "github.com/docker/go/canonical/json"
	"github.com/lukehinds/graphparse/pkg/utils"
	dsselib "github.com/secure-systems-lab/go-securesystemslib/dsse"

)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the file flag from viper
		fileflag, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Read the file into a byte array
		slsa_statement, err := os.ReadFile(fileflag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Unmarshal the byte array into a Envelope struct
		prov := &utils.Envelope{}
		if err := cjson.Unmarshal(slsa_statement, prov); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// // Get the cert from the envelope
		cert, err := utils.GetCert(slsa_statement)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// // Get the public key from the cert
		pubKey, err := utils.GetPubKeyFromCert(string(cert))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// base64 decode the signature
		sig, err := base64.StdEncoding.DecodeString(prov.Signatures[0].Sig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		body, err := base64.StdEncoding.DecodeString(prov.Payload)
		if err != nil {
			fmt.Println(err)
		}
		// Generate PAE(payloadtype, serialized body)
		paeEnc := dsselib.PAE(prov.PayloadType, body)

		// // Verify the signature
		result, err := utils.VerifySignature(pubKey, []byte(paeEnc), sig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if result {
			fmt.Println("Signature is valid")
		} else {
			fmt.Println("Signature is invalid")
		}
	},

}

func init() {
	rootCmd.AddCommand(loadCmd)
	loadCmd.PersistentFlags().StringP("file", "f", "", "file to load")
}
