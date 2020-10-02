package sockets

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/andersondalmina/golang-sockets/services"
	"github.com/vmihailenco/msgpack/v5"
)

const dataBufferSize = 1024 * 10

// SocketServer is a struct to socket server
type SocketServer struct {
	listener *net.TCPListener
}

// CreateSocketServer return a socket server
func CreateSocketServer(host string, port string) SocketServer {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	return SocketServer{
		listener: listener,
	}
}

// Listen for messages to socket
func (s *SocketServer) Listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}

		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, dataBufferSize)
	_, err := conn.Read(buf)
	checkError(err)

	var d SocketData
	err = msgpack.Unmarshal(buf, &d)
	checkError(err)

	fmt.Printf("Received : %+v", d)

	var r SocketReply
	handleSocketData(d, &r)

	b, err := msgpack.Marshal(r)

	checkError(err)
	conn.Write(b)
}

// Close the connection with the socket
func (s *SocketServer) Close() {
	s.listener.Close()
}

func handleSocketData(d SocketData, r *SocketReply) {
	switch d.Action {
	case "searchBookByTitle":
		books := services.SearchBooksByTitle(d.Param)

		json, err := json.Marshal(books)
		checkError(err)

		r.Status = 0
		r.Data = json

	case "deleteBookByTitle":
		err := services.DeleteBooksByTitle(d.Param)
		if err != nil {
			r.Status = 404
			return
		}

		r.Status = 0
	}
}
