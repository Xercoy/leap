package leap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

/* Use the json pkg to write content to a file. It needs to know what it's
   writing to, as well as the content to write. Return an error if any. */
func encodeJSON(filePath string, content interface{}) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()
	if err != nil {
		return err
	}

	log.Printf("Encode JSON, Writing to file (%v) (content): %v", filePath, content)

	err = os.Truncate(filePath, 0)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(file)
	err = enc.Encode(content)
	if err != nil {
		return err
	}

	/*
		var e []Entry
		byteContent, err := json.Marshal(content)
		err = json.Unmarshal(byteContent, &e)
		if err != nil {
			return err
		}*/

	return nil
}

// MIGHT WANNA CHANGE FILE NAME TO FILE PATH, and maybe return a pointer to an interface
/* Use the json pkg to decode JSON content from file according to the Entry
   type, which is just a struct of two strings, the dir and its alias. Return
   the parsed result as an Entry slice. */
func decodeJSON(fileName string) ([]Entry, error) {
	var entries []Entry

	// A json.Decoder requires an io.Reader. Get a *os.File accordingly.
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	byteBuffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteBuffer, &entries)
	if err != nil {
		return nil, err
	}

	log.Printf("Decode, entries: %v", entries)

	return entries, nil
}
