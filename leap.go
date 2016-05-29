package leap

import (
	"log"
	"os"
	//	"os/exec"
)

/* A place is a pair of strings: the destination (Directory), and its alias which leap will resolve the destination from. */
type Place struct {
	Directory string
	Alias     string
}

// Struct to hold a pointer to the config file.
type LeapInfo struct {
	configPath string
	Places     []Place
}

func NewLeapInfo(cfgPath string) *LeapInfo {
	lC := new(LeapInfo)

	lC.configPath = cfgPath

	// Try to get Stats on file. Err if nonexistent.
	_, err := os.Stat(cfgPath)

	// Ensure that the error type describes an existing file.
	if os.IsNotExist(err) == false {
		file, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_RDWR, 0755)
		defer file.Close()

		// Panic if there was an error opening the file.
		if err != nil {
			panic(err.Error())
		}

		// Ensures error type describes a nonexistent file and creates it.
	} else if os.IsNotExist(err) {

		// Create a file with read write access.
		file, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_RDWR, 0755)
		defer file.Close()

		// Panic if there was an error opening the file.
		if err != nil {
			panic(err.Error())
		}
	}

	// Try to get os.Stat info from the newly created file, store result.
	_, err = os.Stat(cfgPath)

	// Check to make sure that the newly opened file exists based on err.
	if err != nil && (os.IsNotExist(err) != false) {
		panic(err.Error())
	}

	lC.Places, err = lC.readConfigFile()
	if err != nil {
		panic(err.Error())
	}

	return lC
}

/* Open the cfg file, add the entry. Might want to check if the dir is valid.
   Need to somehow have the path of the config file open too. */
func (lC *LeapInfo) AddPlace(dir string, alias string) error {

	// Read the config file and update the object's Places field.
	placesFromCfg, err := lC.readConfigFile()
	if err != nil {
		log.Printf("AddPlace: Error Reading Config file...")
		return err
	}

	lC.Places = placesFromCfg

	newPlace := Place{dir, alias}

	// Add the new path to the Place slice.
	lC.Places = append(lC.Places, newPlace)

	// Write the updated Config file
	err = lC.writeToConfig()
	if err != nil {
		return err
	}

	return nil
}

// Attempt to write the content stored in defaultLeapInfo to file.
func (lC *LeapInfo) writeToConfig() error {
	return encodeJSON(lC.configPath, lC.Places)
}

func (lC *LeapInfo) readConfigFile() ([]Place, error) {
	return decodeJSON(lC.configPath)
}

/* In our case, by leaping, we simply return to stdout the destination that the givenalias resolves to. */
func (lC *LeapInfo) Leap(alias string) string {
	// Default value is a dot, so cd . won't do anything.
	resolvedDir := "."

	resolvedDir = lC.resolveAlias(alias)

	return resolvedDir
}

// Return full directory
func (lC *LeapInfo) resolveAlias(alias string) string {
	lC.readConfigFile()

	for _, p := range lC.Places {
		if p.Alias == alias {
			return p.Directory
		}
	}

	return "."
}
