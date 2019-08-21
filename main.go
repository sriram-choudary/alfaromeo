package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type logWriter struct{}

// func (logWriter) Write(bs []byte), (int, error){
// 	fmt.Println(string(bs))
// 	return len(bs), nil

// }

func main() {
	//lw := logWriter{}
	fmt.Println(os.Args)
	//os.Args[1]
	//take everything that the user provides
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	fmt.Println(string(b))
}
