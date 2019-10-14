package main

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"github.com/kuruvi-bits/transform/services"
)

func main() {
	http.HandleFunc("/exif", func(rw http.ResponseWriter, req *http.Request) {
		albumName := req.URL.Query().Get("albumName")
		photoName := req.URL.Query().Get("photoName")
		message := services.Message{
			AlbumName: albumName,
			PhotoName: photoName,
		}
		exif := services.Exif(message)
		result, _ := json.Marshal(&exif)
		rw.WriteHeader(http.StatusOK)
 	    rw.Write(result)
	})
	http.ListenAndServe(":1515", nil)
}
