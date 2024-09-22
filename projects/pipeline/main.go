package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(in, out chan string) {
	usedStrokes := ""
	for value := range in {
		if strings.Contains(usedStrokes, value) {
			continue
		} else {
			usedStrokes += value + " "
		}
	}
	for _, value := range strings.Split(usedStrokes, " ") {
		out <- value
	}
	close(out)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	
	go func() {
	   defer close(inputStream)
	
	   for _, r := range "112334456" {
		  inputStream <- string(r)
	   }
	}()
	
	for x := range outputStream {
	   fmt.Print(x)
	}
}
