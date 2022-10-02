package main


import (
"fmt"
"net"
//"io/ioutil"
"os"
"bufio"

)




func main() {

var ip string

	f, err := os.Open("ips.txt")
	if err != nil {
    	fmt.Printf("error opening file: %v\n",err)
	}


	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        fmt.Println(scanner.Text())

				ip = scanner.Text()
				ip = ip+":445"

				d := net.Dialer{Timeout: 999999999}
				conn, err := d.Dial("tcp", ip)

				if err != nil {
					fmt.Println("Close")
					} else {
						conn.Close()
						fmt.Println("Open")
					}

    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

//	r := bufio.NewReader(f)
//	s, e := bufio.ReadLine(r)
//	for e == nil {
//    fmt.Println(s)
//    s,e = bufio.ReadLine(r)
//	}


	//////////

	//	conn, err := net.Dial("tcp", ":445")
//	if err != nil {
//		fmt.Println("Close")
//	} else {
//		conn.Close()
//		fmt.Println("Open")
//		}



}
