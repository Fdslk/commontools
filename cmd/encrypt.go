/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"crypto/cipher"

	"github.com/spf13/cobra"
)

var encrpty string

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "the string you want to encrypt",
	Long:  `the string you want to encrypt`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Encryption(encrpty); err != nil {
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&encrpty, "encrpty", "e", "", "the string you want to encrypt")
	encryptCmd.MarkFlagRequired("encrpty")
	rootCmd.AddCommand(encryptCmd)
}

func Encryption(plainText string) error {
	bytes := []byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	err := ioutil.WriteFile(fmt.Sprintf(encryptedFile), bytes, 0644)
	if err != nil {
		log.Fatalf("Writing encryption file: %s", err)
	} else {
		fmt.Printf("Message encrypted in file: %s\n\n", string(bytes))
	}
	return err
}
