package main

import (
    "os"
    "fmt"
    "strings"
    "net/http"
    "path/filepath"
    //"encoding/json"
    "github.com/codegangsta/martini"
)

type Visitor struct{}

func (v *Visitor) VisitDir(path string, f *os.FileInfo) bool {
    fmt.Println(path)
    return true
}

func (v *Visitor) VisitFile(path string, f *os.FileInfo) {
    fmt.Println(path)
}


func main() {
    path_prefix := "Z:\\"

    m := martini.Classic()

    // Parse the URL
    m.Get("/files/**", func(w http.ResponseWriter, r *http.Request) {
        message := r.URL.RequestURI()
        prefix := "/files/"
        elite := strings.Split(message, prefix)[1]

        v := &Visitor{}
        errors := make(chan os.Error, 64)

        if len(elite) > 0 {
            // Step in the real folder
            real_path := filepath.ToSlash(path_prefix) + elite
            filepath.Walk(real_path, v, errors)

            select {
                case err := <- errors:
                    panic(err)
                default:
            }
        }
    });

    m.Run()
}
