package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}
func handler22(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hel12312lo World, %s!", request.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/2", handler22)
	http.ListenAndServe(":8080", nil)
}

#121121231
#添加代码
