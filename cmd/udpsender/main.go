package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)


 func main(){
	//println("Starting listner")
	listner,err := net.ResolveUDPAddr("udp",":42069")
	if err != nil {
		log.Fatalf("Unable to resolve the udp address %v",err)
		//os.Exit(1)
	
	}
	udpConnection,err2 :=net.DialUDP("udp",nil,listner)
	if err2 != nil {
		log.Fatalf("Unable to create a udp connection %v",err2)
		//os.Exit(1)
	
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		print(">")
		st,err := reader.ReadString(',')
		if err != nil {
			fmt.Printf("errrrr %v",err)
		}
		udpConnection.Write([]byte(st))
	}
 }
