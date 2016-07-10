package main

import (
    "io"
	"log"
	"net/http"
	"time"
)

/*
func main(){
	// 设置路由
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "this is a Hello work")
}
*/
/*
func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "URL:" + r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is a Hello version2")
}
*/

var mux map[string]func(http.ResponseWriter, *http.Request)

func main(){
	server := http.Server{
		Addr: ":8080",
		Handler: &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "this is a hello version 4")
}

func sayBye(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "this is a bye version 4")
}
