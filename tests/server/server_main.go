package main

import (
	"fmt"
	"log"

	"github.com/fiojay1997/proglog/internal/server"
)

func main() {
	fmt.Println("hello")
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
