package proxy

/* A Bunch of needed functions to later call*/

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
)

var Ports_c2 [2]string //Fortnite grr structure to retrieve the Opened Ports on the System, (For c2 use later on.)

func main_ReturnRndmPorts(original_port string, port_to_be_forwarded_to string) {
	Ports_c2[0] = original_port
	Ports_c2[1] = port_to_be_forwarded_to
}

func Rndm_Port() string {
	port_num := rand.Intn(20000)
	port_num += 1000

	return fmt.Sprint(port_num)
}

func prepare() (string, string, string) {

	var typeOfConnection = "tcp"
	original_port := Rndm_Port()
	port_to_be_forwarded_to := Rndm_Port()

	/* With my Luck, i know i will get 2 of the same Ports, so lets add some retarded logic that excludes an equal port on BOTH Rndm Ports :D .*/
	for original_port == port_to_be_forwarded_to {
		port_to_be_forwarded_to = Rndm_Port()
	}

	return typeOfConnection, original_port, port_to_be_forwarded_to

}

// Input Port to be proxied to (On Other Machine) (This Software will activate the Port and io pipe it to the other socket. )
func SocketConnect(typeOfConnection string, ip string, port string) net.Conn {
	fmt.Println("Establishing Connection...")

	var address = net.JoinHostPort(ip, port) //Bruh they wrote their own function for non-ipv6 based hosts, as String concatination does not work with ipv6.
	connection, err := net.Dial(typeOfConnection, address)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Port-Closed or Filtered")
		os.Exit(1)
	}

	println("Port: %d, is open.", port)
	return connection
}

func Listen(typeOfConnection string, port string) net.Listener {
	println("Enabling Forwarding, starting listener on %d:%s", port)
	listener, err := net.Listen(typeOfConnection, port)

	if err != nil {
		println("Something went wrong, maybe Port is being used / no perms?")
		os.Exit(1)
	}

	return listener
}

func EnableForwarding_tcp(ip string) {
	println("Getting Piping STDIN to STDOUT )")

	// written, err := io.Copy()

	typeOfConnection, port_original, port_to_be_forwarded_to := prepare()
	main_ReturnRndmPorts(port_original, port_to_be_forwarded_to)
	println("Port_Original: %s, Port_Forwarded_To: %s", port_original, port_to_be_forwarded_to)

	origin_list := Listen(typeOfConnection, port_original)

	for {
		/*
			Client A connects to :1234 → Proxy creates Pipe A to :forwarded_port

			Client B connects to :1234 → Proxy creates Pipe B to :forwarded_port

			Client C connects to :1234 → Proxy creates Pipe C to :forwarded_port

			Even i forgot, a port can be just considered as the identifier to connect to, and things to be served over, meaning multiple pipeliens of DataStreams can be obviously achieved.
		*/

		origin_conn, err := origin_list.Accept()

		if err != nil {
			println("Something happend while ACCEPTING() connection..., exiting.")
			os.Exit(1)
		}

		/*
			var address = net.JoinHostPort(ip, port_to_be_forwarded_to)
			pipe_conn, err := net.Dial("tcp", address) // Where to forward.
		*/

		pipe_conn := SocketConnect(typeOfConnection, ip, port_to_be_forwarded_to)

		if err != nil {
			fmt.Println("Pipe Connection, using net_dial was not initialized, exiting...)")
			os.Exit(1)
		}

		// 2 Go Routines running at the same time.
		go func() {
			defer origin_conn.Close()
			defer pipe_conn.Close()
			io.Copy(origin_conn, pipe_conn)
		}()
		go func() {
			defer pipe_conn.Close()
			defer origin_conn.Close()
			io.Copy(pipe_conn, origin_conn)
		}()
	}
}
