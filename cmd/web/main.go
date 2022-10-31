package main

import (
	"log"
	"os"

	"github.com/CloudyKit/jet"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
}

type server struct {
	host string
	port string
	url  string
}

func main() {

	server := server{
		host: "localhost",
		port: "8080",
		url:  "http://localhost:8080",
	}

	app := &application{
		server:  server,
		appName: "CloneNews",
		debug:   true,
		infoLog: log.New(os.Stdout, "Info\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "Info\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))
	}

	if err := app.listenAndServe(); err != nil {
		log.Fatal(err)
	}
}
