// golang的websocket
// author: baoqiang
// time: 2019/2/27 下午5:23
package lang

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
	"math/rand"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Println("can't receive")
		}

		log.Println("received back from client: " + reply)

		msg := "Received: " + reply
		log.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			log.Println("can't send")
			break
		}

		for i := 0; i < 3; i++ {
			ts := time.Now()

			if err = websocket.Message.Send(ws, ts.String()); err != nil {
				log.Println("can't send")
				break
			}

			time.Sleep(time.Duration(rand.Float32()) * time.Second)
		}

	}
}

func RunWs() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}

}
