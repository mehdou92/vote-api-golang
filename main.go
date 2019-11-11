package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

// func main() {

// 	print.Terminal("coucou")
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	//print.Terminal("hello")

	r := gin.Default()
	//handler.InitUser(r, postgresql.New())
	r.Run(":8080")
}
