package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetToken() interface{} {
	file, err := os.Open("token.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	val, _ := ioutil.ReadAll(file)
	var result map[string]interface{}
	json.Unmarshal(val, &result)
	return result["token"]
}
