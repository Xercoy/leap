package leap

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	cfgFileName = ".leap"
	cfgFilePath = "./"
)

// +build !testing
func init() {
	cfgFullPath, err := filepath.Abs(cfgFilePath + cfgFileName)
	if err != nil {
		panic(err.Error())
	}

	_, err = os.Stat(cfgFullPath)

	if os.IsNotExist(err) == false {
		fmt.Println("Config file already exists.")

		return
	}

	if os.IsNotExist(err) {
		fmt.Println("File doesn't exist. Creating...")

		file, err := os.OpenFile(cfgFullPath, os.O_CREATE, 0777)
		defer file.Close()
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("File successfully created.")
	}

	_, err = os.Stat(cfgFullPath)

	if err != nil && (os.IsNotExist(err) != false) {
		panic(err.Error())
	}
}
