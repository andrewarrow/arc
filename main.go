package main

import "github.com/andrewarrow/arc/client"
import "fmt"

func main() {
	c := client.NewClient("127.0.0.1")
	fmt.Println(c.Get(""))
}
