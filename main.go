package main

import (
	"fmt"
	"os"
	"strings"
)

 func main(){
	filePointer,err := os.Open("messages.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fullLine := ""
	for {
		b := make([]byte,8)
		_,readError := filePointer.Read(b)
		if readError != nil {
			//println(readError)
			fmt.Printf("read: %s\n",fullLine)
			break
		}
		currentString := string(b)
		splitedString := strings.Split(currentString,"\n")
		//fmt.Printf("parts %s",splitedString[0])
		lengthOfString := len(splitedString)
		//fmt.Println("length of splitted array %v and value ",len(splitedString))
		fullLine += splitedString[0]
		//fmt.Printf("*****%s",fullLine)
		if lengthOfString > 1 {
			//this mean end of line character is in the stream so we need to break the line
			fmt.Printf("read: %s\n",fullLine)
			fullLine =  ""
			fullLine += splitedString[1]
		}
	}
 }
