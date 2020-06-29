package playground

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
)

func GR() {
	server := &http.Server{Addr: "0.0.0.0:8888"}
	var gracefulChild bool
	var err error
	var l net.Listener

	flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")
	if gracefulChild {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		l, err = net.FileListener(f)
	} else {
		log.Print("main: listening on a new file descriptor.")
		l,err = net.Listen("tcp",server.Addr)
	}
	if gracefulChild{
		parent := syscall.Getppid()
		log.Printf("main:kill parent pid %v",parent)
		//syscall
	}

}
