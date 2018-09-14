package client

import "fmt"
import "net"
import "bufio"

import "errors"

type Client struct {
	ip   string
	conn net.Conn
}

func NewClient(ip string) *Client {
	c := Client{}
	c.ip = ip
	c.conn = connect(ip)
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

func (c *Client) Get(key string) string {
	c.write("*2\r\n$3\r\nGET\r\n$2\r\nhi\r\n")
	p := make([]byte, 1024)
	leni, _ := bufio.NewReader(c.conn).Read(p)
	if leni > 0 {
		payload := string(p[0:leni])
		fmt.Printf("%s\n", payload)
	}
	return ""
}

func (c *Client) write(s string) error {
	if c.conn == nil {
		return errors.New("")
	}
	_, err := c.conn.Write([]byte(s))
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		c.conn.Close()
		c.conn = nil
	}
	return err
}
