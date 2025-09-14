package main

import (
	"fmt"

	"./proxy"
)

func main() {
	fmt.Println("Starting D4ckf3ce....")
	/* What should happen is , 2 rndmly ports in the range from 1200 - 22000 should be picked, with port != port2, those will be printed out to the console.
	   We enable forwarding on the listening Port, connect to the other rndmly choosen Port, and pipe our connection through there, this makes it harder to track.*/
	proxy.EnableForwarding_tcp("127.0.0.1")
	port_original := proxy.Ports_c2[0]
	port_to_be_forwarded_to := proxy.Ports_c2[1]
	println("Port_Original: %s, Port_Forwarded_To: %s", port_original, port_to_be_forwarded_to)

}
