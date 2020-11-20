package helpers

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// PromptForPassword Prompt an input for password
func PromptForPassword(message string) (string, error) {
	fmt.Printf("%v: ", message)

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}
