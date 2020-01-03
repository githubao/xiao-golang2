// 使用原生的tcp网络连接
// author: baoqiang
// time: 2019/3/3 下午5:49
package gogo

import (
	"bufio"
	"log"
	"net"
	"github.com/pkg/errors"
	"sync"
	"io"
	"strings"
	"encoding/gob"
	"strconv"
	"flag"
)

// go run main.go
// go run main.go -connect localhost

type complexData struct {
	N int
	S string
	M map[string]int
	P []byte
	C *complexData
}

const (
	Port = ":61000"
)

func Open(addr string) (*bufio.ReadWriter, error) {
	log.Println("Dial: ", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "Dialing "+addr+" failed")
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

type HandlerFunc func(*bufio.ReadWriter)

type Endpoint struct {
	listener net.Listener
	handler  map[string]HandlerFunc
	m        sync.RWMutex
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		handler: map[string]HandlerFunc{},
	}
}

func (e *Endpoint) AddHandlerFunc(name string, f HandlerFunc) {
	e.m.Lock()
	e.handler[name] = f
	e.m.Unlock()
}

func (e *Endpoint) Listen() error {
	var err error
	e.listener, err = net.Listen("tcp", Port)
	if err != nil {
		return errors.Wrapf(err, "Unable to listen on port %s\n", Port)
	}
	for {
		log.Println("Accept a connection request.")
		conn, err := e.listener.Accept()
		if err != nil {
			log.Println("Failed accept a connection request: ", err)
			continue
		}
		log.Println("Handle incoming messages.")
		go e.handleMessages(conn)
	}
}
func (e *Endpoint) handleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()

	for {
		log.Print("Reveive command '")
		cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			log.Println("Reached EOF - close this connection.\n ---")
			return
		case err != nil:
			log.Println("\nError reading command. Got: '"+cmd+"'\n", err)
			return
		}

		cmd = strings.Trim(cmd, "\n ")
		log.Println(cmd + "'")

		e.m.Lock()
		handleCommand, ok := e.handler[cmd]
		e.m.Unlock()

		if !ok {
			log.Println("Command '" + cmd + "' is not registered.")
			return
		}

		handleCommand(rw)
	}

}

func handleStrings(rw *bufio.ReadWriter) {
	log.Print("Receive STRING message:")

	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
	}

	s = strings.Trim(s, "\n ")
	log.Println(s)

	_, err = rw.WriteString("Thank you.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}

	rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}

func HandleGob(rw *bufio.ReadWriter) {
	log.Print("Receive GOB data:")

	var data complexData

	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("Error decoding GOB data:", err)
		return
	}

	log.Printf("Outer complexData struct: \n%#v\n", data)
	log.Printf("Inner complexData struct: \n%#v\n", data.C)
}

func client(ip string) error {
	testStruct := complexData{
		N: 23,
		S: "string data",
		M: map[string]int{"one": 1, "two": 2, "three": 3},
		P: []byte("abc"),
		C: &complexData{
			N: 256,
			S: "Recursive structs? Piece of cake!",
			M: map[string]int{"01": 1, "10": 2, "11": 3},
		},
	}

	//open
	rw, err := Open(ip + Port)
	if err != nil {
		return errors.Wrap(err, "Client: Failed to open connection to "+ip+Port)
	}

	//send string
	log.Println("Send the string request.")
	n, err := rw.WriteString("STRING\n")
	if err != nil {
		return errors.Wrap(err, "Could not send the STRING request ("+strconv.Itoa(n)+" bytes written)")
	}

	// send data
	n, err = rw.WriteString("Additional data.\n")
	if err != nil {
		return errors.Wrap(err, "Could not send additional STRING data ("+strconv.Itoa(n)+" bytes written)")
	}
	log.Println("Flush the buffer.")

	// flush
	err = rw.Flush()
	if err != nil {
		return errors.Wrap(err, "Flush failed.")
	}

	// read reply
	log.Println("Read the reply.")
	response, err := rw.ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "Client: Failed to read the reply: '"+response+"'")
	}

	// print data
	log.Println("STRING request: got a response:", response)

	log.Println("Send a struct as GOB:")
	log.Printf("Outer complexData struct: \n%#v\n", testStruct)
	log.Printf("Inner complexData struct: \n%#v\n", testStruct.C)

	// 编码器
	enc := gob.NewEncoder(rw)
	n, err = rw.WriteString("GOB\n")
	if err != nil {
		return errors.Wrap(err, "Could not write GOB data ("+strconv.Itoa(n)+" bytes written)")
	}

	//编码
	err = enc.Encode(testStruct)
	if err != nil {
		return errors.Wrapf(err, "Encode failed for struct: %#v", testStruct)
	}

	// 发送数据
	err = rw.Flush()
	if err != nil {
		return errors.Wrap(err, "Flush failed.")
	}
	return nil
}

func server() error {
	endpoint := NewEndpoint()

	endpoint.AddHandlerFunc("STRING", handleStrings)
	endpoint.AddHandlerFunc("GOB", HandleGob)

	return endpoint.Listen()
}

func RunCS(){
	connect := flag.String("connect","","IP address of process to join. If empty, go into listen mode.")
	flag.Parse()

	if *connect != ""{
		err := client(*connect)
		if err != nil{
			log.Println("Error: ", errors.WithStack(err))
		}
		log.Println("client done.")
		return
	}

	err := server()
	if err != nil{
		log.Println("Error: ",errors.WithStack(err))
	}

	log.Println("server done.")

}

func init(){
	log.SetFlags(log.Lshortfile)
}