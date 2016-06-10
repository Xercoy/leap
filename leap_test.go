package leap

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

var (
	testFilePath = filepath.Join(GetTempDir(), "./sampleConfig.json")
)

func init() {
	//	testPath := filepath.Join(GetTempDir(), cfgFileName)
	defaultLeapInfo = NewLeapInfo(testFilePath)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestAddPlace(t *testing.T) {
	testDir := "foo"
	testAlias := "bar"

	err := defaultLeapInfo.AddPlace(testDir, testAlias)
	if err != nil {
		t.Error(err.Error())
	}

	resolvedVal, _ := defaultLeapInfo.ResolveAlias(testAlias)
	if testDir != resolvedVal {
		t.Error("Newly added Place was not found in config.")
	}
}

func TestGetTmpDir(t *testing.T) {
	tmpDir := GetTempDir()
	if tmpDir == "" {
		t.Errorf("Error retrieving tmp dir.")
	}
}

/* Run linux cmd to get home directory. Can't really go wrong, so fail only if
   there's a legitimate error that arises. */
func TestGetHomeDir(t *testing.T) {
	homeDir := GetHomeDir()
	if homeDir == "" {
		t.Errorf("Error retrieving envar value via os.Getenv().")
	}
}

func TestPlaceRemoval(t *testing.T) {
	testDir := "fizz"
	testAlias := "buzz"

	err := defaultLeapInfo.AddPlace(testDir, testAlias)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = defaultLeapInfo.ResolveAlias(testAlias)
	if err != nil {
		t.Error(errors.New("Newly added Place was not found in config."))
	}

	err = defaultLeapInfo.RemovePlace(testAlias)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = defaultLeapInfo.ResolveAlias(testAlias)
	if err == nil {
		t.Error(errors.New("Removed place still exists in config."))
	}
}
