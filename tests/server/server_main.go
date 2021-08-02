package main

import (
	"fmt"
	"log"

	"github.com/fiojay1997/proglog/internal/server"
)

func main() {
	fmt.Println("hello")
	fmt.Println("world")
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
	defer srv.Close()
}
