package main

import (
	"fmt"

	"./proxy"
)

func main() {
	fmt.Println("Starting Software...")
	/* What should happen is , 2 rndmly ports in the range from 1200 - 22000 should be picked, with port != port2, those will be printed out to the console.
	   We enable forwarding on the listening Port, connect to the other rndmly choosen Port, and pipe our connection through there, this makes it harder to track.*/
	proxy.EnableForwarding_tcp("127.0.0.1")

}
