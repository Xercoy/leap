package leap

import (
	"os"
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
