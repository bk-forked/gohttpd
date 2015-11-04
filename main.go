package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	webdir := flag.String("d", "./www", "Host directory")
	port := flag.Int("p", 80, "Listen port")
	flag.Parse()
	fmt.Println(fmt.Sprintf("Now hosting %s on port %d", *webdir, *port))
	go func(port int, webdir string) {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(webdir)))
		if err != nil {
			fmt.Println(err)
		}
	}(*port, *webdir)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
forever:
	for {
		select {
		case <-sig:
			fmt.Println("Signal recieved, now stop and exit")
			break forever

		}

	}
}
