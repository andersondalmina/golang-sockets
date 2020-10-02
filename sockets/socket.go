package sockets

import (
	"fmt"
	"os"
)

type SocketData struct {
	Action string
	Param  string
}

type SocketReply struct {
	Status int    `json:"status"`
	Data   []byte `json:"data"`
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
