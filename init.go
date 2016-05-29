package leap

import (
	"log"
	"path/filepath"
)

func createDefaultLeapConfig() {
	log.Print("Pkg init running...")

	if defaultLeapConfig == nil {
		defaultPath, err := filepath.Abs(filepath.Join(GetHomeDir(),
			cfgFileName))

		if err != nil {
			panic(err.Error())
		}

		defaultLeapConfig = NewLeapConfig(defaultPath)
	}
}
