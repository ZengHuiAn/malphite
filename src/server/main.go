package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data,err := ioutil.ReadFile("conf/server.json.json")
	fmt.Println(string(data),err)
}