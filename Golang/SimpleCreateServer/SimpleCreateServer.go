//package SimpleCreateServer
//SourceCode:====https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
package main

import (
	"fmt"
	"log"
	"net/http"
	//"strconv"
	//"sync"
)

//var counter int
//var mutex = &sync.Mutex{}
//
//func echoString(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, ("hello Counter :=="+ strconv.Itoa(counter)))
//}
//
//func incrementCounter(w http.ResponseWriter, r *http.Request) {
//	mutex.Lock()
//	counter++
//	fmt.Fprintf(w, strconv.Itoa(counter))
//	mutex.Unlock()
//}

func main() {
	//http.HandleFunc("/", echoString)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	//http.HandleFunc("/increment", incrementCounter)

	//http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	//	fmt.Fprintf(w, "Hi")
	//})


	log.Fatal(http.ListenAndServe(":8081", nil))

}
