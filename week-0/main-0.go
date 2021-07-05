package main

import (
	"net/http"
	"fmt"
	"log"
	"io/util"
)

func main(){
	http HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		log.Println("Hello World")
		data, err := ioutil.ReadAll(req.body)
		if(err!nil){
			http.Error(res, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(res, "Hello %s\n", data)
	}

	http.HandleFunc("/goodbye", func(){
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
})

