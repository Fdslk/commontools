/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbersChars = "0123456789"
	specialChars = "!@#$%^&*()-_=+{}[]|;:,.<>/?"
)

var length int32
var containDigit bool
var containUpperCase bool
var containLowerCase bool
var containSpecialCase bool

// pwdgeneratorCmd represents the pwdgenerator command
var pwdgeneratorCmd = &cobra.Command{
	Use:   "pwdgenerator",
	Short: "By this Command you can generate a new password with different requirement",
	Long:  `By this Command you can generate a new password with different requirement`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := generatePassword()
		for {
			if validatePassword(pwd) && err == nil {
				fmt.Println(pwd)
				break
			}
			fmt.Println(pwd + strconv.FormatBool(validatePassword(pwd)))
			pwd, err = generatePassword()
		}
	},
}

func init() {
	pwdgeneratorCmd.Flags().Int32VarP(&length, "length", "l", 0, "The length of the password")
	pwdgeneratorCmd.Flags().BoolVarP(&containDigit, "containDigit", "d", false, "The password should contain a digit")
	pwdgeneratorCmd.Flags().BoolVarP(&containUpperCase, "containUpperCase", "u", false, "The password should contain a upper case letter")
	pwdgeneratorCmd.Flags().BoolVarP(&containLowerCase, "containLowerCase", "o", false, "The password should contain a lower case letter")
	pwdgeneratorCmd.Flags().BoolVarP(&containSpecialCase, "containSpecialCase", "s", false, "The password should contain a special letter")
	pwdgeneratorCmd.MarkFlagRequired("length")
	rootCmd.AddCommand(pwdgeneratorCmd)
}

func generatePassword() (string, error) {
	allChars := lowerChars + upperChars + numbersChars + specialChars

	// Generate a random password string with the specified length
	password := ""
	for i := 0; i < int(length); i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		password += string(allChars[index.Int64()])
	}
	return password, nil
}

func validatePassword(password string) bool {
	containsLowercase := false
	containsUppercase := false
	containsDigit := false
	containsSpecial := false

	for _, c := range password {
		if unicode.IsDigit(c) || !containDigit {
			containsDigit = true
		}
		if unicode.IsLower(c) || !containLowerCase {
			containsLowercase = true
		}
		if unicode.IsUpper(c) || !containUpperCase {
			containsUppercase = true
		}
		if strings.Contains(specialChars, string(c)) || !containSpecialCase {
			containsSpecial = true
		}
	}

	return containsLowercase && containsUppercase && containsDigit && containsSpecial
}
