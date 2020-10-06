package sockets

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/andersondalmina/golang-sockets/services"
	"github.com/vmihailenco/msgpack/v5"
)

const dataBufferSize = 1024 * 100

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

// Listen for messages on socket
func (s *SocketServer) Listen() {
	for {
		conn, err := s.listener.Accept()
		fmt.Printf("Client Connected: %s\n", conn.LocalAddr())

		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, dataBufferSize)
	_, err := conn.Read(buf)

	if err == io.EOF {
		return
	}

	checkError(err)

	var d SocketData
	err = msgpack.Unmarshal(buf, &d)
	checkError(err)

	fmt.Printf("Received : %+v\n", d)

	var r SocketReply
	handleSocketData(d, &r)

	b, err := msgpack.Marshal(r)
	checkError(err)

	_, err = conn.Write(b)
	checkError(err)
}

// Close the connection with the socket
func (s *SocketServer) Close() {
	s.listener.Close()
}

func handleSocketData(d SocketData, r *SocketReply) {
	switch d.Action {
	case "createBook":
		book, err := services.CreateBook(d.Param)

		if err == nil {
			json, err := json.Marshal(book)
			checkError(err)

			r.Data = json
		}

		r.Error = err

	case "updateBook":
		book, err := services.UpdateBook(d.Param)

		json, err := json.Marshal(book)
		checkError(err)

		r.Data = json
		r.Error = err

	case "searchBookByTitle":
		books := services.SearchBooksByTitle(d.Param["title"])

		json, err := json.Marshal(books)
		checkError(err)

		r.Data = json
		r.Error = err

	case "searchBookByAuthor":
		books := services.SearchBooksByAuthor(d.Param["author"])

		json, err := json.Marshal(books)
		checkError(err)

		r.Data = json
		r.Error = err

	case "searchBookByEdition":
		books := services.SearchBooksByEdition(d.Param["edition"])

		json, err := json.Marshal(books)
		checkError(err)

		r.Data = json
		r.Error = err

	case "searchBookByYear":
		books := services.SearchBooksByYear(d.Param["year"])

		json, err := json.Marshal(books)
		checkError(err)

		r.Data = json
		r.Error = err

	case "deleteBookByTitle":
		err := services.DeleteBooksByTitle(d.Param["title"])
		r.Error = err
	}
}
