// http router的一个小的demo
// author: baoqiang
// time: 2019/3/1 下午6:33
package impl

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

// http://127.0.0.1:8234/github/githubao
func RunRServer() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/github/*proxypath", proxy)
	router.GET("/github/*proxypath", proxy)
	log.Fatal(http.ListenAndServe(":8234", router))
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("hello"))
}

func proxy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	remote := "https://github.com"
	u := remote + ps.ByName("proxyPath")

	// 构建新的请求
	req, err := http.NewRequest(r.Method, u, r.Body)
	req.Header = r.Header
	client := http.DefaultClient

	resp, err := client.Do(req)

	// 出错
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err)
		return
	}

	// 读取数据
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err)
		return
	}

	//add header
	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	// set cookie
	for _, value := range resp.Request.Cookies() {
		w.Header().Add(value.Name, value.Value)
	}

	// 回写
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyRes)
}
