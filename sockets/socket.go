package sockets

import (
	"fmt"
	"os"
)

type SocketData struct {
	Action string
	Param  map[string]string
}

type SocketReply struct {
	Status int    `json:"status"`
	Error  error  `json:"error"`
	Data   []byte `json:"data"`
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
