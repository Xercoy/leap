package leap

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	testFilePath = filepath.Join(GetTempDir(), "./sampleConfig.json")
)

func init() {
	//	testPath := filepath.Join(GetTempDir(), cfgFileName)
	defaultLeapConfig = NewLeapConfig(testFilePath)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestAddPlace(t *testing.T) {
	err := defaultLeapConfig.AddPlace("~/Base/workspace", "base")
	if err != nil {
		t.Error(err.Error())
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
