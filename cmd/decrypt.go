/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var decrpty string

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "the string you want to decrypt",
	Long:  `the string you want to decrypt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(string(Decryption()))
	},
}

func init() {
	// decryptCmd.Flags().StringVarP(&decrpty, "decrpty", "d", "", "the string you want to decrypt")
	decryptCmd.MarkFlagRequired("decrpty")
	rootCmd.AddCommand(decryptCmd)
}

func Decryption() []byte {
	bytes, err := ioutil.ReadFile(fmt.Sprintf(encryptedFile))
	if err != nil {
		log.Fatalf("Reading encrypted file: %s", err)
	}
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return bytes
}
