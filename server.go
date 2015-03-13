package main

//http://demo.vemic.com:5000/files/AAPP/旅行箱/1.0/test

import (
    "log"
    "net/http"
    "os"
    "sort"
    "io/ioutil"
)


func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}

func FileMTime(file string) (int64, error) {
    f, err := os.Stat(file)
    if err != nil {
        return 0, err
    }
    retunr f.ModT