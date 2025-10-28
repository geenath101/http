package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// func getLinesChannel(f io.ReadCloser) chan string {
// 	readChannel := make(chan string)
// 	go func() {
// 		fullLine := ""
// 		for {
// 			b := make([]byte,8)
// 			_,readError := f.Read(b)
// 			if readError ==  io.EOF {
// 				//fmt.Printf("read: %s\n",fullLine)
// 				readChannel <- fullLine
// 				break
// 			}
// 			currentString := string(b)
// 			splitedString := strings.Split(currentString,"\n")
// 			//fmt.Printf("parts %s",splitedString[0])
// 			lengthOfString := len(splitedString)
// 			//fmt.Println("length of splitted array %v and value ",len(splitedString))
// 			fullLine += splitedString[0]
// 			//fmt.Printf("*****%s",fullLine)
// 			if lengthOfString > 1 {
// 				//this mean end of line character is in the stream so we need to break the line
// 				//fmt.Printf("read: %s\n",fullLine)
// 				readChannel <- fullLine
// 				fullLine =  ""
// 				fullLine += splitedString[1]
// 			}
// 		}
// 		defer close(readChannel)
// 	}()
// 	return  readChannel
// }


func getLinesChannel(l net.Conn) chan string {
	readChannel := make(chan string)
	go func() {
		fullLine := ""
		for{
			b := make([]byte,8)
			l.Read(b)
			currentString := string(b)
			splitedString := strings.Split(currentString,"\n")
			lengthOfString := len(splitedString)
			fullLine += splitedString[0]
			//fmt.Printf("*****%s",fullLine)
			if lengthOfString > 1 {
				//this mean end of line character is in the stream so we need to break the line
				//fmt.Printf("read: %s\n",fullLine)
				readChannel <- fullLine
				fullLine =  ""
				fullLine += splitedString[1]
			}
		}	
		//defer close(readChannel)
	}()
	return  readChannel
}

 func main(){
	//println("Starting listner")
	listner,err := net.Listen("tcp",":42069")
	if err != nil {
		log.Fatalf("Unable to create the listner %v",err)
		//os.Exit(1)
	
	}
	defer listner.Close()
	// readChannel := getLinesChannel(listner)
	// for v := range readChannel{
	// 	fmt.Printf("read: %s\n",v)
	// }
	conn,err2 := listner.Accept()
	if err2 != nil{
		log.Fatalf("Unable to create the connection %v",err2)
	}
	if conn != nil{
		//println("Connection created")
	}
	readChannel := getLinesChannel(conn)
	for v := range readChannel{
		fmt.Printf("%s\n",v)
	}
 }
