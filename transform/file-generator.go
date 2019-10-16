package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func createMDFile(fileName string) {
    content := fmt.Sprintf(
    `---
title: "%s"
date: "%v"
description: "%s"
---`, "my album", time.Now(), "/pics/family/coverimage.jpg")
    d1 := []byte(content)
    // d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile(fileName, d1, 0644)
    check(err)
        
}

func addToFile(fileName string) {
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()
    
    content := fmt.Sprintf(
        `---
title2: "%s"
date: "%v"
description: "%s"
---`, "my album", time.Now(), "/pics/family/coverimage.jpg")
        // d1 := []byte(content)

    if _, err = f.WriteString("\n" + content); err != nil {
        panic(err)
    }
}

func main() {
    createMDFile("./dat1.txt")
    addToFile("./dat1.txt")
    f, err := os.Create("./dat2.txt")
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()

}