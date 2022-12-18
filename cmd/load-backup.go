// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>

// */
package cmd

// import (
// 	"fmt"
// 	"os"
// 	"encoding/json"
// 	b64 "encoding/base64"

// 	"github.com/spf13/cobra"
// 	"github.com/lukehinds/graphparse/pkg/utils"
// 	// "github.com/sigstore/sigstore/pkg/cryptoutils"
// )

// type SignedSLSAProvenance struct {
// 	PayloadType string `json:"payloadType"`
// 	Payload     string `json:"payload"`
// 	Signatures  []struct {
// 		Keyid string `json:"keyid"`
// 		Sig   string `json:"sig"`
// 		Cert  string `json:"cert"`
// 	} `json:"signatures"`
// }
  

// // loadCmd represents the load command
// var loadCmd = &cobra.Command{
// 	Use:   "load",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fileflag, err := cmd.Flags().GetString("file")
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		prov_file, err := os.ReadFile(fileflag)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		var sigstore SignedSLSAProvenance

// 		err = json.Unmarshal(prov_file, &sigstore)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// //  decode sigstore.Payload from base64
// 		decoded_payload, err := b64.StdEncoding.DecodeString(sigstore.Payload)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		pubKey, err := utils.GetPubKeyFromCert(sigstore.Signatures[0].Cert)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// load signature and decode from base64
// 		decoded_sig, err := b64.StdEncoding.DecodeString(sigstore.Signatures[0].Sig)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// verify signature
// 		verified, err := utils.VerifySignature([]byte(decoded_payload), []byte(decoded_sig), pubKey)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		fmt.Println("Verified: ", verified)

// 	},
// }

// func init() {
// 	rootCmd.AddCommand(loadCmd)
// 	loadCmd.PersistentFlags().StringP("file", "f", "", "file to load")
// }
