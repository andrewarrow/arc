package client

import "fmt"
import "time"
import "sync"
import "net"
import "bufio"

type Client struct {
	ip      string
	conns   []net.Conn
	markers []bool
	mu      sync.Mutex
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

func (c *Client) findConn() (net.Conn, int) {
	for {
		conn, i := c.conn()
		if i > -1 {
			return conn, i
		}
		time.Sleep(time.Millisecond * 20)
	}
}

func (c *Client) conn() (net.Conn, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, conn := range c.conns {
		if c.markers[i] == false {
			c.markers[i] = true
			return conn, i
		}
	}
	return nil, -1
}

func (c *Client) release(i int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.markers[i] = false
}

func (c *Client) Get(key string) string {
	conn, i := c.findConn()
	defer c.release(i)
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
