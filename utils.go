package arraymethods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LoadJSONAsMap loads JSON files from a path give in argument
func LoadJSONAsMap(filePath string) map[string]interface{} {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Json file loading failed: ", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
