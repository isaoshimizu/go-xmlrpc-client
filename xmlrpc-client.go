package main

import (
	"bytes"
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"log"
	"net/http"
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

func main() {
	args := MessageArgs{os.Args[1]}
	var reply MessageReply

	reply, err := XmlRpcCall("MessageService.Send", args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", reply.ResponseBody)
}
