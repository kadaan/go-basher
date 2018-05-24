package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dustin/go-jsonpointer"
	"github.com/studioetrange/go-basher"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func jsonPointer(args []string) {
	if len(args) == 0 {
		os.Exit(3)
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	var o map[string]interface{}
	assert(json.Unmarshal(bytes, &o))
	println(jsonpointer.Get(o, args[0]).(string))
}

func reverse(args []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	runes := []rune(strings.Trim(string(bytes), "\n"))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	println(string(runes))
}

func main() {

	basher.Application(
		map[string]func([]string){
			"reverse":      reverse,
			"jsonPointer":	jsonPointer,
		}, []string{
			"bash/example.bash",
		},
		"main"
		Asset,
		nil,
		true,
	)

}
