package go_route

import (
	"fmt"
	"net"
	"os"
	"time"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	fmt.Print("Starting 3 clients\n")

	for i:=0; i<3; i++ {

		client := func(inName string) {

			defer wg.Done() // added this line
			fmt.Printf("Client <%s> started\n", inName)

			conn, err := net.Dial("tcp", ":8000")
			if err != nil {
				fmt.Printf("[%s] Error while connecting to the server: %s", inName, err.Error())
				os.Exit(1)
			}

			n := 0
			sleepDuration, _ := time.ParseDuration("2s")
			for {
				message := fmt.Sprintf("[%s] > message %d\n", inName, n)
				fmt.Printf("%s", message)
				_, err := conn.Write([]byte(message))
				if nil != err {
					fmt.Sprintf("[%s] Error while writing data to the socket: %s", inName, err.Error())
					os.Exit(1)
				}
				time.Sleep(sleepDuration)
				n++
			}
		}

		name := fmt.Sprintf("Client%d", i)
		fmt.Printf("Starting client <%s>...\n", name)
		wg.Add(1) // moved this line
		go client(name)
		fmt.Print("Done\n")
	}

	wg.Wait()
}