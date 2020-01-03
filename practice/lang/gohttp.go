// http with middleware
// author: baoqiang
// time: 2018/12/21 下午5:18
package lang

import (
	"net/http"
	"context"
	"log"
	"strconv"
)

func RpcTimeoutMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "rpcTimeout", "5000")
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func handler(w http.ResponseWriter, req *http.Request) {
	value := req.Context().Value("rpcTimeout").(string)
	log.Printf("Got rpc time out: %v", value)

	// core logic

	w.Write([]byte("hello"))
}

func RunServerHttp() {
	http.Handle("/", RpcTimeoutMw(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}

type StrIntFunc func(str string) int64

func realFunc(str string) int64 {
	a, _ := strconv.ParseInt(str, 10, 64)
	return a
}

func RunFuncAdapter() {
	//var a = 5
	var b = realFunc
	var c = StrIntFunc(b)
	log.Printf("%d\n", c("321"))
}
