package main

//http://demo.vemic.com:5000/files/AAPP/旅行箱/1.0/test

import (
    "github.com/codegangsta/martini"
    "net/http"
    "fmt"
)


func main() {
    m := martini.Classic()

    // Parse the URL
    m.Get("/files/**", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, r.URL.RequestURI())
    });

    m.Run()
}

func stepFolder() {
}

func generateCollection() {
}

func callback() {
}
