package sockets

import (
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

const bufferSize = 1024 * 1024

// SocketClient is a struct to socket client
type SocketClient struct {
	conn  *net.TCPConn
	Reply SocketReply
}

// CreateSocketClient return a client connected with the socket
func CreateSocketClient(host string, port string) SocketClient {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	return SocketClient{
		conn: conn,
	}
}

// Send a message to socket
func (s *SocketClient) Send(action string, Param map[string]string) {
	b, err := msgpack.Marshal(&SocketData{
		Action: action,
		Param:  Param,
	})

	checkError(err)

	s.conn.Write(b)

	buf := make([]byte, bufferSize)
	_, err = s.conn.Read(buf)
	checkError(err)

	err = msgpack.Unmarshal(buf, &s.Reply)
	checkError(err)
}

// Close the connection with the socket
func (s *SocketClient) Close() {
	s.conn.Close()
}
