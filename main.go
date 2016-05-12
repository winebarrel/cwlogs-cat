package main

import (
	"bufio"
	"cwlogs_cat"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

func scan() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := ioutil.ReadAll(reader)
	return string(input)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	params := cwlogs_cat.ParseFlag()
	message := scan()
	err := cwlogs_cat.Cat(params, message)

	if err != nil {
		log.Fatal(err)
	}
}
