package leap

import (
	"log"
	"os"
	"path/filepath"
)

const (
	cfgFileName = ".leap"
	cfgFilePath = "./"
)

// +build !testing
func init() {
	cfgFullPath, err := filepath.Abs(filepath.Join(GetHomeDir(), cfgFileName))
	if err != nil {
		panic(err.Error())
	}

	_, err = os.Stat(cfgFullPath)

	if os.IsNotExist(err) == false {
		log.Println("Config file already exists.")

		return
	}

	if os.IsNotExist(err) {
		log.Println("File doesn't exist. Creating...")

		file, err := os.OpenFile(cfgFullPath, os.O_CREATE|os.O_RDWR, 0755)
		defer file.Close()
		if err != nil {
			panic(err.Error())
		}

		log.Println("File successfully created.")
	}

	_, err = os.Stat(cfgFullPath)

	if err != nil && (os.IsNotExist(err) != false) {
		panic(err.Error())
	}
}

// Execute the unix cmd echo $HOME and return it as a string.
func GetHomeDir() string {
	return os.Getenv("HOME")
}
