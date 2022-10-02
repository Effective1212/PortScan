package main


import (
"fmt"
"net"


)



func main() {


	d := net.Dialer{Timeout: 999999999}
	conn, err := d.Dial("tcp", "1.1.1.1:445")
	if err != nil {
		fmt.Println("Close")
		fmt.Println(err)
	} else {
		conn.Close()
		fmt.Println("Open")
		}


    }
