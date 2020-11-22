package helpers

import (
	"time"

	"github.com/stewie1520/pm/data"
)

// RequireForMasterPassword return if user need to provide master password to continue their action
func RequireForMasterPassword() bool {
	lastLoginTime := data.GetLastActivityTime()
	return time.Now().Sub(*lastLoginTime) > 15 * time.Minute
}
