package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
)

func main() {
	fmt.Println("this is a  thing!")

	client := twitch.NewClient("sneakyfoxtrot", "oauth:yove6r5nqa74duvxqvf9za2qf220sb")
}
