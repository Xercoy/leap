package leap

import (
	"testing"
)

/* Run linux cmd to get home directory. Can't really go wrong, so fail only if
   there's a legitimate error that arises. */
func TestGetHomeDir(t *testing.T) {
	homeDir := GetHomeDir()
	if homeDir == "" {
		t.Errorf("Error retrieving envar value via os.Getenv().")
	}
}
