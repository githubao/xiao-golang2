// all about http
// author: baoqiang
// time: 2019-08-26 17:41
package network

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var s = `
http协议是无状态的，面向连接的，可靠的协议

req: GET https://www.google.com HTTP/1.0
resp: HTTP/1.0 200 OK

常见的状态码：
"200" ; OK
"201" ; Created
"202" ; Accepted
"204" ; No Content
"301" ; Moved permanently
"302" ; Moved temporarily
"304" ; Not modified
"400" ; Bad request
"401" ; Unauthorised
"403" ; Forbidden
"404" ; Not found
"500" ; Internal server error
"501" ; Not implemented
"502" ; Bad gateway
"503" | Service unavailable
`

func RunHttp() {
	proxy := "http://127.0.0.1:8899"
	target := "http://www.google.com"
	auth := "user:mypassword"

	proxyUrl, err := url.Parse(proxy)
	checkError(err)

	targetUrl, err := url.Parse(target)
	checkError(err)

	// transport
	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}

	// request
	request, err := http.NewRequest("GET", targetUrl.String(), nil)

	//auth
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Proxy-Authorization", basic)

	// log it
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	// send request
	resp, err := client.Do(request)
	checkError(err)

	// read resp
	if resp.StatusCode != 200 {
		fmt.Printf("err: %v\n", resp.Status)
		return
	}

	var buf []byte
	reader := resp.Body
	defer reader.Close()

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			break
		}
		fmt.Printf("got: %v\n", string(buf[0:n]))
	}

}
