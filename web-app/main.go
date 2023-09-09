// Main package of the web app
package main

import (
	_ "embed"
	"encoding/base64"
	"strings"

	slingshot "github.com/bots-garden/slingshot/go-pdk"
	"github.com/valyala/fastjson"
)

var (
	//go:embed resources/dist/index.html
	html []byte
	parser = fastjson.Parser{}
)

func helloHandler(httpRequestData []byte) []byte {

	
	message := "Hello From Slingshot"
	htmlPage := strings.Replace(
		string(html), 
      	"{{message}}", 
      	message, 
     	 -1,
	)
	
	// avoid special characters in jsonString
	html := base64.StdEncoding.EncodeToString([]byte(htmlPage))
	
	
	code := "200"

	headers := []string{
		`"Content-Type": "text/html; charset=utf-8"`, //Set the response format
		`"X-Slingshot-version": "0.0.3"`,
	}

	headersStr := strings.Join(headers, ",")

	response := `{"headers":{` + headersStr + `}, "textBody": "` + html + `", "statusCode": ` + code + `}`

	return []byte(response)

}

//export callHandler
func callHandler() {
	slingshot.Print("👋 callHandler function")
	slingshot.ExecHandler(helloHandler)
}

func main() {}

