package main

import "fmt"

func main() {
	msg := sayhello("Test")
	fmt.Println(msg)
}

func sayhello(name string) string{
	return fmt.Sprintf("Hello %s", name)
}


