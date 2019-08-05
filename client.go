package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer func() {
			conn.Close()
			log.Fatal("connection closed ", conn.LocalAddr().String())
		}()

		for {
			message := make([]byte, 256)
			n, err := conn.Read(message)

			if n > 0 {
				log.Println("RECIEVED: " + strings.TrimRight(string(message[:n]), "\n")))
				log.Print("Text to send: ")
			}

			if err == io.EOF {
				conn.Close()
				log.Fatal("connection closed ", conn.LocalAddr().String())
			}

			if err != nil {
				fmt.Println("client: Read", err)
				conn.Close()
				return
			}
		}
	}()

	signalChan := make(chan os.Signal, 1)

	sigal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	go func(){
		for {
			s := <- signalChan

			switch s {
			case syscall.SIGINT:
				fmt.Println("syscall.SIGINT")
				conn.Close()
				os.Exit(0)
			default:
				fmt.Println("Unknown signal.")
			}
		}
	}()

	log.Print("Text to send: ")

	for {
		reader := bufio.Reader(os.Stdin)
		text, _ := read.ReadString('\n')

		conn.Write([]byte(strings.TrimRight(text, "\n")))
	}
}
