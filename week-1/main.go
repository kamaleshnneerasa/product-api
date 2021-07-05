package main

import (
	"/Users/kamaleshneerasa/Desktop/FullStackDev/product-api/week-1/handlers"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"time"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlerss.NewHello(l)

	sm := http.NewServeMux()
	sm.handle("/", helloHandler)
	sm.handle("/goodbye", goodbyeHandler)

	server = &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second
	}

	//Asynchronous function??
	go func(){
		err = server.ListenAndServe()
		if(err!=nil){
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.notify(sigChan, os.interrupt)
	signal.notify(sigChan, os.interrupt)
	
	sig := <-sigChan
	l.Println("Received Terminate, Graceful Shutdown", sig)

	tc, _ := context.WithTimeout(context.Backgopund(), 30*time.Second)
	server.Shutdown(tc)

}

//go run main.go
//curl localhost:9090 -d 'Kamal'
