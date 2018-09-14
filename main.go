package main

import "github.com/andrewarrow/arc/client"
import "fmt"

func main() {
	c := client.NewClient("")
	fmt.Println(c)
}
