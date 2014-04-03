package main

import (
	"bytes"
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"log"
	"net/http"
	"flag"
	"os"
)

type MessageArgs struct{
	MessageBody string
}

type MessageReply struct {
	ResponseBody string
}

func XmlRpcCall(method string, args MessageArgs) (reply MessageReply, err error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	body := bytes.NewBuffer(buf)

	resp, err := http.Post("http://localhost:8000/", "text/xml", body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	xml.DecodeClientResponse(resp.Body, &reply)
	return
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [message]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	msg := flag.String("m", "", "message string")
	flag.Usage = usage
	flag.Parse()

	if len(*msg) == 0 {
		usage()
	}

	msgargs := MessageArgs{*msg}
	var reply MessageReply

	reply, err := XmlRpcCall("MessageService.Send", msgargs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", reply.ResponseBody)
}
