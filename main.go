package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	fmt.Println(url.PathEscape(os.Args[1]))
}
