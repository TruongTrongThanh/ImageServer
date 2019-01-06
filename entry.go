package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/TruongTrongThanh/ImageServer/goenv"
	midd "github.com/TruongTrongThanh/ImageServer/middleware"
	"github.com/TruongTrongThanh/ImageServer/repository"
	"github.com/TruongTrongThanh/ImageServer/router"
)

func main() {
	// Environment variables setup
	goenv.LoadEnv()

	// Database setup
	repository.Connect()
	repository.CreateImageTable()

	// Router setup
	mux := http.NewServeMux()
	rtr := router.Init("")
	rtr.MappingTo(mux)

	// File Server setup
	fsPath := os.Getenv("FileServerPath")
	fmt.Printf("File Server serve at: %s\n", fsPath)
	dir := http.Dir(filepath.ToSlash(os.Getenv("StoredPath")))
	fsHandler := http.StripPrefix(fsPath, http.FileServer(dir))
	mux.Handle(fsPath, midd.Filter(fsHandler, fsPath))

	// Serve
	isSSL, err := strconv.ParseBool(os.Getenv("ssl"))
	if err != nil {
		panic(err)
	}
	hostname, port := os.Getenv("hostname"), os.Getenv("port")
	if isSSL {
		panic("Not supported SSL right now")
	} else {
		fmt.Printf("Serve at http://%s:%s\n", hostname, port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", hostname, port), mux))
	}
}
