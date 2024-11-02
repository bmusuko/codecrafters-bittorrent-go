package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jackpal/bencode-go"
	"os"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

type Detail struct {
	Length int64 `json:"length"`
}
type Info struct {
	Announce string `json:"announce"`
	Info     Detail `json:"info"`
}

// Ensures gofmt doesn't remove the "os" encoding/json import (feel free to remove this!)
var _ = json.Marshal

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345
func decodeBencode(bencodedString []byte) (interface{}, error) {
	buf := bytes.NewBuffer(bencodedString)
	return bencode.Decode(buf)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	command := os.Args[1]

	if command == "decode" {
		bencodedValue := os.Args[2]

		decoded, err := decodeBencode([]byte(bencodedValue))
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else if command == "info" {
		filename := os.Args[2]
		dat, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		decoded, err := decodeBencode(dat)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonOutput, _ := json.Marshal(decoded)
		info := new(Info)
		_ = json.Unmarshal(jsonOutput, info)
		fmt.Printf("Tracker URL: %s\n", info.Announce)
		fmt.Printf("Length: %d\n", info.Info.Length)
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
