package leap

import (
	"os"
	"path/filepath"
)

const (
	cfgFileName = ".leap"
	cfgFilePath = "./"
)

var (
	defaultLeapInfo   *LeapInfo
	defaultConfigPath = filepath.Join(GetHomeDir(), cfgFileName)
	DefaultCfgPath    = filepath.Join(GetHomeDir(), cfgFileName)
)

func openFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Execute the unix cmd echo $HOME and return it as a string.
func GetHomeDir() string {
	return os.Getenv("HOME")
}

func GetTempDir() string {
	return os.TempDir()
}
