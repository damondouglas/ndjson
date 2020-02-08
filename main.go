package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var records []interface{}
	f, err := source()
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(f).Decode(&records)
	for _, k := range records {
		data, err := json.Marshal(k)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}


func source() (result io.Reader, err error) {
	args := os.Args[1:]
	if len(args) == 0 {
		result = os.Stdin
		return
	}
	return os.Open(args[0])
}