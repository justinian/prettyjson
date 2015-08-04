package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("Error reading stdin: " + err.Error())
	}
	if len(data) < 2 {
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "    "); err != nil {
		panic("Bad JSON input: " + err.Error())
	}
	out.Write([]byte("\n"))

	if _, err := os.Stdout.Write(out.Bytes()); err != nil {
		panic("Error printing result: " + err.Error())
	}
}
