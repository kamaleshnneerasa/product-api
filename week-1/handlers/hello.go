package handlers //Handlers package also called controller

import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct{
	l *log.Logger
}

//func name(',' seperated arguments) <return type> 
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}  //NewHello returns a pointer to a Hello object
}

//This method ca ne called on any pointer to Hello object
func (h *Hello) ServerHTTP(res http.ResponseWriter, req *http.Request){
	
	h.l.Println("Hello World")

	data, err := ioutil.ReadAll(req.body)
	if(err!nil){
		http.Error(res, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Hello %s\n", data)
}



