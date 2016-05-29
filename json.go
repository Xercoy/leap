package leap

import (
	"encoding/json"
	"io/ioutil"
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

	err = os.Truncate(filePath, 0)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(file)
	err = enc.Encode(content)
	if err != nil {
		return err
	}

	return nil
}

func decodeJSON(fileName string) ([]Place, error) {
	var entries []Place

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

	// Return early on empty files.
	if len(byteBuffer) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(byteBuffer, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
