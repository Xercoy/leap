package leap

import (
	"errors"
	"fmt"
	"log"
	"os"
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

func (lI *LeapInfo) StrOfPlaces() string {
	var strCollection string
	var formatString = "\n%-10s   %-10s\n"
	var header = fmt.Sprintf(formatString, "Aliases", "Directories")

	// Display appropriate header if there's only one place.
	if len(lI.Places) == 1 {
		header = fmt.Sprintf(formatString, "Alias", "Directory")
	}

	// Append header
	strCollection += header

	// Iterate through the slice and append the formatted strings.
	for _, p := range lI.Places {
		strCollection += fmt.Sprintf(formatString, p.Alias, p.Directory)
	}

	return strCollection
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

func (lC *LeapInfo) aliasInUse(testAlias string) bool {
	for _, p := range lC.Places {
		if p.Alias == testAlias {
			return true
		}
	}

	return false
}

func (lI *LeapInfo) RemovePlace(alias string) error {
	var err error
	var rmIndex int = -1

	lI.Places, err = lI.readConfigFile()
	if err != nil {
		return err
	}

	for i, place := range lI.Places {
		if place.Alias == alias {
			rmIndex = i
		}
	}

	if rmIndex > -1 {
		lI.Places = append(lI.Places[:rmIndex], lI.Places[rmIndex+1:]...)
	} else {
		return errors.New("Removal Error: Place with given alias was not found.")
	}

	// Write the updated Config file
	err = lI.writeToConfig()
	if err != nil {
		return err
	}

	return nil
}

func (lI *LeapInfo) readConfig() error {
	// Read the config file and update the object's Places field.
	updatedPlaces, err := lI.readConfigFile()
	if err != nil {
		log.Printf("AddPlace: Error Reading Config file...")
		return err
	}

	lI.Places = updatedPlaces

	return nil
}

/* Open the cfg file, add the entry. Might want to check if the dir is valid.
   Need to somehow have the path of the config file open too. */
func (lI *LeapInfo) AddPlace(dir string, alias string) error {
	err := lI.readConfig()
	if err != nil {
		return err
	}

	if lI.aliasInUse(alias) {
		errorStr := fmt.Sprintf("Error: Alias (%s) is already in use. Unable to add Place.\n", alias)

		return errors.New(errorStr)
	}

	newPlace := Place{dir, alias}

	// Add the new path to the Place slice.
	lI.Places = append(lI.Places, newPlace)

	// Write the updated Config file
	err = lI.writeToConfig()
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

	resolvedDir, _ = lC.ResolveAlias(alias)

	return resolvedDir
}

// Return full directory
func (lC *LeapInfo) ResolveAlias(alias string) (string, error) {
	lC.readConfigFile()

	unresolvedErr := errors.New("Could not resolve alias.")

	for _, p := range lC.Places {
		if p.Alias == alias {
			return p.Directory, nil
		}
	}

	return ".", unresolvedErr
}
