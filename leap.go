package leap

import (
	"os"
)

const (
	cfgFileName = ".leap"
	cfgFilePath = "./"
)

var (
	defaultLeapConfig *LeapConfig
)

type Place struct {
	Directory string
	Alias     string
}

// Struct to hold a pointer to the config file.
type LeapConfig struct {
	configPath string
	Places     []Place
}

func NewLeapConfig(cfgFullPath string) *LeapConfig {
	lC := new(LeapConfig)

	lC.configPath = cfgFullPath

	// Try to get Stats on file. Err if nonexistent.
	_, err := os.Stat(cfgFullPath)

	// Ensure that the error type describes an existing file.
	if os.IsNotExist(err) == false {
		file, err := os.OpenFile(cfgFullPath, os.O_CREATE|os.O_RDWR, 0755)
		defer file.Close()

		// Panic if there was an error opening the file.
		if err != nil {
			panic(err.Error())
		}

		// Ensures error type describes a nonexistent file and creates it.
	} else if os.IsNotExist(err) {

		// Create a file with read write access.
		file, err := os.OpenFile(cfgFullPath, os.O_CREATE|os.O_RDWR, 0755)
		defer file.Close()

		// Panic if there was an error opening the file.
		if err != nil {
			panic(err.Error())
		}
	}

	// Try to get os.Stat info from the newly created file, store result.
	_, err = os.Stat(cfgFullPath)

	// Check to make sure that the newly opened file exists based on err.
	if err != nil && (os.IsNotExist(err) != false) {
		panic(err.Error())
	}

	return lC
}

func (lC *LeapConfig) readConfigFile() ([]Place, error) {
	return decodeJSON(lC.configPath)
}

/* Open the cfg file, add the entry. Might want to check if the dir is valid.
   Need to somehow have the path of the config file open too. This probably
   should not be public, or it should be a method. */
func (lC *LeapConfig) AddPlace(dir string, alias string) error {

	// Read the config file and update the object's Places field.
	placesFromCfg, err := lC.readConfigFile()
	if err != nil {
		return err
	}
	lC.Places = placesFromCfg

	newPlace := Place{dir, alias}

	// Add the new path to the Place slice.
	lC.Places = append(lC.Places, newPlace)

	// Write the updated Config file
	err = lC.writeToFile()
	if err != nil {
		return err
	}

	return nil
}

// Attempt to write the content stored in defaultLeapConfig to file.
func (lC *LeapConfig) writeToFile() error {
	// Truncate the file, write the new file content instead.
	err := os.Truncate(lC.configPath, 0)
	if err != nil {
		return err
	}

	err = encodeJSON(lC.configPath, lC.Places)
	if err != nil {
		return err
	}

	return nil
}
