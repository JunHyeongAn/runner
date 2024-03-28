package runner

import (
	"fmt"
	"log"
	"net"

	"acky.io/socket_server/config"
)

var (
	CONNECTION map[net.Conn]int
)

func Run(handler func(net.Conn)) {
	// config.yml 파일 읽어오기
	var c config.Config
	c.GetConfig()

	// 현재 접속 connection map
	CONNECTION = make(map[net.Conn]int)
	listener, err := net.Listen("tcp", ":"+fmt.Sprint(c.Server.Port))
	fmt.Println("Start ACKY TCP Server at ", c.Server.Port)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		CONNECTION[conn] = len(CONNECTION)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if handler != nil {
			fmt.Println("call")
			go handler(conn)
		} else {
			fmt.Println("Please give a Handler to Runner.Run(*)")
		}
	}
}

// func tcpConnHandler(conn net.Conn) {
// 	fmt.Println("Accept connection from", conn.RemoteAddr())

// 	buffer := make([]byte, 1024)
// 	for {
// 		n, err := conn.Read(buffer)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// resultChannel <- buffer[:n]
// 		fmt.Printf("CLIENT[%d] : %s\n", CONNECTION[conn], buffer[:n])
// 	}
// }
