package client

import "fmt"
import "net"
import "bufio"

type Client struct {
	ip      string
	conns   []net.Conn
	markers []bool
}

func NewClient(ip string, size int) *Client {
	c := Client{}
	c.ip = ip
	c.conns = make([]net.Conn, size)
	c.markers = make([]bool, size)
	for i, _ := range c.conns {
		c.conns[i] = connect(ip)
	}
	return &c
}

func connect(ip string) net.Conn {
	var connp net.Conn
	var err error
	connp, err = net.Dial("tcp", ip+":6379")

	if err != nil {
		fmt.Printf("Some error %v", err)
		return nil
	}
	return connp
}

func (c *Client) conn() net.Conn {
	for i, conn := range c.conns {
		if c.markers[i] == false {
			c.markers[i] = true
			return conn
		}
	}
	return nil
}

func (c *Client) Get(key string) string {
	conn := c.conn()
	write(conn, "*2\r\n$3\r\nGET\r\n$2\r\nhi\r\n")
	p := make([]byte, 1024)
	leni, _ := bufio.NewReader(conn).Read(p)
	if leni > 0 {
		payload := string(p[0:leni])
		fmt.Printf("%s\n", payload)
	}
	return ""
}

func write(conn net.Conn, s string) error {
	_, err := conn.Write([]byte(s))
	if err != nil {
		fmt.Printf("Some error %v\n", err)
	}
	return err
}
