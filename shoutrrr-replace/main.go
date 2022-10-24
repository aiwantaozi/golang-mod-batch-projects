package main

import (
	"fmt"
	"github.com/containrrr/shoutrrr"
)

func main() {
	url := "slack://token-a/token-b/token-c"
	err := shoutrrr.Send(url, "Hello world (or slack channel) !")
	if err != nil {
		fmt.Println(err)
	}
}
