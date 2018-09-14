package client

//import "fmt"

type Client struct {
	ip string
}

func NewClient(ip string) *Client {
	c := Client{}
	c.ip = ip
	return &c
}

func (c *Client) Get(key string) string {
	return ""
}
