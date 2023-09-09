// Main package (plugin print)
package main

import (
	"encoding/base64"
	"os"
	"strings"

	slingshot "github.com/bots-garden/slingshot/go-pdk"
	"github.com/valyala/fastjson"
)

func replaceHandler(argHandler []byte) []byte {
	input := string(argHandler)
	//slingshot.Print("input: " + string(input))

	p := fastjson.Parser{}
	jsonValue, err := p.Parse(input)
	if err != nil {
		slingshot.Log("😡 " + err.Error())
		os.Exit(1)
	}

	find := jsonValue.GetStringBytes("find")
	replace := jsonValue.GetStringBytes("replace")
	text := jsonValue.GetStringBytes("text")

	if find == nil {
		slingshot.Log("😡 find is nil")
		os.Exit(1)
	}

	if replace == nil {
		slingshot.Log("😡 replace is nil")
		os.Exit(1)
	}

	if text == nil {
		slingshot.Log("😡 text is nil")
		os.Exit(1)
	}

	//slingshot.Print("📝 find: " + string(find))
	//slingshot.Print("📝 replace: " + string(replace))

	// Decoding
	var content string
	decodedStrAsByteSlice, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		content = string(text)
	} else {
		content = string(decodedStrAsByteSlice)
	}

	//slingshot.Print("📝 text: " + content)

	newContent := strings.ReplaceAll(content, string(find), string(replace))

	return []byte(newContent)
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(replaceHandler)
}

func main() {}
