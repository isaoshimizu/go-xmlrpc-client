package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"log"
	"net/http"
	"os"
)

type MessageArgs struct {
	MessageBody string
}

type MessageReply struct {
	ResponseBody string
}

func XmlRpcCall(hostname string, port int, method string, args MessageArgs) (reply MessageReply, err error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	body := bytes.NewBuffer(buf)

	url := fmt.Sprintf("http://%s:%d/", hostname, port)
	resp, err := http.Post(url, "text/xml", body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	xml.DecodeClientResponse(resp.Body, &reply)
	return
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s -m [message] -h [host] -p [port]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	msg := flag.String("m", "", "message string test")
	hostname := flag.String("h", "localhost", "server address")
	port := flag.Int("p", 8000, "port")
	flag.Usage = usage
	flag.Parse()

	if len(*msg) == 0 {
		usage()
	}

	msgargs := MessageArgs{*msg}
	var reply MessageReply

	reply, err := XmlRpcCall(*hostname, *port, "MessageService.Send", msgargs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", reply.ResponseBody)
}
