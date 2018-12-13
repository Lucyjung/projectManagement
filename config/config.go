package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration Struct
type Configuration struct {
	DBUser   string `json:"DBUser"`
	DBPass   string `json:"DBPass"`
	DBServer string `json:"DBServer"`
}

// Read and parse the configuration file
func Read() Configuration {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(file)

	// we initialize our Users array
	var c Configuration

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &c)

	fmt.Println("Load Config Successfully")
	return c
}
