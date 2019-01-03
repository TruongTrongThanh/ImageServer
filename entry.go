package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TruongTrongThanh/ImageServer/goenv"
	"github.com/TruongTrongThanh/ImageServer/repository"
	"github.com/TruongTrongThanh/ImageServer/router"
)

func main() {
	// Setup environment variables
	goenv.LoadEnv()

	// Setup Database
	repository.Connect()
	repository.CreateImageTable()

	// Setup router
	mux := http.NewServeMux()
	rtr := router.Init("")
	rtr.MappingTo(mux)

	// Serve
	fmt.Println("Serve at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
