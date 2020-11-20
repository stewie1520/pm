package helpers

import (
	"log"
	"time"

	"github.com/stewie1520/pm/data"
)

// RequireForMasterPassword return if user need to provide master password to continue their action
func RequireForMasterPassword() bool {
	lastLoginTime, err := data.GetLastLogin()
	if err != nil {
		log.Fatal(err)
	}

	if time.Now().Sub(*lastLoginTime) > 15*time.Minute {
		return false
	}

	return true
}
