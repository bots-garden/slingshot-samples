// Main package (plugin print)
package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func helloHandler(argHandler []byte) []byte {
	input := string(argHandler)	
	return []byte(" " + string(input) + " three ✅ 🎉")
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(helloHandler)
}

func main() {}
