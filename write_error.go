package errjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Errors Error

func CreateFolder() {
	// create folder
	err := os.Mkdir("log", 0700)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func CreateFile(name string) *os.File {
	path := fmt.Sprintf("%s/%s", "log/", name)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return file
}

func WriteIntoJson(fileName string, data []byte) {
	file := CreateFile(fileName)
	file.Write(data)
	file.Sync()
	file.Close()
}

func WriteError(fileName string, errMessage string) {
	// exist error
	var current []Error

	// check is folder exist
	_, err := os.Stat("log")
	if os.IsNotExist(err) == true {
		CreateFolder()
	}

	// check file is exist
	path := fmt.Sprintf("%s/%s", "log/", fileName)
	_, err = os.Stat(path)

	if os.IsNotExist(err) == true {
		// file not exist
		current = append(current, Error{Time: time.Now(), Error: errMessage})
		json, _ := json.Marshal(current)
		WriteIntoJson(fileName, json)

	} else {
		// file exist
		reader, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		json.Unmarshal([]byte(reader), &current)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		current = append(current, Error{Time: time.Now(), Error: errMessage})
		json, err := json.MarshalIndent(current, "", " ")
		WriteIntoJson(fileName, json)
	}
}
