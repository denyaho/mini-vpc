package main

import (
	"github.com/daiogawa/mini-vpc/netns-lab"
)

func main() {
	err := netnslab.createNamespace("example")
	if err != nil {
		panic(err)
	}
}