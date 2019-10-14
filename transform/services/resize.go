package services

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
	utils "github.com/kuruvi-bits/transform/utils"
)

func Resize(message utils.Message) {
    albumName := message.AlbumName
    photoName := message.PhotoName
    path := "./" + albumName + "/" + photoName
    utils.CreateDirIfNotExist(albumName)

    url := utils.GetResizeURL(message)

    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()

    //open a file for writing
    file, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Success!")
}
