package main

import (
	"fmt"
	"log"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {
	fmt.Println("Hello Clone Website")
}
