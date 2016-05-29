package leap

import (
	"log"
	"path/filepath"
)

func createDefaultLeapInfo() {
	log.Print("Pkg init running...")

	if defaultLeapInfo == nil {
		defaultPath, err := filepath.Abs(filepath.Join(GetHomeDir(),
			cfgFileName))

		if err != nil {
			panic(err.Error())
		}

		defaultLeapInfo = NewLeapInfo(defaultPath)
	}
}
